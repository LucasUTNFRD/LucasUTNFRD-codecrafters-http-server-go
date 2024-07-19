package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/echo/") {
		http.NotFound(w, r)
		return
	}
	message := strings.TrimPrefix(r.URL.Path, "/echo/")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(message)))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, message)
}

func userAgentHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(userAgent)))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, userAgent)
}

func fileHandler(w http.ResponseWriter, r *http.Request, dir string) {
	log.Println("debugg")
	fileName := strings.TrimPrefix(r.URL.Path, "/files/")
	filePath := filepath.Join(dir, fileName)
	log.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		http.NotFound(w, r)
		log.Println("not found")
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, filePath)
}
