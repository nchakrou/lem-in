package funces

import (
	"log"
	"strconv"
	"strings"
)

func Validation(Lines []string) {
	var Start Room
	var End Room
	IsStarted := false
	IsinEnd := false
	IsLink := false
	rooms := make(map[string]bool)
	coords := make(map[string]bool)
	links := make(map[string]bool)
	NumAnts, err := strconv.Atoi(Lines[0])
	if err != nil || NumAnts <= 0 {
		log.Fatalln("ERROR: invalid data format, invalid number of Ants")
	}

	for i := 1; i < len(Lines); i++ {
		Line := Lines[i]
		Split := []string{}
		if strings.Contains(Line, " ") {
			if !strings.HasPrefix(Line, "#") {

				Split = strings.Split(Line, " ")
				if len(Split) == 3 && !strings.HasPrefix(Split[0], "L") && !strings.HasPrefix(Split[0], "l") {
					XandY(Split)
					if coords[Split[1]+" "+Split[2]] {
						log.Fatalln("ERROR: invalid data format. repeat coords")
					}
					coords[Split[1]+" "+Split[2]] = true
					if !rooms[Split[0]] {
						rooms[Split[0]] = true
					} else {
						log.Fatalln("ERROR: invalid data format. repeat Room")
					}
				} else {
					log.Fatalln("ERROR: invalid data format. invalid Room")
				}

			}
		} else if strings.Contains(Line, "-") {
			if !strings.HasPrefix(Line, "#") {
				Split = strings.Split(Line, "-")
				if len(Split) == 2 {
					if strings.HasPrefix(Split[0], "L") || strings.HasPrefix(Split[1], "L") || strings.HasPrefix(Split[1], "#") {
						log.Fatalln("ERROR: invalid data format, invalid Links")
					}
					IsLink = true
				}

			}
		}
		if Line == "##start" {
			if IsLink {
				log.Fatalln("ERROR: invalid data format, invalid Links")
			}

			IsStarted = true
			continue
		} else if Line == "##end" {
			if IsLink {
				log.Fatalln("ERROR: invalid data format, invalid Links")
			}
			IsinEnd = true
		}
		if IsStarted && len(Split) == 3 {
			Start.Name, Start.x, Start.y = XandY(Split)
			IsStarted = false
		}
		if IsinEnd && len(Split) == 3 {
			End.Name, End.x, End.y = XandY(Split)
			IsinEnd = false
		}
		if IsLink {
			if Split[0] == Split[1] {
				log.Fatalln("ERROR: invalid data format. to the same room", Line)
			} else if !rooms[Split[0]] || !rooms[Split[1]] {
				log.Fatal("ERROR: invalid data format. unexisting room")
			} else if links[Split[0]+"-"+Split[1]] {
				log.Fatal("ERROR: invalid data format. double Links")
			} else if links[Split[1]+"-"+Split[0]] {
				log.Fatal("ERROR: invalid data format. double Links")
			}
			links[Line] = true
		}
	}
}
func XandY(Split []string) (string, int, int) {
	var err error
	Name := Split[0]
	x := 0
	y := 0
	if x, err = strconv.Atoi(Split[1]); err != nil {
		log.Fatalln("ERROR: invalid data format, invalid x")
	}
	if y, err = strconv.Atoi(Split[2]); err != nil {
		log.Fatalln("ERROR: invalid data format, invalid y")
	}
	return Name, x, y
}
