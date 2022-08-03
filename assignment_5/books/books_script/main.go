package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type book struct {
	Isbn       string   `csv:"isbn13"`
	Title      string   `csv:"title"`
	Last_Name  string   `csv:"authorLN"`
	First_Name string   `csv:"authorFN"`
	Subjects   []string `csv:"subjects"`
	Numsub     int      `csv:"numb of subjects"`
	Format     string   `csv:"format"`
	Price      string   `csv:"price"`
	Publisher  string   `csv:"publisher"`
	Pubdate    string   `csv:"pubdate"`
}

func readme() []book {
	// open file
	f, err := os.Open("books_go.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	books := make([]book, 0)

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		// fmt.Printf("%+v\n", rec)

		b := rowToBook(rec)
		books = append(books, b)

	}
	return books[1:]
}

func rowToBook(row []string) book {

	sub := strings.Split(row[4], ",")
	num, _ := strconv.Atoi(row[5])

	book := book{}
	book.Isbn = row[0]
	book.Title = row[1]
	book.Last_Name = row[2]
	book.First_Name = row[3]
	book.Subjects = sub
	book.Numsub = num
	book.Format = row[6]
	book.Price = row[7]
	book.Publisher = row[8]
	book.Pubdate = row[9]
	return book
}

type sortedBooksLN []book

func (a sortedBooksLN) Len() int {
	return len(a)
}

func (a sortedBooksLN) Less(i, j int) bool {
	return a[i].Last_Name < a[j].Last_Name
}

func (a sortedBooksLN) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type sortedBooksPrice []book

func (a sortedBooksPrice) Len() int {
	return len(a)
}

func (a sortedBooksPrice) Less(i, j int) bool {
	return a[i].Price < a[j].Price
}

func (a sortedBooksPrice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//"getter” function that returns the last name of the book's author
func getAuthorLName(b book) string {
	return b.Last_Name
}

// "getter” function that returns a book's price
func getPrice(b book) float64 {
	v, _ := strconv.ParseFloat(b.Price[1:], 64)
	return v
}

func GetMaxCol[T ~string | ~float64](f func(book) T, books []book) T {
	max := f(books[0])
	for _, b := range books {
		if f(b) > max {
			max = f(b)
		}
	}
	return max
}

func main() {
	books := readme()
	sort.Sort(sortedBooksLN(books))
	for _, b := range books {
		fmt.Printf("%+v\n", b.Last_Name)
	}
	sort.Sort(sortedBooksPrice(books))
	for _, b := range books {
		fmt.Printf("%+v\n", b.Price)
	}

	fmt.Printf("%+v\n", books)
	la := GetMaxCol(getAuthorLName, books)
	hp := GetMaxCol(getPrice, books)
	fmt.Printf("Last author alphabetically: %v\n", la)
	fmt.Printf("Highest priced book: $%v \n", hp)
}
