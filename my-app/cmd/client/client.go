package main

import (
	"log"
	"net/http"
)

func main() {
	general()
	// generateErrors()
}

func general() {
	for {
		req("GET", "http://panchanandevops.com:31559/devices")
		req("POST", "http://panchanandevops.com:31559/devices")
		req("PUT", "http://panchanandevops.com:31559/devices/123")

		req("GET", "http://panchanandevops.com:31559/devices")
		req("GET", "http://panchanandevops.com:31559/devices")
		req("GET", "http://panchanandevops.com:31559/devices")
		req("GET", "http://panchanandevops.com:31559/devices")
		req("POST", "http://panchanandevops.com:31559/devices")
		req("POST", "http://panchanandevops.com:31559/devices")
		req("POST", "http://panchanandevops.com:31559/devices")
		req("PUT", "http://panchanandevops.com:31559/devices/123")
		req("PUT", "http://panchanandevops.com:31559/devices/123")
		req("PUT", "http://panchanandevops.com:31559/devices/123")
		req("PUT", "http://panchanandevops.com:31559/devices/123")
		req("GET", "http://panchanandevops.com:31559/devices")
		req("GET", "http://panchanandevops.com:31559/devices")
		req("GET", "http://panchanandevops.com:31559/devices")
		req("GET", "http://panchanandevops.com:31559/devices")
		req("POST", "http://panchanandevops.com:31559/devices")
		req("POST", "http://panchanandevops.com:31559/devices")
		req("POST", "http://panchanandevops.com:31559/devices")
		req("PUT", "http://panchanandevops.com:31559/devices/123")
		req("PUT", "http://panchanandevops.com:31559/devices/123")
		req("PUT", "http://panchanandevops.com:31559/devices/123")
		req("PUT", "http://panchanandevops.com:31559/devices/123")

		req("POST", "http://panchanandevops.com:31559/login")
		req("DELETE", "http://panchanandevops.com:31559/devices/123")
	}
}

func generateErrors() {
	for {
		req("POST", "http://panchanandevops.com:31559/login")
		req("DELETE", "http://panchanandevops.com:31559/devices/123")
		req("PUT", "http://panchanandevops.com:31559/devices/123")
		req("POST", "http://panchanandevops.com:31559/login")
		req("DELETE", "http://panchanandevops.com:31559/devices/123")
		req("POST", "http://panchanandevops.com:31559/login")
		req("DELETE", "http://panchanandevops.com:31559/devices/123")
	}
}

func req(method string, url string) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
