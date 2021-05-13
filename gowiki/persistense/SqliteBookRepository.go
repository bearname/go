package persistense

import (
	"awesomeProject/gowiki/model"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SqliteBookRepository struct {
	db *sql.DB
}

func NewSqliteBookRepository(filePath string) *SqliteBookRepository {
	//err := os.Remove("sqlite-database.db")
	//if err != nil {
	//	return nil
	//}
	//
	//log.Println("Creating sqlite-database.db...")
	//file, err := os.Create(filePath)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//err = file.Close()
	//if err != nil {
	//	return nil
	//}
	//log.Println("sqlite-database.db created")
	p := new(SqliteBookRepository)

	openDatabase, err := sql.Open("sqlite3", filePath)
	if err != nil {
		fmt.Println(err.Error())
	}

	p.db = openDatabase
	p.CreateTable()
	p.InsertBook("Go tutorial", "Bob marline")
	p.InsertBook("Go tutorial", "Bob marline")
	p.InsertBook("Go tutorial", "Bob marline")
	p.InsertBook("Building Web Apps with Go", "Jeremy Saenz")
	return p
}

func (r *SqliteBookRepository) CreateTable() {
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS book (
		"idBook" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"author" TEXT
	  );`

	log.Println("Create student table...")
	statement, err := r.db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.
		Exec()
	if err != nil {
		return
	}
	log.Println("student table created")
}

func (r *SqliteBookRepository) InsertBook(title string, author string) int64 {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO book(title, author) VALUES (?, ?);`
	statement, err := r.db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	result, err := statement.Exec(title, author)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Println("LastInsertId:", id)
			return id
		}
	}

	return 0
}

func (r *SqliteBookRepository) GetBooks() []model.Book {
	row, err := r.db.Query("SELECT * FROM book;")
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {

		}
	}(row)
	var books []model.Book
	for row.Next() {
		var id int64
		var title string
		var author string
		err := row.Scan(&id, &title, &author)
		if err != nil {
			return nil
		}
		books = append(books, model.Book{Id: id, Title: title, Author: author})
		log.Println("Student: ", id, " ", title, " ", author)
	}
	return books
}

func (r *SqliteBookRepository) GetBook(id int64) model.Book {
	stmt, err := r.db.Prepare("SELECT * FROM book WHERE idBook = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var book model.Book
	for rows.Next() {
		var id int64
		var title string
		var author string
		err = rows.Scan(&id, &title, &author)
		if err != nil {
			log.Fatal(err)
		}

		book = model.Book{Id: id, Title: title, Author: author}
		fmt.Println(id, title, author)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return book
}
