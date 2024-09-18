package main;
import (
	"fmt"
	"net/url"
)

func main(){
	// basicUrlParsing()
	urlValidation()
}


func urlValidation(){
// You can validate a URL by checking whether url.Parse returns an error, or by checking the scheme and host to ensure they meet your requirements:

rawURL := "https://example.com:8080/search?q=golang&sort=asc#section";

parsedURL, err := url.Parse(rawURL)
if err != nil {
	fmt.Println("Error:", err)
	return
}

if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
    fmt.Println("Invalid URL scheme");
	return;
}

if parsedURL.Host == "" {
    fmt.Println("URL must have a host")
	return
}

fmt.Println("Valid URL")
return;
}


func basicUrlParsing(){

// 	Scheme: To determine if it's http, https, or another protocol.
// Host and Port: To get the domain and port number.
// Path: To identify the resource being requested (like /users or /products).
// Query Parameters: To capture and handle key-value pairs in the query string (?key=value).
// Fragment: To identify and navigate to specific sections of a page (#section).

	rawURL := "https://example.com:8080/search?q=golang&sort=asc#section"
    
    // Parse the URL
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // // Access different parts of the URL
    // fmt.Println("Scheme:", parsedURL.Scheme)       // https
    // fmt.Println("Host:", parsedURL.Host)           // example.com:8080
    // fmt.Println("Path:", parsedURL.Path)           // /search
    // fmt.Println("Query:", parsedURL.RawQuery)      // q=golang&sort=asc
    // fmt.Println("Fragment:", parsedURL.Fragment)   // section
    // fmt.Println("Port:", parsedURL.Port())         // 8080
    // fmt.Println("Query Params:", parsedURL.Query()) // map[q:[golang] sort:[asc]]


	parameterQueries:=parsedURL.Query();
	fmt.Println(parameterQueries["q"][0]);

}
