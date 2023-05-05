package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func log(writer io.Writer, msg string) {
	timestamp := time.Now().Format(time.RFC3339)
	withTimestamp := fmt.Sprintf("%s: %s", timestamp, msg)
	fmt.Fprint(writer, withTimestamp)
}

func main() {
	endpointStr := os.Getenv("ENDPOINT")
	if endpointStr == "" {
		endpointStr = "https://checkip.amazonaws.com"
	}

	intervalStr := os.Getenv("INTERVAL")
	if intervalStr == "" {
		intervalStr = "1m"
	}
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		log(os.Stderr, fmt.Sprintf("ERROR - Failed to parse interval value %s\n", intervalStr))
		return
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	for {
		resp, err := client.Get(endpointStr)
		if err != nil {
		} else {
			if resp.StatusCode != 200 {
				log(os.Stderr, fmt.Sprintf("ERROR - HTTP request returned status code %d\n", resp.StatusCode))
			} else {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log(os.Stderr, "ERROR - Failed to read response body.\n")
				} else {
					if len(body) > 0 && body[len(body)-1] != '\n' {
						body = append(body, '\n')
					}
					log(os.Stdout, string(body))
				}
			}
			resp.Body.Close()
		}

		time.Sleep(interval)
	}
}
