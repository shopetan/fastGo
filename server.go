package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/shopetan/fastGo/models" // 自分のパッケージ
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		user := models.NewUser("Naoyoshi Aikawa", 29)
		fmt.Fprintln(w, user)
	})
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
