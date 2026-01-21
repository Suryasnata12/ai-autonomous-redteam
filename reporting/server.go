package reporting

import (
	"encoding/json"
	"log"
	"net/http"

	"ai-autonomous-redteam/reporting/model"
)

func StartServer(report model.Report) {
	// ---- API endpoint: JSON report ----
	http.HandleFunc("/report", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(report); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// ---- Static frontend ----
	http.Handle("/", http.FileServer(http.Dir("./web")))

	// ---- Startup log (important for Windows) ----
	log.Println("🚀 Report server running at http://localhost:8080")
	log.Println("📄 JSON report available at http://localhost:8080/report")

	// ---- Start server (blocks forever) ----
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("server failed:", err)
	}
}
