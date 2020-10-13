# Go REST

Teaching myself how to use Golang to make a REST API using gin.

Currently spits out random quotes from a couple of authors. Uses a locally hosted postgresql database of quotes for the moment, to help me teach myself about how to use postgresql with Go (and how to use databases with Go more generally).

Runs with ```run go ./server.go``` and serves on "/", "/perlman", and "/camatte" at the moment.