/*

  Flat DB - A truly simplistic, in-memory and likely terrible NoSQL database. Bear with me.

  Author - Josh McGhee
  Email - hello@joshuamcghee.com

*/

package main

import "net/http"
import "fmt"
import "html"
import "os"

type db struct {
    storedValues  map[string]string
}

/*
  Database commands
*/

func getCommand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I got nothin, %q", html.EscapeString(r.URL.Path))
}

/*
  Initialise the database
*/

func initDb() db {
	var data db = db{storedValues: make(map[string]string)}
	return data
}

/*
  Configure and run the HTTP server
*/

func configureHandlers(data *db) {
	http.HandleFunc("/GET", getCommand)
}

func runServer(port string) {
	http.ListenAndServe(":" + port,nil)
}

/*
  Main program
*/


func main() {
	var port string
	var data db

	port = os.Args[1]

	// Build a KV map, knock it in a struct and return
	data = initDb()

	// Setup the handlers to deal with the struct
	configureHandlers(&data)

	// Run the HTTP server. Eeeek!
        runServer(port)

}
