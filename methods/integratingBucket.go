// File: s3_upload.go
// Description: Uploads regular files and HLS-encoded MP4 videos to AWS S3.
// Requirements:
// - Go 1.18+
// - ffmpeg installed (for HLS conversion)
// - Gin framework (github.com/gin-gonic/gin)
// - AWS SDK v2 (aws-sdk-go-v2)
// - Set your own logic r hardcode values for demo/testing

package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/ratelimit"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
)


const (
	HlsTestCount       = 3  // Canary uploads
	HlsMaxConcurrency  = 10 // Adjust based on memory and CPU
)

func UploadFilesToS3(c *gin.Context, uploadFileName string, fileDirectory string) (string, error) {

	if uploadFileName == "" {
		return "", nil
	}

	fileHeader, err := c.FormFile(uploadFileName)
	if err != nil {
		return uploadFileName, nil
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	// S3 CONFIGURATION
	var environment string = "YOUR_ENVIRONMENT"
	var accessKeyId string = "YOUR_S3_ACCESS_KEY_ID"
	var secretAccessKey string = "YOUR_S3_SECRET_ACCESS_KEY"
	var region string = "YOUR_S3_REGION"
	var bucketName string = "YOUR_S3_BUCKET_NAME"

	var fileName string = fileHeader.Filename
	var mimeType string = fileHeader.Header.Get("Content-Type")
	var timestamp string = time.Now().Format("2006-01-02_15-04-05")

	var fileSuffix string = filepath.Ext(fileName)
	var fileNameWithoutSuffix string = strings.TrimSuffix(fileName, fileSuffix)

	var filePath string = fmt.Sprintf("%s%s/%s_%s%s", environment, fileDirectory, fileNameWithoutSuffix, timestamp, fileSuffix)


	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, "")))
	if err != nil {
		return "", err
	}

	s3Client := s3.NewFromConfig(cfg)

	_, err = s3Client.PutObject(c, &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(filePath),
		Body:        file,
		ContentType: aws.String(mimeType),
	})

	if err != nil {
		return "", err
	}

	var fileURL string = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, filePath)

	return fileURL, nil
}

