package logging

import (
	"log"
	"os"
)

var (
	Request  *log.Logger
	Info     *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
	Critical *log.Logger
)

func init() {
	Request = log.New(os.Stdout, "", log.LstdFlags)
	Info = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.LstdFlags|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
	Critical = log.New(os.Stderr, "CRITICAL: ", log.LstdFlags|log.Lshortfile)
}
