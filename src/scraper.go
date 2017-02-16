package main

import (
	"log"
 	"github.com/PuerkitoBio/goquery"
)


// Scrapes BookFinder.com and reads information of a book
// Returns a book struct
func scrape(url string) Book {

  var book Book

  doc, err := goquery.NewDocument(url)
  if err != nil {
    log.Fatal(err)
  }

  book.url = url

  // Name
  doc.Find("span[itemprop='name']").Each(func(i int, s *goquery.Selection) {
    title := s.Text()
    book.title = title
  })

  // Author
  doc.Find("span[itemprop='author']").Each(func(i int, s *goquery.Selection) {
    author := s.Text()
    book.author = author
  })

  // ISBN
  doc.Find("span[itemprop='isbn']").Each(func(i int, s *goquery.Selection) {
    isbn := s.Text()
    book.isbn = isbn
  })

  // Prices
  doc.Find(".results-price a").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    price := s.Text()
    book.prices = append(book.prices, price)
  })

  return book
}
