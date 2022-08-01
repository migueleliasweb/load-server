package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/migueleliasweb/load-server/pkg/limiter"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Please use \"/loadtest/[0-999]\" for loadtesting.\n"))
	})

	for i := 0; i <= 999; i++ {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Greetings from \"" + r.URL.Path + "\"\n"))
		})

		router.Handle(fmt.Sprintf("/loadtest/%d", i), limiter.AddLimiter(100, handler))
	}

	http.Handle("/", router)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
