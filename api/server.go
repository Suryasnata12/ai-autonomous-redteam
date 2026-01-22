package api

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"ai-autonomous-redteam/api/handlers"
)

func Start() {
	// 🔹 Resolve absolute path to web/
	exePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	webDir := filepath.Join(exePath, "web")
	log.Println("Serving UI from:", webDir)

	// 🔹 Serve static UI
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	// 🔹 API endpoint
	http.HandleFunc("/api/analyze", handlers.Analyze)

	log.Println("🌐 Server running at http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
