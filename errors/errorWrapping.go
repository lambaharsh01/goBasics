package main;

import ("fmt"; "errors"; "os" );



func openFile(filename string) error {
	// Try to open the file
	result, err := os.Open(filename)  // This line tries to open the file // os.Open function returns a file object (*os.File), not the contents of the file or anything human-readable

	// Check if an error occurred while trying to open the file
	if err != nil {
		// If there is an error, return a new error that wraps the original error with more context
		// %w is used to wrap the original error (err) inside a new error message.

		return fmt.Errorf("failed to open file %s: %w", filename, err)

		// if i do return errors.New(err) it will not work as explained below
		// errors.New expects a string argument, not an error. You cannot pass an error directly to errors.New. Instead, you should use fmt.Errorf to wrap an existing error with additional context.
	}

	// If no error, return nil (meaning everything is OK)
	return nil
}


func main(){
	err := openFile("doesNotExists.txt")
	if err != nil {
		fmt.Println("Error:", err)
		
		// Use errors.Is to check if the error is due to a specific underlying error (e.g., os.ErrNotExist)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file does not exist.")
		}
	}
}