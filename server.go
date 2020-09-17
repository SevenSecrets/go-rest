package main

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"os"
	"math/rand"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

func quoteGenerator(author string) string {
	jsonFile, err := os.Open("quotes.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("successfully opened JSON file.")

	defer jsonFile.Close()

	byt, _ := ioutil.ReadAll(jsonFile)

	var res map[string]interface{}
	json.Unmarshal([]byte(byt), &res)

	quotes := res[author].([]interface{})
	quote := quotes[rand.Intn(len(quotes))].(string)

	return quote
}

func main() {
	r := gin.Default()

	r.StaticFile("/", "./static/index.html")

	r.GET("/perlman", func(c *gin.Context) {
		quote := quoteGenerator("perlman")
		var response struct {
			Quote string
		}
		response.Quote = quote
		c.JSON(http.StatusOK, response)
	})

	r.GET("/camatte", func(c *gin.Context) {
		quote := quoteGenerator("camatte")
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