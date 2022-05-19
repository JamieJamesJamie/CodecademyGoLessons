package main

import (
	"fmt"
	"time"
)

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber uint
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title,
		"was written by", writer,
		"and drawn by", artist,
		"and published by", publisher,
		"in", year,
		"with", pageNumber, "pages.",
		"The grade is", grade, "according to submitted reviews.",
		"The cost of the comic book is",
		(float32(pageNumber*(uint(time.Now().Year())-year+1))*grade)/100, "pounds.",
	)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title,
		"was written by", writer,
		"and drawn by", artist,
		"and published by", publisher,
		"in", year,
		"with", pageNumber, "pages.",
		"The grade is", grade, "according to submitted reviews.",
		"The cost of the comic book is",
		(float32(pageNumber*(uint(time.Now().Year())-year+1))*grade)/100, "pounds.",
	)
}
