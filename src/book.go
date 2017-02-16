package main

type Book struct {

  url string
  title string
  author string
  isbn string
  prices []string

}

func (b Book) BestPrice() string {
    return b.prices[0] // TODO : Sort
}