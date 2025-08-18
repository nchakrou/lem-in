package main

import (
	lemin "lemin/funces"
)

func main() {
	lemin.CheckArgs()
	Lines := lemin.ReadFile()
	Start, End, Graph := lemin.Validation(Lines)
	lemin.Bfs(Start, End, Graph)

}