func UploadHlsVideoToS3(c *gin.Context, uploadFileName string) (string, error) {

	fileHeader, err := c.FormFile(uploadFileName)
	if err != nil {
		return "", err
	}
	if fileHeader == nil || fileHeader.Filename == "" {
		return "", nil
	}

	var uploadedFileName = fileHeader.Filename
	var uploadedFileExtension = filepath.Ext(uploadedFileName)
	if uploadedFileExtension != ".mp4" {
		return "", fmt.Errorf("only MP4 files are supported")
	}

	rawFileName := strings.TrimSuffix(uploadedFileName, uploadedFileExtension)
	sanitizedFileName := SanitizeFileName(rawFileName)

	rawVideoUploadDir := filepath.Join("public", "rawVideos")
	if err := os.MkdirAll(rawVideoUploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}
	rawVideoPath := filepath.Join(rawVideoUploadDir, sanitizedFileName+".mp4")
	if err := c.SaveUploadedFile(fileHeader, rawVideoPath); err != nil {
		return "", fmt.Errorf("failed to save video file: %w", err)
	}

	// 2️⃣ HLS CONVERSION
	var hlsDirectory string = "/hls"
	hlsConversionDir := filepath.Join("public", "hlsVideos", sanitizedFileName)
	if err := os.MkdirAll(hlsConversionDir, 0755); err != nil {
		_ = os.Remove(rawVideoPath)
		return "", fmt.Errorf("failed to create HLS directory: %w", err)
	}
	var hlsConversionPath string = filepath.Join(hlsConversionDir, "playlist.m3u8")
	var hlsSegmentPattern string = filepath.Join(hlsConversionDir, "segment_%04d.ts")
	// var hlsBaseUrl string = fmt.Sprintf("%s/", sanitizedFileName)

	cmd := exec.Command("ffmpeg",
		"-i", rawVideoPath,
		"-codec", "copy",
		"-start_number", "0",
		"-hls_time", "30",
		"-hls_list_size", "0",
		"-hls_segment_filename", hlsSegmentPattern,
		// "-hls_base_url", hlsBaseUrl,
		"-f", "hls",
		hlsConversionPath,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		_ = os.Remove(rawVideoPath)
		_ = os.RemoveAll(hlsConversionDir)
		return "", fmt.Errorf("FFmpeg conversion failed: %s, %w", string(out), err)
	}

	// 3️⃣ AWS S3 CLIENT SETUP
	var environment string = "YOUR_ENVIRONMENT"
	var accessKeyId string = "YOUR_S3_ACCESS_KEY_ID"
	var secretAccessKey string = "YOUR_S3_SECRET_ACCESS_KEY"
	var region string = "YOUR_S3_REGION"
	var bucketName string = "YOUR_S3_BUCKET_NAME"

	// custom HTTP client to avoid stale keep-alives
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     30 * time.Second,
		},
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, ""),
		),
		config.WithHTTPClient(httpClient),
		config.WithRetryer(func() aws.Retryer {
			return retry.NewStandard(func(o *retry.StandardOptions) {
				o.MaxAttempts = 5
				o.RateLimiter = ratelimit.None
			})
		}),
	)
	if err != nil {
		_ = os.Remove(rawVideoPath)
		_ = os.RemoveAll(hlsConversionDir)
		return "", fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(s3Client)

	// 4️⃣ GATHER ALL HLS FILES
	var allFiles []string
	if err := filepath.Walk(hlsConversionDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			allFiles = append(allFiles, path)
		}
		return nil
	}); err != nil {
		return "", err
	}

	if len(allFiles) == 0 {
		return "", fmt.Errorf("no HLS files found in %s", hlsConversionDir)
	}

	s3BasePath := fmt.Sprintf("%s%s/%s", environment, hlsDirectory, sanitizedFileName)
	for i, p := range allFiles {
		if i >= HlsTestCount {
			break
		}
		if err := uploadSingle(context.TODO(), uploader, bucketName, s3BasePath, hlsConversionDir, p); err != nil {
			_ = os.Remove(rawVideoPath)
			_ = os.RemoveAll(hlsConversionDir)
			return "", fmt.Errorf("canary upload failed on %s: %w", filepath.Base(p), err)
		}
	}

	// 6️⃣ FULL UPLOAD WITH BOUNDED CONCURRENCY
	var semaphore chan struct{} = make(chan struct{}, HlsMaxConcurrency)
	var waitGroup sync.WaitGroup
	var errChannel chan error = make(chan error, 1)

	for _, hlsSegmentPath := range allFiles {
		semaphore <- struct{}{}
		waitGroup.Add(1)

		go func(path string) {
			defer waitGroup.Done()
			defer func() { <-semaphore }()

			if err := uploadSingle(context.TODO(), uploader, bucketName, s3BasePath, hlsConversionDir, path); err != nil {
				select {
				case errChannel <- err:
				default:
				}
			}
		}(hlsSegmentPath)
	}

	go func() {
		waitGroup.Wait()
		close(errChannel)
	}()
	if err := <-errChannel; err != nil {
		_ = os.Remove(rawVideoPath)
		_ = os.RemoveAll(hlsConversionDir)
		return "", err
	}

	// 7️⃣ CLEAN UP & RETURN URL
	_ = os.Remove(rawVideoPath)
	_ = os.RemoveAll(hlsConversionDir)

	directoryURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s/", bucketName, region, s3BasePath)
	return directoryURL, nil
}

