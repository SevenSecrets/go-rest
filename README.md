# Go REST

Teaching myself how to use Golang to make a REST API using gin.

Currently spits out random quotes from a couple of authors. Uses a locally hosted postgresql database of quotes for the moment, to help me teach myself about how to use postgresql with Go (and how to use databases with Go more generally).

### Running the server

Before running, create a .env file with the ```GO_DATABASE_PASS=YOUR_DB_PASSWORD``` inside, this is the password to your local PostgreSQL server. The PostgreSQL server is currently set to be on port 5432, called "go-rest" and using the default user "postgres". I'd make these all options you can change but this is for personal use, so the password being an environment variable one can set is only there for security.

Runs with ```run go ./server.go``` and serves on "/", "/perlman", and "/camatte" at the moment, although "/" is just a page telling you to go to one of the other two.


### Current plans: 
  - split the queries into a separate method, to keep it a bit cleaner.
  - add a few more authors and quotes.
  - why do I have a static html page being served on index? well, it was originally for me learning to use gin but now I'll have to find something else to put on "/".
  - see about making it possible to see all of the quotes together by author, and then select individual ones.
  - this is also probably going to be what I use to teach myself how to use AWS properly, and NGINX.