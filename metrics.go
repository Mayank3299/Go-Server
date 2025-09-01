package main

import (
	"fmt"
	"net/http"
)

func (ac *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	resp := fmt.Sprintf(`<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>`, ac.fileserverHits.Load())
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
