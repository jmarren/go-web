package loadbalancer

import (
	// "context"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Server() {
	proxyStr := "http://localhost:8080"

	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Fatalf("Error parsing proxy URL: %v", err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		// DialContext: func(ctx context.Context, network, addr string) (net.Conn, error)  {
		//
		// }
		// DialContext func(ctx context.Context, network, addr string) (net.Conn, error)
	}

	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get("https://example.com")
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	log.Printf("Response: %s", body)
}
