package database

import (
	"CLH/internal/modals"
)

func (s *service) AddBook(book *modals.Book) error {
	q := `
  INSERT INTO books(isbn, name, author, genre, available)
  VALUES($1, $2, $3, $4, $5)
  `

	_, err := s.db.Exec(q, book.ISBN, book.Name, book.Available, book.Genre)

	if PrimaryKeyError(err) {
		return ErrItemPrimaryKeyExits
	}

	if IsUniqueViolation(err) {
		return ErrItemAlreadyExists
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateBookName(ISBN string, name string) error {
	q := `
	UPDATE books
	SET name = $1
	WHERE isbn = $2
  `
	_, err := s.db.Exec(q, name, ISBN)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateBookGenre(ISBN string, genre string) error {
	q := `
	UPDATE books
	SET genre = $1
	WHERE isbn = $2
  `
	_, err := s.db.Exec(q, genre, ISBN)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteBook(ISBN string) error {

	q := "DELETE FROM books WHERE isbn = $1"

	res, err := s.db.Exec(q, ISBN)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrItemNotFound
	}

	return nil
}

func (s *service) SearchBooks(name string) ([]modals.Book, error) {
	var recipes []modals.Book

	searchTerm := "%" + name + "%"
	query := "SELECT * FROM books WHERE name ILIKE $1"

	rows, err := s.db.Query(query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe modals.Book
		err := rows.Scan(&recipe.ISBN, &recipe.Name, &recipe.Available, &recipe.Genre, &recipe.Author)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (s *service) GetAllBooks() ([]modals.Book, error) {
	var recipes []modals.Book

	query := "SELECT * FROM books"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe modals.Book
		err := rows.Scan(&recipe.ISBN, &recipe.Name, &recipe.Available, &recipe.Genre, &recipe.Author)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
