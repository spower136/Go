package sort

type byAuthorLName interface {
	len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// type byPrice []book

// func (a byPrice) Len() int           { return len(a) }
// func (a byPrice) Less(i, j int) bool { return a[i] < a[j] }
// func (a byPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
