package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func scanFolder(ipPort string) {
    // Send HTTP request to server at ipPort to get directory listing
    resp, err := http.Get("http://" + ipPort + "/files/")
    if err != nil {
        log.Fatal("Error scanning folder:", err)
    }
    defer resp.Body.Close()

    // Read response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error reading response body:", err)
    }

    // Print directory listing
    fmt.Printf("Directory listing from %s:\n", ipPort)

    // Extract file names from HTML links
    fileNames := extractFileNames(string(body))
    for _, fileName := range fileNames {
        fmt.Println(fileName)
    }
}

func extractFileNames(htmlContent string) []string {
    var fileNames []string

    // Find all occurrences of "<a href=" in the HTML content
    startTag := "<a href=\""
    endTag := "\">"
    startPos := strings.Index(htmlContent, startTag)

    for startPos != -1 {
        // Move startPos to the beginning of the file name
        startPos += len(startTag)

        // Find end position of file name
        endPos := strings.Index(htmlContent[startPos:], endTag)
        if endPos == -1 {
            break
        }

        // Extract file name and add to list
        fileName := htmlContent[startPos : startPos+endPos]
        fileNames = append(fileNames, fileName)

        // Move to next "<a href=" occurrence
        startPos = strings.Index(htmlContent[startPos+endPos:], startTag)
    }

    return fileNames
}