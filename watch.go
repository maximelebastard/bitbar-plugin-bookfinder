package main

import (
	"fmt"
	"log"
  "github.com/PuerkitoBio/goquery"
)

type Book struct {

  url string
  title string
  author string
  isbn string
}

func main() {
	var book = scrape("http://www.bookfinder.com/search/?author=&title=&lang=en&isbn=978-0134682334&new_used=*&destination=fr&currency=EUR&mode=basic&st=sr&ac=qr")

  fmt.Printf("%+v\n", book)
}

func scrape(url string) Book {

  var book Book

  doc, err := goquery.NewDocument(url)
  if err != nil {
    log.Fatal(err)
  }

  // Set book url
  book.url = url

  // Find the book name
  doc.Find("span[itemprop='name']").Each(func(i int, s *goquery.Selection) {
    title := s.Text()
    fmt.Printf("Title : %s \n", title)
    book.title = title
  })

  // Find the book author
  doc.Find("span[itemprop='author']").Each(func(i int, s *goquery.Selection) {
    author := s.Text()
    fmt.Printf("Author: %s \n", author)
    book.author = author
  })

  // Find the book ISBN
  doc.Find("span[itemprop='isbn']").Each(func(i int, s *goquery.Selection) {
    isbn := s.Text()
    fmt.Printf("ISBN: %s \n", isbn)
    book.isbn = isbn
  })

  // Find the book prices
  doc.Find(".results-price a").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    price := s.Text()
    fmt.Printf("Price %d: %s\n", i, price)
  })

  return book
}