package main

import (
	//"os"
 	"github.com/PuerkitoBio/goquery"
 	"time"
 	"sort"
 	"fmt"
 	"strings"
)

// Updates a local book from distant data if it is needed
func DfUpdateLocalBookIfNeeded(url string){

	var exists, localBook = DaoReadLocalBook(url)

	if(!exists || DfIsUpdateNeeded(localBook)) {

		bestPriceBefore := localBook.bestPrice

		var _, distantBook = DfFetchDistantData(url)

		distantBook.bestPriceBefore = bestPriceBefore

		DaoWriteLocalBook(distantBook)

	}

}

// Defines if the update of a book is needed
// *needed = Has not been done in the last 12 hours
func DfIsUpdateNeeded(book Book) bool {

	// IMPORTANT NOTICE
	// 	I strongly discourage you to edit the behavior of
	// 	this function.
	// 	BookFinder monitors really tightly the requests
	// 	and bans every client that do not behave like
	// 	a regular webbrowser visitor.
	// 	Their conditions do not allow webscraping, so stay discreet.
	// 	I got banned for a few days because of making more than
	// 	a dozen price requests in an hour.
	// 	I got confirmation from their team that most of the prices do not vary more often
	// 	than every 12 hours. So no need to refresh more frequently.
	twelveHoursAgo := int(time.Now().Unix()) - 60*60*12
	return (book.lastUpdate < twelveHoursAgo && DaoGetDistantFetchesAmountInThePastHour() <= 10)

}



// Fetches data from BookFinder for a given search url
func DfFetchDistantData(url string) (err error, book Book){

println("fff")

	DaoRegisterDistantFetch()

	doc, err := goquery.NewDocument(url)
	deal(err)

	fmt.Printf("%+v\n", doc)

	book.url = url

	// Name
	doc.Find("span[itemprop='name']").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		book.title = title
	})

	// Author
	var authors []string
	doc.Find("span[itemprop='author']").Each(func(i int, s *goquery.Selection) {
		authors = append(authors, s.Text())
	})
	book.author = strings.Join(authors, ", ")

	// ISBN
	doc.Find("span[itemprop='isbn']").Each(func(i int, s *goquery.Selection) {
		isbn := s.Text()
		book.isbn = isbn
	})

	// Prices
	var prices []string;
	doc.Find(".results-price a").Each(func(i int, s *goquery.Selection) {
	// For each item found, get the band and title
		price := s.Text()
		prices = append(prices, price)
	})
	sort.Sort(SortByPrice(prices))

	if(len(prices) > 0){
		book.bestPrice = prices[0]
	} else {
		book.bestPrice = ""
	}
	

	book.lastUpdate = int(time.Now().Unix())

	return;

}

