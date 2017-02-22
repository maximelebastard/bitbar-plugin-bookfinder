package main

import(
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
	"regexp"
	"os"
	"time"
	"bufio"
)

// Reads the book urls file and returns them in an array
func DaoReadBooksUrlsFile(uri string) []string {

	filename := DaoGetBooksUrlsFilePath()

	content, err := ioutil.ReadFile(filename)
	deal(err)

	lines := strings.Split(string(content), "\n")

	return lines

}

// Reads local data of a book if exists
func DaoReadLocalBook(url string) (exists bool, book Book) {

	pathes := DaoGetBookLocalDataDirPathes(url)

	exists = true
	if _, err := os.Stat(pathes.dir); 
	os.IsNotExist(err) {
		exists = false
	}

	if(exists){
		book.url = DaoReadDataFile(pathes.url)
		book.isbn = DaoReadDataFile(pathes.isbn)
		book.title = DaoReadDataFile(pathes.title)
		book.author = DaoReadDataFile(pathes.author)
		book.bestPrice = DaoReadDataFile(pathes.bestPrice)
		book.bestPriceBefore = DaoReadDataFile(pathes.bestPriceBefore)
		book.lastUpdate, _ = strconv.Atoi(DaoReadDataFile(pathes.lastUpdate))
	}


	return;
}

func DaoReadDataFile(path string) string {
	content, err := ioutil.ReadFile(path)
	deal(err)

	strContent := string(content)

	re := regexp.MustCompile(`\r?\n`)
	cleanContent := strings.Trim(re.ReplaceAllString(strContent, ""), " ")

	return cleanContent
}

func DaoWriteDataFile(path string, data string) {

	var err = ioutil.WriteFile(path, []byte(data), 0664)
    deal(err)

}

// Writes local data of a book
func DaoWriteLocalBook(book Book) {

	fmt.Printf("%+v\n", book)

	pathes := DaoGetBookLocalDataDirPathes(book.url)

	_ = os.Mkdir(pathes.dir, 0664) // TODO Fix these rights !

	DaoWriteDataFile(pathes.url, book.url)
	DaoWriteDataFile(pathes.isbn, book.isbn)
	DaoWriteDataFile(pathes.title, book.title)
	DaoWriteDataFile(pathes.author, book.author)
	DaoWriteDataFile(pathes.bestPrice, book.bestPrice)
	DaoWriteDataFile(pathes.bestPriceBefore, book.bestPriceBefore)
	DaoWriteDataFile(pathes.lastUpdate, strconv.Itoa(book.lastUpdate))


}

// Register a distant request
func DaoRegisterDistantFetch() {

	path := DaoGetDistantFetchesFilePath()

	// Create file if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err := os.Create(path)
		deal(err)
	}

	// Open file
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0664)
	deal(err)

	defer f.Close()

	// write in it
	var line string = strconv.Itoa(int(time.Now().Unix())) + "\r\n"
	_, err = f.WriteString(line)
	deal(err)

}

// Counts the amount of requests made in the past hour
// This is to avoid to make more than 10 requests per hour
// (when you have a lot of books)
func DaoGetDistantFetchesAmountInThePastHour() int {
		
		// TODO : Optimize this function to delete too old lines

	path := DaoGetDistantFetchesFilePath()
	nbFetches := 0

	// The limit is now - 1 hour
	limit := int(time.Now().Unix()) - 60*60

	file, err := os.Open(path)
    deal(err)

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        var lineValue, err = strconv.Atoi( scanner.Text() )

        if(err != nil) {
		    panic(err)
		}

        if( lineValue >= limit ) {
        	nbFetches++
        }

    }

    err = scanner.Err()
    deal(err)

    return nbFetches
}

// Gets the absolute path to the books url file
func DaoGetBooksUrlsFilePath() string {
	return "./.bitbar-plugin-bookfinder/books-urls.txt";
}

// Gets the absolute path to the fetch register
func DaoGetDistantFetchesFilePath() string {
	return "./.bitbar-plugin-bookfinder/fetches.txt";
}

// Gets the absolute path to a book local data directory
func DaoGetBookLocalDataDirPathes(url string) LocalDataPathes {


	dirPath := "./.bitbar-plugin-bookfinder/booksdata/"+md5string(url) //TODO optimize concat

	var pathes LocalDataPathes;
	pathes.dir = dirPath
	pathes.url = dirPath+"/url"
	pathes.isbn = dirPath+"/isbn"
	pathes.title = dirPath+"/title"
	pathes.author = dirPath+"/author"
	pathes.bestPrice = dirPath+"/bestPrice"
	pathes.bestPriceBefore = dirPath+"/bestPriceBefore"
	pathes.lastUpdate = dirPath+"/lastUpdate"


	return pathes;
}


// Represents a book data local directory pathes
type LocalDataPathes struct {
		dir string
		url string
		isbn string
		title string
		author string
		bestPrice string
		bestPriceBefore string
		lastUpdate string
	}