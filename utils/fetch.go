package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var IPCONFIG_URL string = "https://ipconfig.org"

func FetchUsersIp() string {
	// Create a new request
	req, err := http.NewRequest("GET", IPCONFIG_URL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// Set the headers
	req.Header.Set("User-Agent", "curl/8.7.1")
	req.Header.Set("Accept", "*/*")

	// Create HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("No response sadly :(")
		return ""
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return strings.Trim(string(body), "\n")
}
