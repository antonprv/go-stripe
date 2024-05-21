package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

const version = "1.0.0"
const cssVersion = "1."

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config 
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func (app * application) serve() {
	srv := &http.Server {
		Addr: fmt.Sprintf(":%d", app.config.port),
		Hanlder : app.routes(),
		IdleTimeout: 30* time.Second,
		ReadTimeout: 30* time.Second,
	}
}

func main() {
	var cfg config

	// flag.Int just produces an int.
	// flag.IntVar changes some predefined value with pointers.
	flag.IntVar(
		&cfg.port,
		"port",
		4000,
		"server port to listen on",
	)

	flag.StringVar(
		&cfg.env, "working environment",
		"development",
		"Application environment {development|production}",
	)

	flag.StringVar(
        &cfg.api,
        "api",
        "http://localhost:4001",
        "URL to api",
	)

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
        infoLog:       infoLog,
        errorLog:      errorLog,
        templateCache: tc,
        version:       version,
	}

}