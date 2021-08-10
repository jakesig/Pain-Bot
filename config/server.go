package config

import (
  "fmt"
  "net/http"
)

func Server() {
  http.HandleFunc("/", HelloServer)
  http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Bot is alive!")
}