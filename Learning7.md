### 71) Command Line Arguments:
```go
package main

import (
	"fmt"
	"os"
)

// Command Line args are a common way to parameterize execution of programs. For example, go run hello.go uses run and hello.go arguments to teh go program.

func mian() {
	// os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// go build CmdArg.go
// ./CmdArg a b c d
// [./CmdArg a b c d]
// [a b c d]
// c
```

### 72) Command Line Flags:
```go
package main

import (
	"flag"
	"fmt"
)

// Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag.

func main() {
	// Basic flag declarations are available for string, integer, and boolean options. Here we declare a string flag word with a default value "foo" and a short description.
	// This flag.String function returns a string pointer (not a string value); we’ll see how to use this pointer below.
	wordPtr := flag.String("word", "foo", "a string")

	// This declares numb and fork flags, using a similar approach to the word flag.
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	// Here we’ll just dump out the parsed options and any trailing positional arguments. Note that we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/*
To experiment with the command-line flags program it’s best to first compile it and then run the resulting binary directly.
$ go build command-line-flags.go

Try out the built program by first giving it values for all flags.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

Note that if you omit flags they automatically take their default values.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

Trailing positional arguments can be provided after any flags.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

Use -h or --help flags to get automatically generated help text for the command-line program.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

If you provide a flag that wasn’t specified to the flag package, the program will print an error message and show the help text again.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
*/
```

### 72) Command Line SubCommands:
```go
package main

import (
	"flag"
	"fmt"
	"os"
)

// Some command-line tools, like the go tool or git have many subcommands, each with its own set of flags. For example, go build and go get are two different subcommands of the go tool. The flag package lets us easily define simple subcommands that have their own flags.

func main() {
	// We declare a subcommand using the NewFlagSet function, and proceed to define new flags specific for this subcommand.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// For a different subcommand we can define different supported flags.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// The subcommand is expected as the first argument to the program.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Check which subcommand is invoked.
	switch os.Args[1] {
	// For every subcommand, we parse its own flags and have access to trailing positional arguments.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

/*
$ go build command-line-subcommands.go

First invoke the foo subcommand.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

Now try bar.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

But bar won’t accept foo’s flags.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
        level
*/
```

### 73) Environment Variables:
```go
package main

import (
	"fmt"
	"os"
	"strings"
)

// Environment variables are a universal mechanism for conveying configuration information to Unix programs. Let’s look at how to set, get, and list environment variables.

func main() {
	// To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv. This will return an empty string if the key isn’t present in the environment.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	// FOO: 1
	fmt.Println("BAR:", os.Getenv("BAR"))
	// BAR:

	// Use os.Environ to list all key/value pairs in the environment. This returns a slice of strings in the form KEY=value. You can strings.SplitN them to get the key and value. Here we print all the keys.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
	// TERM_PROGRAM
	// PATH
	// SHELL
	// ...
	// FOO
}
```

### 74) Logging:
```go
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
```

### 75) HTTP Client:
```go
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// The Go standard library comes with excellent support for HTTP clients and servers in the net/http package.
// In this example we'll use it to issue a simple HTTP requests.

func main() {
	// Issue an HTTP GET request to a server. http.Get is a convenient shortcut around creating an http.Client object and calling its Get method;
	// it uses the http.DefaultClient object which has useful default settings.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP respinse status
	fmt.Println("Response Status:", resp.Status)
	// Response Status: 200 OK

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	// <!DOCTYPE html>
	// <html>
	//   <head>
	//     <meta charset="utf-8">
	//     <title>Go by Example</title>
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
```

### 76) HTTP server
```go
package main

import (
	"fmt"
	"net/http"
)

// Writing a basic HTTP server is easy using the net/http pacakge.

// A fundamental concept in net/http servers is handlers. A handler is an object implementing the http.Handler interface.
// A common way to write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature.
func hello(w http.ResponseWriter, req *http.Request) {
	// Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments. The response writer is used to fill in the HTTP response.
	// Here our simple response is just “hello\n”.
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	// This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// We register our handlers on server routes using the http.HandleFunc convenience function. It sets up the default router in the net/http package and takes a function as an argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Finally, we call the ListenAndServe with the port and a handler. nil tells it to use the default router we’ve just set up.
	http.ListenAndServe(":8090", nil)
}
```

### 77) Context:
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

// In the previous example we looked at setting up a simple HTTP server. HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation.
// A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.

