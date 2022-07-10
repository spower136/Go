package calculator

import ("fmt"
	"unicode/utf8"
	"unicode"
	// "strings"
)

func Calc(shape string, op string, length float64, width float64) float64 {
	switch shape {
	case "rectangle":
		switch op {
		case "area":
			return length * width
		case "perimeter":
			return 2 * (length + width)
		}
	case "square":
		switch op {
		case "area":
			return length * length
		case "perimeter":
			return 4 * length
		}
	}
	return 0

}

func Sums(maxNum int) (threeMultSum, fiveMultSum int) {
	for i := 1; i <= maxNum; i++ {
		if i%3 == 0 {
			threeMultSum += i
		}
		if i%5 == 0 {
			fiveMultSum += i
		}
	}
	return threeMultSum, fiveMultSum
}

func Transform(s string) {
	fmt.Println(s, " has length ", utf8.RuneCountInString(s))
	t := ""
	for pos, char := range s {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
		switch {
		case unicode.IsLetter(char) :
			if unicode.IsUpper(char) {
				t += string(unicode.ToLower(char))
			} else {
				t += string(unicode.ToUpper(char))
			}
		case unicode.IsDigit(char) :
			continue
		case unicode.IsSpace(char) :
			t += string(char)
		case unicode.IsPunct(char) :
			continue
		case unicode.IsSymbol(char) :
			continue
		}
	}
	fmt.Println("Transformed", s, "to", t)
}