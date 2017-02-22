package main

import (
)

// Main function
func main() {

  // Read urls file
  var urls = DaoReadBooksUrlsFile("./.booksUrls")

  // Init books array
  var books []Book;

  // Browse urls
  for _, url := range urls {

    // Update book local data if it is needed
    DfUpdateLocalBookIfNeeded(url)

    // Read local data and append it to books array
    var exists, localReading = DaoReadLocalBook(url)
    if exists {
      books = append(books, localReading)
    }
   
  }

  // Displays infos for BitBar
  BitbarPrint(books)
}