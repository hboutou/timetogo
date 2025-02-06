package main

import (
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"time"
)

var fruits = [...]string{"🍇", "🍈", "🍉", "🍊", "🍋", "🍌", "🍍", "🥭", "🍎", "🍏", "🍐", "🍑", "🍒", "🍓", "🫐", "🥝", "🍅", "🥥"}

func home(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(len(fruits))
	fmt.Fprint(w, fruits[i])
}

func main() {
	port := flag.Int("port", 8080, "port to run the server on")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)

	rand.Seed(time.Now().UnixNano())

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)

	slog.Info("Starting server ", "address", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		slog.Error("oups!", err)
	}
}
