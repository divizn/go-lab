package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/divizn/go-lab/pkg/utils"
)

func main() {
	staticDir, err := utils.JoinWithCWD("cmd/web-server/static")
	if err != nil {
		panic(err)
	}

	fmt.Println(staticDir)
	fmt.Println(os.ReadDir(staticDir))

	// API endpoint
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello from API!"})
	})

	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
