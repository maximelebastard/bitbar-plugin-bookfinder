package main

import (
	"fmt"
)

func BitbarPrint(books []Book){
	BitbarPrintMenuIcon()
	BitbarPrintBooks(books)
	BitbarPrintUrlsFileEdition()
}

// Prints books list in Bitbar format (more at https://github.com/matryer/bitbar#plugin-api)
func BitbarPrintBooks(books []Book) {

	for _, book := range books {
		fmt.Println("---")
		fmt.Printf("%s : %s |href=%s \n", book.DisplayTitle(), book.DisplayPrice(), book.url)
	}
	
}

// Prints the menu icon displayed in status bar
func BitbarPrintMenuIcon() {
	fmt.Println("BF")

}

// Prints the edit books file item
func BitbarPrintUrlsFileEdition() {
    fmt.Println("---")
    fmt.Println("Edit books file...")
}