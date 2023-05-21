package logger

import (
	"log"
	"os"
)

var (
	Info *log.Logger = log.New(
		os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Warning *log.Logger = log.New(
		os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Error *log.Logger = log.New(
		os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Critical *log.Logger = log.New(
		os.Stderr,
		"CRITICAL: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
)
