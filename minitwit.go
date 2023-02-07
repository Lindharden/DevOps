package main

import "github.com/gin-gonic/gin"
import "net/http"
import "log"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "fmt"

//global database variable
var DB *sql.DB

// Configuration
const DATABASE = "/tmp/minitwit.db"

func setupRouter() *gin.Engine {
	router := gin.Default();
	router.LoadHTMLGlob("templates/*.tmpl");
	//router.LoadHTMLFiles("templates/index.tmpl");
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":"yo dudes",
		})
	})
	return router
}

func main() {
	connect_db()
	id := get_user_id("Leonora")
	fmt.Println("id: " + id)
	r := setupRouter()
	r.Run(":8080")
}


func connect_db() *sql.DB {
	db, err := sql.Open("sqlite3", DATABASE)
	if err != nil {
		log.Fatal(err)
	}
	
	DB = db
	return db
}

func get_user_id(username string) string {
	var id string
	err := DB.QueryRow("select user_id from user where username = ?", username).Scan(&id)
	if err != nil { 
		log.Fatal(err)
	}
	return id
}