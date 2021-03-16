package models

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"os"

	//Needed by database/sql/
	_ "github.com/go-sql-driver/mysql"
)

var (
	//ConnectionString used in main to init database
	ConnectionString = os.Getenv("CONNECTIONSTRING")
)

// User class
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// GetUsers get all users from database
func GetUsers() (*[]User, error) {
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		return nil, errors.New("Unable to connect to database")
	}

	defer db.Close()

	allUsers, err := db.Query("select * FROM users")
	if err != nil {
		return nil, errors.New("Unable to fetch users")
	}

	var arrUsers []User
	for allUsers.Next() {
		var user User

		err = allUsers.Scan(&user.ID,
			&user.FirstName,
			&user.LastName,
		)
		if err != nil {
			return nil, errors.New("unable to parse user data")
		}

		arrUsers = append(arrUsers, user)
	}
	return &arrUsers, nil
}

// AddUser add a user to the database
func AddUser(u User) (User, error) {
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		return User{}, errors.New("Unable to connect to database")
	}

	if u.ID != 0 {
		return User{}, errors.New("New User must no include id")
	}

	u.ID = rand.Intn(100)
	queryString := fmt.Sprintf("INSERT users Values('%v', '%v', '%v')", u.ID, u.FirstName, u.LastName)

	insert, err := db.Query(queryString)

	if err != nil {
		return User{}, errors.New("could not add user")
	}
	insert.Close()
	return u, nil
}

//GetUserByID returns a user based on id
func GetUserByID(id int) (User, error) {
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		return User{}, errors.New("Unable to connect to database")
	}

	defer db.Close()

	allUsers, err := db.Query("select * FROM users")
	if err != nil {
		return User{}, errors.New("Unable to fetch users")
	}

	for allUsers.Next() {
		var user User

		err = allUsers.Scan(&user.ID,
			&user.FirstName,
			&user.LastName,
		)
		if err != nil {
			return User{}, errors.New("unable to parse user data")
		}

		if id == user.ID {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

//UpdateUser updates a user account
func UpdateUser(u User) (User, error) {
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		return User{}, errors.New("Unable to connect to database")
	}

	defer db.Close()

	allUsers, err := db.Query("select * FROM users")
	if err != nil {
		return User{}, errors.New("Unable to fetch users")
	}

	for allUsers.Next() {
		var candidate User

		err = allUsers.Scan(&candidate.ID,
			&candidate.FirstName,
			&candidate.LastName,
		)
		if err != nil {
			return User{}, errors.New("unable to parse user data")
		}

		if candidate.ID == u.ID {
			candidate = u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID removes a user by id
func RemoveUserByID(id int) error {
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		return errors.New("Unable to connect to database")
	}

	defer db.Close()

	allUsers, err := db.Query("select * FROM users")
	if err != nil {
		return errors.New("Unable to fetch users")
	}

	for allUsers.Next() {
		var candidate User

		err = allUsers.Scan(&candidate.ID,
			&candidate.FirstName,
			&candidate.LastName,
		)
		if err != nil {
			return errors.New("unable to retreive user ids")
		}

		if id == candidate.ID {
			deleteString := fmt.Sprintf("DELETE FROM users WHERE id = %v;", id)
			db.Query(deleteString)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
