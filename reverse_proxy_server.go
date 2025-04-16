package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync/atomic"
	"time"
)

var (
	requestCount uint64
	startTime    = time.Now()
)

func main() {
	//targetURL, err := url.Parse("http://httpbin.org")
	targetURL, err := url.Parse("http://localhost:9000")
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/status" {
			statusHandler(w, targetURL)
			return
		}

		// ✅ Increment request count only on actual proxy calls
		atomic.AddUint64(&requestCount, 1)
		log.Printf("Proxying request: %s", r.URL.Path)
		proxy.ServeHTTP(w, r)
	})

	log.Println("✅ Reverse proxy server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func statusHandler(w http.ResponseWriter, target *url.URL) {
	count := atomic.LoadUint64(&requestCount)
	uptime := time.Since(startTime).Round(time.Second)
	ip := getLocalIP()

	status := map[string]interface{}{
		"status":         "running",
		"target_backend": target.String(),
		"total_requests": count,
		"uptime":         uptime.String(),
		"server_ip":      ip,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, "Failed to encode status", http.StatusInternalServerError)
	}
}

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		hostname, _ := os.Hostname()
		return hostname
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
