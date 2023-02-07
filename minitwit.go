package main

import "github.com/gin-gonic/gin"
import "net/http"
import "log"
import "database/sql"


// Configuration
const DATABASE = "itu-minitwit.db"

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
	r := setupRouter()
	r.Run(":8080")
}


func connect_db() *sql.DB {
	//os.Remove(DATABASE)

	db, err := sql.Open("sqlite3", DATABASE)

	if (err != nil) {
		log.Fatal(err)
	}
	return db;
}

func get_user_id(username) {
	
}