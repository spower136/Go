package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type book struct {
	Isbn       string   "csv:'isbn13'"
	Title      string   "csv:'title'"
	Last_Name  string   "csv:'authorLN'"
	First_Name string   "csv:'authorFN'"
	Subjects   []string "csv:'subjects'"
	Numsub     int      "csv:'numb of subjects'"
	Format     string   "csv:'format'"
	Price      string   "csv:'price'"
	Publisher  string   "csv:'publisher'"
	Pubdate    string   "csv:'pubdate'"
}

func readme() {
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
		fmt.Printf("%+v\n", rec)

		b := rowToBook(rec)
		books = append(books, b)

	}

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

type Bk interface {
	getAuthorLName() string
	getPrice() float64
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

func GetMaxCol[T ~string | ~float64](f func(Book) T, books []Bk) T {
	max := f(books[0])
	for _, b := range books {
		if f(b) > max {
			max = f(b)
		}
	}
	return max
}

func main() {
	bk := readme()
	la := GetMaxCol(getAuthorLName, bk)
	hp := GetMaxCol(getPrice, bk)
	fmt.Printf("Last author alphabetically: %v\n", la)
	fmt.Printf("Highest priced book: $%v \n", hp)

}
