package funces

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ReadFile reads the input file line by line, trims spaces,
// checks that there is exactly one ##start, one ##end, and
// at least one valid link, then returns all lines as a slice of strings.
func ReadFile() []string {
	Lines := []string{}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	IsStart := false
	IsEnd := false
	IsLink := false
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "##start" {
			if IsStart {
				log.Fatalln("ERROR: invalid data format, duplicate start")
			}
			IsStart = true
		} else if text == "##end" {
			if IsEnd {
				log.Fatalln("ERROR: invalid data format, duplicate end")
			}
			IsEnd = true
		} else if strings.Contains(text, "-") {
			part := strings.Split(text, "-")
			if len(part) == 2 && !strings.Contains(text, " ") {

				IsLink = true
			}

		}
		Lines = append(Lines, text)
	}
	if !IsStart {
		log.Fatalln("ERROR: invalid data format, no start room found")
	} else if !IsEnd {
		log.Fatalln("ERROR: invalid data format, no end room found")
	} else if !IsLink {
		log.Fatalln("ERROR: invalid data format, no Links  found")
	}
	return Lines
}
