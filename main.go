package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	runtime.GOMAXPROCS(1)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Please use \"/loadtest/[0-999]\" for loadtesting.\n"))
	})

	for i := 0; i <= 999; i++ {
		router.HandleFunc(fmt.Sprintf("/loadtest/%d", i), func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Millisecond * 100)
			w.Write([]byte("Greetings from \"" + r.URL.Path + "\"\n"))
		})
	}

	http.Handle("/", router)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
