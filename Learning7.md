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
