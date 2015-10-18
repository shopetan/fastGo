package main

import (
	"encoding/json"
	"fmt"

	"github.com/codegangsta/negroni"
	"github.com/shopetan/fastGo/models"
	"github.com/unrolled/render"
	"net/http"
)

func main() {
	ren := render.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		user := models.NewUser("TANAKA tarou", 10)
		marshalUser, _ := json.Marshal(user)
		ren.JSON(w, http.StatusOK, user)
	})
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
