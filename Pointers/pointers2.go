package main

/*
	Each time we take the address of a variable or copy a pointer, we create new aliases or ways
	to identify the same variable.For example, *p is an alias for v. Pointer aliasing is useful because it allows
	us to access a variable without using its name, but this is a double-edged sword: to find al the statements that
	access a variable, we have to know all its aliases. It's not just pointers that create aliases; aliasing also
	occurs when we copy values of other reference types like slices, maps, and channels, andeven structs, arrays, and
	interfaces that contain these types.

	pointers are key to the flag package, which uses a program's command-line arguents to set the value of certain
	variables distributed throughout the program.
*/

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")

/*
	the function flag.Bool creates a new flag variable of type bool. It takes three arguments: the name of the flag
	("n"), the variable's default value (false), and a message that will be printed if the user provides invalid argument,
	an invalid flag, or -h or -help.
*/

var sep = flag.String("s", " ", "separator")

/*
	Similarly, flag.String takes a name, a default value, and a message, and creates a string variable, which must be accessed
	indirectly as *sep and *n.
*/

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

/*
	When the program runs, it must call flag.Parse before the flags are used, to update the flag variables from their default values. the non-flag argument
	are available from flag.Args() as a slice of strings. If flag.Parase encounterss an error,it prints a usage message
	and calls os.Exit(2) to terminate the program.

	Some test cases for the program
	$ go build pointers2.go
	$ ./pointers2 a bc def
	$ ./pointers2 -s / a bc def
	$ ./pointers2 -n a bc def
	$ ./pointers2 -help


*/
