package main

import (
	"examples/models"
	"fmt"
)

func main() {

	myBook := models.Books{}
	myBook.Name = "Understanding Go in a visual way"
	myBook.Size.Width = 15
	myBook.Size.Height = 22
	fmt.Println(myBook)
}
