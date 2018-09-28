package main

import (
	"net/http/"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", Write([]byte(name)))
	http.ListenAndServe(port(), name)
	fmt.Println(name)
}

func port() string {
        port := os.Getenv("PORT")
        if len(port) == 0 {
                port = "80"
        }
        return ":" + port
}