func hello(w http.ResponseWriter, req *http.Request) {
	// A context.Context is created for each request by the net/http machinary, and is available with the Context() method.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait for a few seconds before sending a reply to the client. This could simulate soem work the server is doing.
	// While working, keep an eye on the context's Done() channel for a signal that we should cancle the work adn return as soon as possible.
	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintf(w, "hellooo....\n")
	case <-ctx.Done():
		// The context's Err() method returns an error that explains why the Done() channel was closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// As before, we register our handler on the "/hello" route, and start serving.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
```

### 78) Spawning Process:
```go
package main

import (
	"fmt"
	"io"
	"os/exec"
)

// Sometimes our Go programs need to spawn other, non-Go processes.

func main() {
	// We'll start with a simple command that takes no arguments or input and just prints something to stdout.
	// The exec.Command helper creates an object to represent this external process.
	dateCmd := exec.Command("date")

	// The output method runs the command, waits for it to finish and collects its standard output.
	// If there were no errors, dateOut will hold bytes with hte date info.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	// > date
	// Wed Aug 21 17:27:40 IST 2024

	// Output and other methods of Command will return *execError if there was a problem executing the command(e.g. wrong path),
	// and *exec.ExitError if the command ran but exited with a non-zero return code.
	_, err = exec.Command("date", "-X").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}
	// command exit rc = 1

	// Next we'll look at a slightly more involved case where we pipe data to teh external process on its stdin and collect the result from its stdout.
	grepCmd := exec.Command("grep", "hello")

	// Here we explicitly grap input/output pipes, start the process, write some input to it, read the resulting output and finally wait for the process to exit.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// We ommitted error checks in the above example, but you could use the usual if err != nil pattern for all of them.
	// We also only collect the StdoutPipe results, but you could collect the SterrPipe in exactly the same way.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
	// > grep hello
	// hello grep

	// Note that when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string.
	// If you want to spawn a full command with a string, you can use bash's -c option:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOUt, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOUt))
	// > ls -a -l -h
	// total 1016
	// drwxr-xr-x@ 106 prashantsingh  staff   3.3K Aug 21 15:51 .
	// drwxr-xr-x    4 prashantsingh  staff   128B Aug 16 08:39 ..
	// drwxr-xr-x   15 prashantsingh  staff   480B Aug 21 15:48 .git
	// -rw-r--r--    1 prashantsingh  staff   1.4K Aug 16 07:37 Array.go
	// ...
	// -rw-r--r--    1 prashantsingh  staff    71B Aug 16 07:37 main.go
}
```

### 79) Executing Processes:
```go
package main

import (
	"os"
	"os/exec"
	"syscall"
)

// In the previouse example we looked at spawning external processes. We do this when we need and external process accesible to a running Go process.
// Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one. To do this we'll use Go's implementation of the classic exec function.

func main() {
	// For our example we'll exec ls. Go requires an absolute path to the binary we want to execute, so we'll use exec.LookPath to find it (probably /bin/ls).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in slice form (as opposed to one big string). We'll give ls a few common arguments. Note that first argument should be the program name.
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of Env variables to use. Here we just provide our current environment.
	env := os.Environ()

	// Here's the actual syscall.Exec call. If this call is successful, the execution of our process will end here and be replaced by the /bin/ls -a -l -h process.
	// If there is an error we'll get a return value.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	// total 1016
	// drwxr-xr-x@ 106 prashantsingh  staff   3.3K Aug 21 15:51 .
	// drwxr-xr-x    4 prashantsingh  staff   128B Aug 16 08:39 ..
	// drwxr-xr-x   15 prashantsingh  staff   480B Aug 21 15:48 .git
	// -rw-r--r--    1 prashantsingh  staff   1.4K Aug 16 07:37 Array.go
	// ...
	// -rw-r--r--    1 prashantsingh  staff    71B Aug 16 07:37 main.go
}
```

### 80) Signals:
```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Sometimes we'd like our Go programs to intelligently handle Unix Signales.
// For example, we might want a server to gracefully shutdown when it receives a SIGTERM, or a command-line tool to stop processing input if it receives a SIGINT.
// Here's how to handle signals in Go with Channels.

func main() {
	// Go signal notification works by sending os.Signal values on a channel.
	// We'll create a channel to receive these notifications. Note that thic channel should be buffered.
	sigs := make(chan os.Signal, 1)

	// signal.Notify registers the given channel to receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// We could receive from sigs here in the main function, but let's see how this could also be done in a sperate goroutine, to demonstrate a more realistic scenario of graceful shutdown.
	done := make(chan bool, 1)

	// This goroutine executes a blocking recevie for signals. When it gets one it'll print it out and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// The program will wait here until it gets the expected signal(as indicated by the goroutine above sending a value on done) and then exit.
	fmt.Println("await signal")
	<-done
	fmt.Println("exiting")

	// await signal
	// ^C
	// interrupt
	// exiting
}
```

### 81) Exit:
```go
package main

import (
	"fmt"
	"os"
)

// Use os.Exit to immediately exit with a given status.

func main() {
	// defers will not be run when using os.Exit, so this fmt.Println will never be called
	defer fmt.Println("!")

	// Exit with status 3.
	os.Exit(3)
}

// Note that unlike e.g. C, Go does not use an integer return value from main to indicate exit status. If you’d like to exit with a non-zero status you should use os.Exit.
```
