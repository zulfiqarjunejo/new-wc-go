package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	c := flag.Bool("c", false, "count number of bytes")
	l := flag.Bool("l", false, "count number of lines")

	flag.Parse()

	filename := flag.Arg(0)

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read %s: %s\n", filename, err.Error())
	}

	if *c {
		fmt.Printf("  %d %s\n", len(content), filename)
	} else if *l {
		lines := strings.Split(string(content), "\n")
		fmt.Printf("  %d %s\n", len(lines)-1, filename)
	}
}
