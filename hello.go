package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"database/sql"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"log"
	"html/template"
)

func main() {
	// two entry points: "localhost:8080" which is where you enter a url to be shortened
	// and "localhost:8080/url/<shortened url>" which will redirect to the long url via the short url
    http.HandleFunc("/", pageHandler)
    http.HandleFunc("/url/", urlHandler)
    err := http.ListenAndServe(":8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// loading "localhost:8080" will read the html file "result.html"
    if r.Method == "GET" {
        t, _ := template.ParseFiles("result.html")
        t.Execute(w, nil)
    } else {
    	// when "shorten" buttton pressed, the url the user entered is
    	// sent to the database, the returned id is converted to a short string
    	// and that string is printed for the user
        r.ParseForm()
        id := addURL(r.Form["URL"][0])
        encoded_string := numToShortString(id)
        fmt.Fprintf(w, "Here's your shortened URL:\n/url/%s", encoded_string)
    }
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
    // the short url is converted to the id of the id in the database
    // where the long url is stored, and user is redirected to the url they
    // requested
    id := strToNum(r.URL.Path[5:])
    url := queryDB(id)
    http.Redirect(w,r, url, http.StatusSeeOther)
}

func addURL(url string)(int64){
	// add URL to database, return id

	db, err := sql.Open("sqlite3", "./urls.db")
    checkErr(err)
    
    stmt, err := db.Prepare("INSERT INTO urltable(url) values(?)")
    checkErr(err)

    res, err := stmt.Exec(url)
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    db.Close()

    fmt.Println("Added URL to database")

	return(id)
}

func queryDB(id uint64)(url_final string){
	// gets a URL from the database from a given id
	db, err := sql.Open("sqlite3", "./urls.db")
    checkErr(err)

    var query bytes.Buffer

    query.WriteString("SELECT url FROM urltable WHERE id==")
    query.WriteString(strconv.FormatUint(id, 10))

    instance, err := db.Query(query.String())
    checkErr(err)
    
    var url string

    for instance.Next() {
        err = instance.Scan(&url)
        checkErr(err)
        fmt.Println(url)
    }

    db.Close()

    return(url)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func numToShortString(strID int64)(string){
	// converts a number (the id from the database entry) to a shorter
	// string that acts as a url
	var strForConversion string= "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ$-_.+!*'(),"
	var l int64 = int64(len(strForConversion))

	var remainder int64= strID % l
	var result string= string(strForConversion[remainder])
	var div int64 = int64(math.Floor(float64(strID) / float64(l)))
	for div > 0{
		remainder = div % l
		div = int64(math.Floor(float64(div) / float64(l)))
		result = string(strForConversion[remainder]) + result
	}
	return result
}

func strToNum(myURL string)(uint64){
	// converts a short url into a number that corresponds to the 
	// database id of the longer url
	var strForConversion string= "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ$-_.+!*'(),"
	var l uint64 = uint64(len(strForConversion))

	var length uint64 = uint64(len(myURL))
	var result uint64 = 0
	for i := 0; uint64(i) < length; i++ {
		result = l * result + uint64(strings.Index(strForConversion, string(myURL[i])))
	}
	return result
}

