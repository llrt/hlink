package main

import (
	"flag"
    "fmt"
	"log"
	"os"
)

func main() {
    unln := flag.Bool("u", false, "Unlinks the specified destination path")

    flag.Parse()

    // receives arguments from command line
	args := flag.Args()


    if *unln { // if -u flag was provided, do unlink destination
        // check if there is exactly one arguments (destination of hard link)
        if len(args) != 1 {
            log.Println(args)
            log.Fatal("Link script must receive exactly one argument when unlinking: destination")
        }

        dst := args[0] // destination

        yn := "n"
        fmt.Println("WARNING: removing hard link does delete all structure on destination. Confirm removal? [y|n]")
        fmt.Scanf("%s", &yn)

        if yn == "y"{ // removal confirmed: proceed with removal
            log.Println("Removal confirmed. Hard link will be removed")
            // make a direct call to the syscall of current OS
            if err := os.Remove(dst); err != nil {
                // if there is an error, prints error message and exits
                log.Println("An error occurred while trying to remove specified link")
                log.Fatal(err)
            }

            log.Println("Successfully removed specified link")
        } else { // removal cancelled
            log.Println("Removal cancelled. Nothing done")
        }

    } else { // otherwise, do link source to destination
        // check if there are exactly two arguments (source and destination of hard link)
        if len(args) != 2 {
            log.Println(args)
            log.Fatal("Link script must receive exactly two arguments when linking: source and destination")
        }

        src := args[0] // source
        dst := args[1] // destination

        // make a direct call to the syscall of current OS
        if err := os.Link(src, dst); err != nil {
            // if there is an error, prints error message and exits
            log.Println("An error occurred while trying to create specified link")
            log.Fatal(err)
        }

        log.Println("Successfully created specified link")
    }

}
