package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

var fruits = [...]string{"🍇", "🍈", "🍉", "🍊", "🍋", "🍌", "🍍", "🥭", "🍎", "🍏", "🍐", "🍑", "🍒", "🍓", "🫐", "🥝", "🍅", "🥥"}

func handleFruits(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		fruit := fruits[id%len(fruits)]
		fmt.Fprintf(w, "%s", fruit)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Get some fruits at /fruits"))
}

func main() {
	address := ":5000"
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/fruits", handleFruits)

	slog.Info("starting server.", "address", address)
	err := http.ListenAndServe(address, mux)
	if err != nil {
		slog.Error("oops!", err)
	}
}
