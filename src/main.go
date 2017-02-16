package main

import (
	"fmt"
)

// Main function
func main() {

  var books []Book

  // Read urls from file
  // TODO
  var urls = [3]string{"http://www.bookfinder.com/search/?author=&title=&lang=en&isbn=978-0134682334&new_used=*&destination=fr&currency=EUR&mode=basic&st=sr&ac=qr","http://www.bookfinder.com/search/?author=&title=&lang=en&isbn=9780321767530&new_used=*&destination=fr&currency=EUR&mode=basic&st=sr&ac=qr","http://www.bookfinder.com/search/?keywords=+978-1780671840&new=&used=&ebooks=&classic=&lang=en&st=sh&ac=qr&submit="}
	
  // Scrape BookFinder.com to get informations
  for _, url := range urls {
    books = append(books, scrape(url))
  }

  // Displays infos for BitBar
  bitbarPrint(books)
}

// bitbarPrint formats books array for BitBar app
// and prints it
func bitbarPrint(books []Book) {
  fmt.Printf("BF \n---\n")
  for _, book := range books {
    fmt.Printf("%s - %s : %s |href=%s \n", book.title, book.author, book.BestPrice(), book.url)
  }

}