package main

import "time"

type ID string

func (i ID) New() string {
	return time.Now().Format("20060102150405")
}

func main() {
	var id ID
	println(id.New())
}
