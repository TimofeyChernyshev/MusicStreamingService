package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	track := r.URL.Query().Get("track")
	safePath := filepath.Join("/app/storage/song", filepath.Base(track))
	file, err := os.Open(safePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Println(err)
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "audio/mpeg")
	http.ServeContent(w, r, track, stat.ModTime(), file)
}
