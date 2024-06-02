package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// For export to javascript.
const version = "1.0.0"
const cssVersion = "1."

// Basic configuration.
type config struct {
	port int
	env  string
	api  string

	// dsn - Data Source Name.
	// Includes the database type, server address,
	// database name, and authentication credentials.
	db struct {
		dsn string
	}

	// Stripe configuration parameters.
	stripe struct {
		secret string
		key    string
	}
}

// Define the basic structure of the application.
type application struct {
	config config

	// Logging everything that happens with the application.
	infoLog  *log.Logger
	errorLog *log.Logger

	// Caching our server's most popular requests.
	// This might require Redis.
	templateCache map[string]*template.Template

	// Changed with the const ot the top of the file.
	version string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf(
		"Starting HTTP server in %s mode on port %d\n",
		app.config.env,
		app.config.port,
	)

	return srv.ListenAndServe()

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
		&cfg.env, "environ",
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

	// "INFO\t" and the like is the prefix for each logging message.
	// Position of log.Ldate and log.Ltime is irrelevant. They are displayed
	// all the same in the end: <date> <time>.
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

	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}

}
