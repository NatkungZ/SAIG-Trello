package main

import (
	"fmt"
	"io/ioutil"
	//	"strconv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)


func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	if body, err := ioutil.ReadAll(r.Body); err==nil {
		log.Println("Data",string(body))
		fmt.Fprint(w, "Data",string(body))
	}else{
		log.Println("Home",string(r.URL.Path))
		fmt.Fprintf(w, "Home %s",string( r.URL.Path))
	}
}

func fbBot(group string, msg string) ([]byte, error) {
	fbHost := os.Getenv("GO_HOST")
	reqUrl := fbHost + "/" + group + "/" + msg
	log.Println(reqUrl)
	res, err := http.Get(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return bot, err
}


func main() {
	// Go Web
	port := 9000

	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}
