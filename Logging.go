package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

// The Go standard library provides straightforward tools for outputting logs from Go programs, with the log package for free form output and log/slog pacakge for strucutred output.

func main() {
	// Simply invoking functions like Println from teh log package uses the standard logger, which is already preconfigured for reasonable logging output to os.Stderr.
	// Additional methods like Fatal* or Panic* will exit the program after logging.
	log.Println("standard logger")
	// 2024/08/21 14:34:24 standard logger

	// Loggers can be configured with flags to set their output format. By default, the standard logger has the log.Ldate and log.Ltime flags set, and these are collected in log.LstdFlags.
	// We can change its flags to emit time with microsecond accuracy, for example.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")
	// 2024/08/21 14:34:57.214256 with micro

	// It also supports emitting the file name and line from which the log function is called.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")
	// 2024/08/21 14:41:10 Logging.go:25: with file/line

	// It may be useful to create a custom logger and pass it around. When creating a new logger, we can set a prefix to distinguish output from other loggers.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")
	// my:2024/08/21 14:36:45 from mylog

	// We can set the prefix on existing loggers (including the standard one) with the SetPrefix method.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")
	// ohmy:2024/08/21 14:37:13 from mylog

	// Loggers can have custom output targets; any io.Writer works.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// This call writes the log output in buf.
	buflog.Println("hello")

	// This will actually show it on stadard output/
	fmt.Print("from bufflog:", buf.String())
	// from bufflog:buf:2024/08/21 14:42:18 hello

	// The slog package provides structured log output. For example, logging in JSON format is straightforward.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	// {"time":"2024-08-21T14:42:49.271193+05:30","level":"INFO","msg":"hi there"}

	// In addition to the message, slog output can contain an arbitrary number of key=value pairs.
	myslog.Info("hello again", "key", "val", "age", 25)
	// {"time":"2024-08-21T14:43:17.122919+05:30","level":"INFO","msg":"hello again","key":"val","age":25}
}
