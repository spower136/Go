package assign_5_functions

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
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

func Readme() []Book {
	// open file
	f, err := os.Open("books_go.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	books := make([]Book, 0)

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

		b := RowToBook(rec)
		books = append(books, b)

	}
	return books[1:]
}

func RowToBook(row []string) Book {

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

type SortedBooksLN []Book

func (a SortedBooksLN) Len() int {
	return len(a)
}

func (a SortedBooksLN) Less(i, j int) bool {
	return a[i].Last_Name < a[j].Last_Name
}

func (a SortedBooksLN) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type SortedBooksPrice []Book

func (a SortedBooksPrice) Len() int {
	return len(a)
}

func (a SortedBooksPrice) Less(i, j int) bool {
	return a[i].Price < a[j].Price
}

func (a SortedBooksPrice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//"getter” function that returns the last name of the book's author
func GetAuthorLName(b Book) string {
	return b.Last_Name
}

// "getter” function that returns a book's price
func GetPrice(b Book) float64 {
	v, _ := strconv.ParseFloat(b.Price[1:], 64)
	return v
}

func GetMaxCol[T ~string | ~float64](f func(Book) T, books []Book) T {
	max := f(books[0])
	for _, b := range books {
		if f(b) > max {
			max = f(b)
		}
	}
	return max
}
