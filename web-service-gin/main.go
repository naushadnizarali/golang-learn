package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Users struct {
	ID           int64  `json:"id"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
}

var users = []Users{
	{ID: 1, UserName: "superadmin", FirstName: "Super Admin", LastName: "Super Administrator", EmailAddress: "admin@hiresafe.com"},
	{ID: 2, UserName: "client-1", FirstName: "Client", LastName: "1", EmailAddress: "client@hiresafe.com"},
	{ID: 3, UserName: "candidate-1", FirstName: "Candidate", LastName: "1", EmailAddress: "candidate@hiresafe.com"},
	{ID: 4, UserName: "altafali", FirstName: "Altaf", LastName: "Ali", EmailAddress: "altaf@hiresafe.com"},
	{ID: 5, UserName: "saad", FirstName: "Saad", LastName: "1", EmailAddress: "saad@hiresafe.com"},
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	godotenv.Load(".env")

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)

	port := os.Getenv("PORT")

	router.Run("localhost:" + port)
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Handle the error, maybe return a bad request response
	}

	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
