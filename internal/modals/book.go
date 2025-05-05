package modals

type Book struct {
	ISBN      string `json:"isbn"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Available bool   `json:"available"`
}

func NewBook(isbn, name string, genre string) *Book {
	return &Book{
		ISBN:      isbn,
		Name:      name,
		Genre:     genre,
		Available: true,
	}
}
