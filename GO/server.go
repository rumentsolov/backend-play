package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rumentsolov/GoLangWeb/config"
	"github.com/rumentsolov/GoLangWeb/pkg/handlers"
)

var portNumber = ":8002"
var app config.AppConfig
var session *scs.SessionManager 

func main() {
	
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cach")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app) 
	handlers.NewHandler(repo)

	render.NewTemplate(&app) 
	fmt.Printf("Starting app on port %s \n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

	//_ = http.ListenAndServe(portNumber, nil)
}
