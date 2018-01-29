package main

import (
	"log"
	"web/lib"
)

func main() {
	log.Println("Weclome to Avalon")
	lib.LogInit(true)
	lib.Web()
}
