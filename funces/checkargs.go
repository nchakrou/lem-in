package funces

import (
	"log"
	"os"
	"strings"
)

// CheckArgs verifies that the program is run with exactly one argument,
// and that the argument is a .txt file. Otherwise, it exits the program.
func CheckArgs() {
	if len(os.Args) != 2 {
		log.Fatalln("<usage>: ./lem-in <file.txt>")
	}
	if !strings.HasSuffix(os.Args[1], ".txt") {
		log.Fatalln("invalid file format please enter a '.txt'file")
	}
}
