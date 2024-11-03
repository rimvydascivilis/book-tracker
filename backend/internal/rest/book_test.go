package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/rimvydascivilis/book-tracker/backend/domain"
	"github.com/rimvydascivilis/book-tracker/backend/internal/rest"
	"github.com/rimvydascivilis/book-tracker/backend/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockJWTToken = jwt.Token{
	Claims: jwt.MapClaims{
		"id": float64(1),
	},
}

func TestGetBooks(t *testing.T) {
	mockSvc := new(mocks.BookService)
	handler := rest.NewBookHandler(mockSvc)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books?page=1&limit=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &mockJWTToken)

	var rating float64 = 5
	var rating2 float64 = 4
	books := []domain.Book{
		{ID: 1, Title: "Book One", Rating: &rating},
		{ID: 2, Title: "Book Two", Rating: &rating2},
	}
	mockSvc.On("GetBooks", mock.Anything, int64(1), int64(1), int64(2)).Return(books, true, nil)

	err := handler.GetBooks(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "books")
	assert.Contains(t, response, "hasMore")
}

func TestCreateBook(t *testing.T) {
	mockSvc := new(mocks.BookService)
	handler := rest.NewBookHandler(mockSvc)

	e := echo.New()
	var rating float64 = 5
	book := domain.Book{Title: "New Book", Rating: &rating}
	body, _ := json.Marshal(book)

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &mockJWTToken)

	createdBook := book
	createdBook.ID = 1
	mockSvc.On("CreateBook", mock.Anything, int64(1), book).Return(createdBook, nil)

	err := handler.CreateBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var responseBook domain.Book
	err = json.Unmarshal(rec.Body.Bytes(), &responseBook)
	assert.NoError(t, err)
	assert.Equal(t, createdBook, responseBook)
}

func TestUpdateBook(t *testing.T) {
	mockSvc := new(mocks.BookService)
	handler := rest.NewBookHandler(mockSvc)

	e := echo.New()
	var rating float64 = 5
	book := domain.Book{ID: 1, Title: "Updated Title", Rating: &rating}
	body, _ := json.Marshal(book)

	req := httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &mockJWTToken)
	c.SetParamNames("id")
	c.SetParamValues(strconv.FormatInt(book.ID, 10))

	mockSvc.On("UpdateBook", mock.Anything, int64(1), book).Return(book, nil)

	err := handler.UpdateBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBook domain.Book
	err = json.Unmarshal(rec.Body.Bytes(), &responseBook)
	assert.NoError(t, err)
	assert.Equal(t, book, responseBook)
}

func TestDeleteBook(t *testing.T) {
	mockSvc := new(mocks.BookService)
	handler := rest.NewBookHandler(mockSvc)

	e := echo.New()
	bookID := int64(1)

	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &mockJWTToken)
	c.SetParamNames("id")
	c.SetParamValues(strconv.FormatInt(bookID, 10))

	mockSvc.On("DeleteBook", mock.Anything, int64(1), bookID).Return(nil)

	err := handler.DeleteBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, "", rec.Body.String())
}

func TestCreateBook_InvalidInput(t *testing.T) {
	mockSvc := new(mocks.BookService)
	handler := rest.NewBookHandler(mockSvc)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader([]byte(`{"title": ""}`))) // Invalid title
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", &mockJWTToken)

	mockSvc.On("CreateBook", mock.Anything, int64(1), mock.Anything).Return(domain.Book{}, domain.ErrInvalidBook)

	err := handler.CreateBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request format")
}
