# 🔁 Golang Reverse Proxy Server

A simple reverse proxy server written in Go that forwards requests to a target backend and provides a `/status` endpoint to monitor usage.

## 🚀 Features

- ✅ Reverse proxy to any backend
- 📊 `/status` endpoint shows:
  - Total proxied requests
  - Server uptime
  - Backend URL
  - Local server IP
- 🌐 Fully HTTP-based (easy to test locally)

## 🛠 Setup

### 1. Clone the repo (or copy the code)

``bash
git clone <your-repo-url>
cd golang-reverse-proxy
2. Run a local backend (optional)
To test the proxy locally, run this simple backend on port 9000:

// backend.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend: %s", r.URL.Path)
	})

	log.Println("✅ Backend running at http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
Run it with:


go run backend.go
3. Update the proxy target (in main.go)
Change the target URL to your backend:

targetURL, err := url.Parse("http://localhost:9000")
4. Run the proxy
go run main.go
5. Test it
Proxy some requests:
curl http://localhost:8080/
curl http://localhost:8080/hello
Check proxy status:
curl http://localhost:8080/status
Example output:
{
  "server_ip": "192.168.1.100",
  "status": "running",
  "target_backend": "http://localhost:9000",
  "total_requests": 2,
  "uptime": "35s"
}
📦 Dependencies
Standard Go library only (net/http, httputil, etc.)

No third-party packages needed

📘 License
MIT License

✍️ Author
Hein (with help from ChatGPT 😉)