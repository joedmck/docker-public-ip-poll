package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	intervalStr := os.Getenv("INTERVAL")
	if intervalStr == "" {
		intervalStr = "1m"
	}

	endpointStr := os.Getenv("ENDPOINT")
	if endpointStr == "" {
		endpointStr = "https://checkip.amazonaws.com"
	}

	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		fmt.Println("Error parsing interval value:", err)
		return
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for {
		resp, err := client.Get(endpointStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: HTTP response status code: %d", resp.StatusCode)
		} else {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: Failed to read response body: ", err)
			} else {
				if len(body) > 0 && body[len(body)-1] != '\n' {
					body = append(body, '\n')
				}
				timestamp := time.Now().Format(time.RFC3339)
				bodyWithTimestamp := fmt.Sprintf("%s: %s", timestamp, string(body))
				fmt.Print(bodyWithTimestamp)
			}
			resp.Body.Close()
		}

		time.Sleep(interval)
	}
}
