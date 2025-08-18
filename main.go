package main

import (
	lemin "lemin/funces"
)

func main() {
	lemin.CheckArgs()
	Lines := lemin.ReadFile()
	lemin.Validation(Lines)

}
