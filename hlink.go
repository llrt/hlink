package main

import (
	"flag"
	"log"
	"os"
)

func main() {
    flag.Parse()

    // receives arguments from command line
	args := flag.Args()

    // check if there are exactly two arguments (source and destination of hard link)
	if len(args) != 2 {
        log.Println(args)
		log.Fatal("Link script must receive exactly two arguments: source and destination")
	}

	src := args[0]
	dst := args[1]

    // make a direct call to the syscall of current OS
	if err := os.Link(src, dst); err != nil {
        // if there is an error, prints error message and exits
        log.Println("An error occurred while trying to create requested link")
		log.Fatal(err)
	}

	log.Println("Successfully created requested link")
}
