package main

import (
	// "context"
	"fmt"
	// "os"
	"math/rand"
	// "io/ioutil"
	// "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	// "github.com/go-pg/pg/v10/orm"
)

type Quote struct {
	author string
	quote string
}

func (q Quote) String() string {
	return fmt.Sprintf("Quote<%s %s>", q.author, q.quote)
}

func quoteGenerator(quotes []string) string {
	quote := quotes[rand.Intn(len(quotes))]
	return quote
}

func main() {

	db := pg.Connect(&pg.Options{
		Addr: ":5432",
		User: "postgres",
		Password: "",
		Database: "go-rest",
	})
	defer db.Close()

	r := gin.Default()

	r.StaticFile("/", "./static/index.html")

	r.GET("/perlman", func(c *gin.Context) {

		var quotes []string
		
		err := db.Model((*Quote)(nil)).
		ColumnExpr("array_agg(quote)").
		Where("author = ?", "Perlman").
		Select(pg.Array(&quotes))
		if err != nil {
			panic(err)
		}

		quote := quoteGenerator(quotes)

		var response struct {
			Quote string
		}
		response.Quote = quote
		c.JSON(http.StatusOK, response)
	})

	r.GET("/camatte", func(c *gin.Context) {
		var quotes []string
		
		err := db.Model((*Quote)(nil)).
		ColumnExpr("array_agg(quote)").
		Where("author = ?", "Camatte").
		Select(pg.Array(&quotes))
		if err != nil {
			panic(err)
		}

		quote := quoteGenerator(quotes)

		var response struct {
			Quote string
		}
		response.Quote = quote
		c.JSON(http.StatusOK, response)
	})

	r.POST("/", func(c *gin.Context) {
		fmt.Println("POST RECEIVED")
	})

	r.Run(":8080")
}