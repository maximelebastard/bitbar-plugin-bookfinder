package main

type Book struct {

  url string
  isbn string
  title string
  author string
  bestPrice string
  bestPriceBefore string
  lastUpdate int

}

/*func (b Book) BestPrice() string {

	var sortedPrices = b.prices
	sort.Sort(SortByPrice(sortedPrices))

	fmt.Printf("%v", sortedPrices)

	if 0<len(sortedPrices) {
		return sortedPrices[0]
	} else {
		return "none"
	}
    
}*/

func (book Book) DisplayTitle() string {

	var res string

	if(book.title != "") {
		res += book.title

		if(book.author != "") {
			res += " - " + book.author
		}
	} else if (book.isbn != "") {
		res += "ISBN " + book.isbn
	} else {
		res += "No title or ISBN found"
	}

	return res
}

func (book Book) DisplayPrice() string {

	var res string

	if(book.bestPrice != ""){
		res += book.bestPrice
	} else {
		res += "No price found"
	}

	return res
}

func (b Book) BestPrice() string {

	return "none"
    
}

func (b Book) hashedUrl() string {

	return md5string(b.url)

}