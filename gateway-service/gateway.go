package main

import (
	"net/http"
)

func main() {
	//"/" matches anything to route map
	//"/abc" will only match "/abc" and not any other path in the route
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I am working and healthy"))
	})

	panic(http.ListenAndServe(":80", nil))

}
