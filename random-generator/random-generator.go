package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func getRandomNumber(limit int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(limit) + 1
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div><form action=\"random\">Informe um valor máximo para sorteio: <input type=text name=\"num\" /></form></div>")
}

func handlerRandom(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["num"]) == 0 {
		http.Error(w, "Not found number in param, try to add it in your next request '?num=X'", http.StatusBadRequest)
	} else {
		value, err := strconv.Atoi(r.Form["num"][0])
		if err != nil || value == 0 {
			http.Error(w, "Invalid number", http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "<h1>%v</h1>", getRandomNumber(value))
		}
	}
}

func main() {
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/random", handlerRandom)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
