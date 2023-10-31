package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type User struct {
	UserID   int
	UserName string
	Email    string
}

func main() {
	e := echo.New()

	// Database connection
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=root dbname=taskManager sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		return
	}
	defer db.Close()

	// Migrate user table
	migrate(db)

	e.POST("/users", createUser)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/users", getAllUsers)

	e.Logger.Fatal(e.Start(":8800"))
}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}
func getAllUsers(c echo.Context) error {
	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		return err
	}
	defer rows.Close()

	var usrs []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user, user.UserID, &user.UserName, &user.Email)
		if err != nil {
			return err
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOk, users)
}

func migrate(db *sql.DB) {
	// Implement the logic to create a User table if it doesn't exist in the PostgreSQL database
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users(
		user_id SERIAL PRIMARY KEY,
		user_name TEXT,
		email TEXT
	)`)
	if err != nil {
		fmt.Println("Failed to create table:", err)
	}
}
