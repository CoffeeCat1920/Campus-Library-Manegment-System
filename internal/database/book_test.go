package database

// import (
// 	"CLH/internal/modals"
// 	"testing"
//
// 	"github.com/DATA-DOG/go-sqlmock"
// )
//
// func TestAddBook(t *testing.T) {
// 	db, mock, err := sqlmock.New()
//
// 	if err != nil {
// 		t.Fatal("Can't Initalize database")
// 	}
//
// 	defer db.Close()
//
// 	s := &service{db: db}
//
// 	testCases := []struct {
// 		name        string
// 		book        *modals.Book
// 		setupMock   func(recipe *modals.Book)
// 		expectedErr error
// 	}{
// 		{
// 			name: "Success - Book Added Successfully",
// 			book: &modals.Book{
// 				ISBN:      "test-book-123",
// 				Name:      "test-book",
// 				Available: true,
// 			},
// 			setupMock: func(book *modals.Book) {
// 				mock.ExpectExec(`INSERT INTO books \( isbn, name, available \) VALUES\(\$1, \$2, \$3, -1\)`).
// 					WithArgs(book.ISBN, book.Name, book.Available).
// 					WillReturnResult(sqlmock.NewResult(1, 1))
// 			},
// 			expectedErr: nil,
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		tc.setupMock(tc.book)
//
// 		err := s.AddBook(tc.book)
//
// 		if err != tc.expectedErr {
// 			if err != nil && tc.expectedErr == nil {
// 				t.Errorf("In %s, Expected no error but got %v", tc.name, err)
// 			} else if err != nil && tc.expectedErr != nil {
// 				t.Errorf("In %s, Expected error %v but got %v", tc.name, tc.expectedErr, err)
// 			} else if err == nil && tc.expectedErr != nil {
// 				t.Errorf("In %s, Expected error %v but got nil", tc.name, tc.expectedErr)
// 			}
// 		}
//
// 	}
//
// }
//
// func TestUpdateBooks(t *testing.T) {
// 	db, mock, err := sqlmock.New()
//
// 	if err != nil {
// 		t.Fatal("Can't Initalize database")
// 	}
//
// 	defer db.Close()
//
// 	s := &service{db: db}
//
// 	testCases := []struct {
// 		name        string
// 		ISBN        string
// 		UpdatedName string
// 		setupMock   func(ISBN, UpdatedName string)
// 		expectedErr error
// 	}{
// 		{
// 			name:        "Success - Updated Book Successfully",
// 			ISBN:        "test-book-123",
// 			UpdatedName: "New_Name",
// 			setupMock: func(ISBN, UpdatedName string) {
// 				mock.ExpectExec("UPDATE books name = \\$1 WHERE isbn = \\$2").
// 					WithArgs(UpdatedName, ISBN).
// 					WillReturnResult(sqlmock.NewResult(0, 1))
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:        "Failure - No Book Found",
// 			ISBN:        "test-book-123",
// 			UpdatedName: "New_Name",
// 			setupMock: func(ISBN, UpdatedName string) {
// 				mock.ExpectExec("UPDATE books name = \\$1 WHERE isbn = \\$2").
// 					WithArgs(UpdatedName, ISBN).
// 					WillReturnResult(sqlmock.NewResult(0, 0))
// 			},
// 			expectedErr: ErrItemNotFound,
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		tc.setupMock(tc.ISBN, tc.UpdatedName)
//
// 		err := s.DeleteBook(tc.ISBN)
//
// 		if err != tc.expectedErr {
// 			if err != nil && tc.expectedErr == nil {
// 				t.Errorf("In %s, Expected no error but got %v", tc.name, err)
// 			} else if err != nil && tc.expectedErr != nil {
// 				t.Errorf("In %s, Expected error %v but got %v", tc.name, tc.expectedErr, err)
// 			} else if err == nil && tc.expectedErr != nil {
// 				t.Errorf("In %s, Expected error %v but got nil", tc.name, tc.expectedErr)
// 			}
// 		}
//
// 	}
// }
//
// func TestDeleteBooks(t *testing.T) {
// 	db, mock, err := sqlmock.New()
//
// 	if err != nil {
// 		t.Fatal("Can't Initalize database")
// 	}
//
// 	defer db.Close()
//
// 	s := &service{db: db}
//
// 	testCases := []struct {
// 		name        string
// 		ISBN        string
// 		setupMock   func(uuid string)
// 		expectedErr error
// 	}{
// 		{
// 			name: "Success - Book Deleted Successfully",
// 			ISBN: "test-book-123",
// 			setupMock: func(ISBN string) {
// 				mock.ExpectExec("DELETE FROM books WHERE isbn = \\$1").
// 					WithArgs(ISBN).
// 					WillReturnResult(sqlmock.NewResult(0, 1))
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name: "Failure - Recipe Not Found",
// 			ISBN: "test-book-123",
// 			setupMock: func(ISBN string) {
// 				mock.ExpectExec("DELETE FROM recipes WHERE isbn = \\$1").
// 					WithArgs(ISBN).
// 					WillReturnResult(sqlmock.NewResult(0, 0))
// 			},
// 			expectedErr: ErrItemNotFound,
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		tc.setupMock(tc.ISBN)
//
// 		err := s.DeleteBook(tc.ISBN)
//
// 		if err != tc.expectedErr {
// 			if err != nil && tc.expectedErr == nil {
// 				t.Errorf("In %s, Expected no error but got %v", tc.name, err)
// 			} else if err != nil && tc.expectedErr != nil {
// 				t.Errorf("In %s, Expected error %v but got %v", tc.name, tc.expectedErr, err)
// 			} else if err == nil && tc.expectedErr != nil {
// 				t.Errorf("In %s, Expected error %v but got nil", tc.name, tc.expectedErr)
// 			}
// 		}
//
// 	}
//
// }
//
// func TestSearchBook(t *testing.T) {
// 	db, mock, err := sqlmock.New()
//
// 	if err != nil {
// 		t.Errorf("Can't initialize sqlmock")
// 	}
//
// 	s := &service{db: db}
//
// 	testCases := []struct {
// 		name        string
// 		searchTerm  string
// 		setupMock   func(searchTerm string)
// 		expectedErr error
// 	}{
// 		{
// 			name:       "Success - Got Searched Recipe",
// 			searchTerm: "test",
// 			setupMock: func(searchTerm string) {
// 				rows := sqlmock.NewRows([]string{"isbn", "name", "available"}).
// 					AddRow(`test-uuid-123`, `test1`, `test-uuid-123`, 0).
// 					AddRow(`test-uuid-123`, `test2`, `test-uuid-123`, 1).
// 					AddRow(`test-uuid-123`, `test3`, `test-uuid-123`, 2)
//
// 				mock.ExpectQuery("SELECT \\* FROM books WHERE name ILIKE \\$1").
// 					WithArgs("%" + searchTerm + "%").
// 					WillReturnRows(rows)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:       "Success - No Recipe Found",
// 			searchTerm: "test",
// 			setupMock: func(searchTerm string) {
// 				rows := sqlmock.NewRows([]string{"isbn", "name", "available"})
//
// 				mock.ExpectQuery("SELECT \\* FROM books WHERE name ILIKE \\$1").
// 					WithArgs("%" + searchTerm + "%").
// 					WillReturnRows(rows)
// 			},
// 			expectedErr: nil,
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			tc.setupMock(tc.searchTerm)
//
// 			_, err := s.SearchBooks(tc.searchTerm)
//
// 			if err != tc.expectedErr {
// 				if err != nil && tc.expectedErr == nil {
// 					t.Errorf("In %s, Expected no error but got %v", tc.name, err)
// 				} else if err != nil && tc.expectedErr != nil {
// 					t.Errorf("In %s, Expected error %v but got %v", tc.name, tc.expectedErr, err)
// 				} else if err == nil && tc.expectedErr != nil {
// 					t.Errorf("In %s, Expected error %v but got nil", tc.name, tc.expectedErr)
// 				}
// 			}
//
// 		})
// 	}
//
// }
