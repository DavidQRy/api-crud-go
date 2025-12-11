package main

import (
	"api-crud-go/internal/service"
	"api-crud-go/internal/store"
	"api-crud-go/internal/transport"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"

	"log"
	"net/http"
)

func main() {
	// connect to SQLite
	db, err := sql.Open("sqlite", "./book.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//Create table if not exist
	q := `CREATE TABLE IF NOT EXISTs books (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	author TEXT NOT NULL
	);
	`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inject the dependencies
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	// Routes Config
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/book/", bookHandler.HandleBookByID)

	// logs
	fmt.Println("🚀 Server running at http://localhost:8000")
	fmt.Println("📚 API Endpoints:")
	fmt.Println("  GET    /books        - Retrieve all books")
	fmt.Println("  POST   /books        - Create a new book")
	fmt.Println("  GET    /book/{id}   - Retrieve a specific book")
	fmt.Println("  PUT    /book/{id}   - Update a book")
	fmt.Println("  DELETE /book/{id}   - Delete a book")

	//Server up
	log.Fatal(http.ListenAndServe(":8000", nil))
}
