package main

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    serverURL := "http://localhost:5000/upload"
    filePath := "1.jpg" // Replace with path to your image
    outputFilePath := "2.jpg"  // File path to save the returned image

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    // Prepare a multipart form request
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
    if err != nil {
        fmt.Printf("Error creating form file: %v\n", err)
        return
    }

    _, err = io.Copy(part, file)
    if err != nil {
        fmt.Printf("Error copying file: %v\n", err)
        return
    }
    writer.Close()

    // Send the request to the Flask server
    request, err := http.NewRequest("POST", serverURL, body)
    if err != nil {
        fmt.Printf("Error creating request: %v\n", err)
        return
    }
    request.Header.Add("Content-Type", writer.FormDataContentType())

    // Perform the request
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        fmt.Printf("Error sending request: %v\n", err)
        return
    }
    defer response.Body.Close()

    // Handle the response
    if response.StatusCode == http.StatusOK {
        // Create a file to save the returned image
        outFile, err := os.Create(outputFilePath)
        if err != nil {
            fmt.Printf("Error creating file: %v\n", err)
            return
        }
        defer outFile.Close()

        // Write the response body to file
        _, err = io.Copy(outFile, response.Body)
        if err != nil {
            fmt.Printf("Error writing response to file: %v\n", err)
            return
        }

        fmt.Println("Processed image saved as:", outputFilePath)
    } else {
        fmt.Printf("Server returned non-OK status: %v\n", response.Status)
    }
}
