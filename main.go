package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"zauth/api"
	"zauth/sessionmanager"
	"zauth/strconv"

	"github.com/gorilla/mux"
)

func main() {
	// router
	r := mux.NewRouter()
	// context
	ctx := context.Background()

	// service instantiation
	svc := strconv.StringService{}
	sess := sessionmanager.SessionService{}

	// api registration
	api.RegisterStrconvRoutes(svc, r)
	api.RegisterSessionRoutes(sess, ctx, r)

	// logging
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(met, tpl, err1, err2)
		return nil
	})

	// launch server
	fmt.Println("Started Server http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
