package main

import (
	"flag"
	"html/template"
	"log"
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

func main() {
	var cfg config

	// flag.Int just produces an int.
	// flag.IntVar changes some predefined value with pointers.
	flag.IntVar(&cfg.port, )

}
