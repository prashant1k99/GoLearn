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
