package sort

type byAuthorLName []book

func (a byAuthorLName) Len() int           { return len(a) }
func (a byAuthorLName) Less(i, j int) bool { return a[i] < a[j] }
func (a byAuthorLName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type byPrice []Book

func (a byPrice) Len() int           { return len(a) }
func (a byPrice) Less(i, j int) bool { return a[i] < a[j] }
func (a byPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
