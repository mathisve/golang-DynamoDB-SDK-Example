package main

import "log"

// PrintAndGet is meant to reduce clutter in the main function
func PrintAndGet(id int) {
	p, err := GetItem(id)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v", p)
}
