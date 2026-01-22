package handlers

import (
	"encoding/json"
	"net/http"

	"ai-autonomous-redteam/api/types"
	"ai-autonomous-redteam/engine"
)

// Analyze handles POST /api/analyze
func Analyze(w http.ResponseWriter, r *http.Request) {
	// 1️⃣ Allow only POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2️⃣ Decode request body
	var req types.AnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.Target == "" {
		http.Error(w, "target is required", http.StatusBadRequest)
		return
	}

	// 3️⃣ Run engine
	report, err := engine.Run(req.Target)
	if err != nil {
		// IMPORTANT: return JSON error so frontend can display it
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// 4️⃣ Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
