package main

import (
	"database/sql"
	"net/http"

	"github.com/2e8f8d9d/go/demoapp/webservice/controllers"
	"github.com/2e8f8d9d/go/demoapp/webservice/models"

	// needed by sql
	_ "github.com/go-sql-driver/mysql"
)

// HTTPRequest is used to make http requests
type HTTPRequest struct {
	Method string
}

func main() {
	// Init table
	err := createTable()
	if err != nil {
		panic(err.Error())
	} else {
		println("users table created with id, first_name, last_name with primary key as id")
	}

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	r := HTTPRequest{Method: "Get"}

	switch r.Method {
	case "GET":
		println("Get request")
	case "DELETE":
		println("Delete request")
	default:
		println("Unhandled method")
	}
}

func createTable() error {

	db, err := sql.Open("mysql", models.ConnectionString)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("CREATE Table users(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id));")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
