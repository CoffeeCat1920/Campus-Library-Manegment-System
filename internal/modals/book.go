package modals

type Book struct {
	ISBN      string
	Name      string
	Available bool
}

func NewBook(isbn, name string) *Book {
	return &Book{
		ISBN:      isbn,
		Name:      name,
		Available: true,
	}
}
