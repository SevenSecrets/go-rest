# Go REST

Teaching myself how to use Golang to make a REST API using gin.

Currently spits out random quotes from a couple of authors. Uses a locally hosted postgresql database of quotes for the moment, to help me teach myself about how to use postgresql with Go (and how to use databases with Go more generally).

Runs with ```run go ./server.go``` and serves on "/", "/perlman", and "/camatte" at the moment.


Current plans: 
  - split the queries into a separate method, to keep it a bit cleaner.
  - add a few more authors and quotes.
  - get to learning how to use environment variables in Go so I can use the password for the database properly.
  - why do I have a static html page being served on index? well, it was originally for me learning to use gin but now I'll have to find something else to put on "/".
