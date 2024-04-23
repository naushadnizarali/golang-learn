package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Users struct {
	id           int64
	userName     string
	firstName    string
	lastName     string
	emailAddress string
}

func main() {
	godotenv.Load(".env")

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_URI"),  //"127.0.0.1:3307",
		DBName: os.Getenv("DB_NAME"), //"hiresafe_db_dev",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	userByName, err := getUserByUsername("superadmin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("userByName found: %v\n", userByName)

	userByID, err := getUserByID(3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("userByID found: %v\n", userByID)
}

func getUserByUsername(userName string) ([]Users, error) {
	var users []Users

	rows, err := db.Query("SELECT id, userName, firstName, lastName, emailAddress from Users WHERE userName = ?", userName)

	if err != nil {
		return nil, fmt.Errorf("user not found %q: %v", userName, err)
	}

	defer rows.Close()

	for rows.Next() {
		var user Users

		if err := rows.Scan(&user.id, &user.userName, &user.firstName, &user.lastName, &user.emailAddress); err != nil {
			return nil, fmt.Errorf("not found %q: %v", userName, err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("user not found %q: %v", userName, err)
	}

	return users, nil
}

func getUserByID(id int) ([]Users, error) {
	var users []Users

	rows, err := db.Query("SELECT id, userName, firstName, lastName, emailAddress from Users WHERE id = ?", id)

	if err != nil {
		return nil, fmt.Errorf("user not found %q: %v", id, err)
	}

	defer rows.Close()

	for rows.Next() {
		var user Users

		if err := rows.Scan(&user.id, &user.userName, &user.firstName, &user.lastName, &user.emailAddress); err != nil {
			return nil, fmt.Errorf("not found %q: %v", id, err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("user not found %q: %v", id, err)
	}

	return users, nil
}