// uploadSingle handles one file’s S3 upload.
func uploadSingle(
	ctx context.Context,
	uploader *manager.Uploader,
	bucketName, s3BasePath, baseDir, filePath string,
) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	relPath, err := filepath.Rel(baseDir, filePath)
	if err != nil {
		return err
	}
	s3Key := fmt.Sprintf("%s/%s", s3BasePath, relPath)

	contentType := "application/octet-stream"
	switch filepath.Ext(filePath) {
	case ".m3u8":
		contentType = "application/vnd.apple.mpegurl"
	case ".ts":
		contentType = "video/mp2t"
	}

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(s3Key),
		Body:        f,
		ContentType: aws.String(contentType),
	})
	return err
}

func GeneratePresignedURL(url string) (string, error) {

	if url == "" {
		return "", nil
	}

	var accessKeyID string = "YOUR_S3_ACCESS_KEY_ID"
	var secretAccessKey string = "YOUR_S3_SECRET_ACCESS_KEY"
	var region string = "YOUR_S3_REGION"
	var bucketName string = "YOUR_S3_BUCKET_NAME"

	var expiry time.Duration = time.Minute * 60

	var urlToBeTrimmed string = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", bucketName, region)

	var filePath string = strings.Replace(url, urlToBeTrimmed, "", 1)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		return "", fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	presigner := s3.NewPresignClient(client)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
	}

	presignedURL, err := presigner.PresignGetObject(context.TODO(), params, func(opts *s3.PresignOptions) {
		opts.Expires = expiry
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return presignedURL.URL, nil
}

func AppendPresignedURL(stringPointer *string, stringValueToBeUpdated *string) error {

	if stringPointer == nil {
		return nil
	}

	presignedUrl, err := GeneratePresignedURL(*stringPointer)
	if err != nil {
		return err
	}

	*stringValueToBeUpdated = presignedUrl
	return nil
}

func DeleteS3Path(s3URL string) error {

	// Parse the S3 URL to extract bucket name and key
	parsedURL, err := url.Parse(s3URL)
	if err != nil {
		return fmt.Errorf("invalid S3 URL: %w", err)
	}

	// Extract key from path (removing leading slash)
	var objectKey string = strings.TrimPrefix(parsedURL.Path, "/")

	var accessKeyId string = "YOUR_S3_ACCESS_KEY_ID"
	var secretAccessKey string = "YOUR_S3_SECRET_ACCESS_KEY"
	var region string = "YOUR_S3_REGION"
	var bucketName string = "YOUR_S3_BUCKET_NAME"

	// INIT S3 CLIENT
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, "")))
	if err != nil {
		return fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	var isDirectory bool = strings.HasSuffix(objectKey, "/")

	if !isDirectory {
		if _, err := s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
		}); err != nil {
			return fmt.Errorf("failed to delete file: %w", err)
		}

		return nil
	}

	// LIST ALL THE OBJECTS IN THE DIR WITH THE PREFIX TO DELETE THE ENTIRE DIR
	listObjectsOutput, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(objectKey),
	})

	if err != nil {
		return fmt.Errorf("failed to list objects in directory: %w", err)
	}

	if len(listObjectsOutput.Contents) == 0 {
		return nil //IF NO OBJECTS FOUND IN THE DIR
	}

	objectsToDelete := make([]types.ObjectIdentifier, 0, len(listObjectsOutput.Contents))

	for _, obj := range listObjectsOutput.Contents {
		objectsToDelete = append(objectsToDelete, types.ObjectIdentifier{
			Key: obj.Key,
		})

		if err != nil {
			return fmt.Errorf("failed to list objects in directory: %w", err)
		}

		objectsToDelete := make([]types.ObjectIdentifier, 0, len(listObjectsOutput.Contents))

		for _, obj := range listObjectsOutput.Contents {
			objectsToDelete = append(objectsToDelete, types.ObjectIdentifier{
				Key: obj.Key,
			})
		}

		// DELETE ALL THE OBJECTS IN THE DIR
		_, err = s3Client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
			Bucket: aws.String(bucketName),
			Delete: &types.Delete{
				Objects: objectsToDelete,
				Quiet:   aws.Bool(false),
			},
		})

	}

	if err != nil {
		return fmt.Errorf("failed to delete directory contents: %w", err)
	}

	return nil

}
