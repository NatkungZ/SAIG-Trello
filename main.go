package main

import (
	"fmt"
	"io/ioutil"
//	"strconv"
	"encoding/json"
	"log"
	"net/http"
//	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	if body, err := ioutil.ReadAll(r.Body); err==nil {
		log.Println("Data",string(body))
		fmt.Fprint(w, "Data",string(body))
		// here
		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err != nil {
			log.Println(err)
			fmt.Fprint(w, "Data",string(body))
			return
		}
		xEvent := dat["action"].(map[string]interface{})["type"]

		if (dat["action"].(map[string]interface{})["data"].(map[string]interface{})["listAfter"] != nil){
			xEvent = "listAfter"
		}
		if (dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["closed"] == true){
                        xEvent = "deleteCard"
		}

		fmt.Println(xEvent)
		
		var msg string
		switch xEvent {
			case "updateCard":
				msg = fmt.Sprintf("UpdateCard in %s by @%s",dat["model"].(map[string]interface{})["name"],
					dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
			case "listAfter":
			msg = fmt.Sprintf("Move card \"%s\" in %s from list \"%s\" to \"%s\" by @%s",dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["name"],dat["model"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["listBefore"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["listAfter"].(map[string]interface{})["name"],
		dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
			case "createCard":
	msg = fmt.Sprintf("Create card \"%s\" on %s Board in list \"%s\" by @%s",dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["name"],dat["model"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["list"].(map[string]interface{})["name"],
		dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
			case "deleteCard":
	msg = fmt.Sprintf("Delete card \"%s\" on %s Board in list \"%s\" by @%s",dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["name"],dat["model"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["list"].(map[string]interface{})["name"],
		dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
			case "commentCard":
	msg = fmt.Sprintf("Comment card \"%s\" text \"%s\" on %s Board in list \"%s\" by @%s",dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["text"],dat["model"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["list"].(map[string]interface{})["name"],
		dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
			case "addChecklistToCard":
	msg = fmt.Sprintf("Create checklist \"%s\" in card \"%s\" on %s Board by @%s",dat["action"].(map[string]interface{})["data"].(map[string]interface{})["checklist"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["name"],dat["model"].(map[string]interface{})["name"],
		dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
			case "createCheckItem":
	msg = fmt.Sprintf("Create checklist item \"%s\" in checklist \"%s\" in card \"%s\" on %s Board by @%s",dat["action"].(map[string]interface{})["data"].(map[string]interface{})["checkItem"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["checklist"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["name"],dat["model"].(map[string]interface{})["name"],
		dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
			case "updateCheckItemStateOnCard":
	msg = fmt.Sprintf("Check item \"%s\" in checklist \"%s\" in card \"%s\" on %s Board by @%s",dat["action"].(map[string]interface{})["data"].(map[string]interface{})["checkItem"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["checklist"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["name"],dat["model"].(map[string]interface{})["name"],
		dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
			case "removeChecklistFromCard":
	msg = fmt.Sprintf("Removed checklist \"%s\" in card \"%s\" on %s Board by @%s",dat["action"].(map[string]interface{})["data"].(map[string]interface{})["checklist"].(map[string]interface{})["name"],dat["action"].(map[string]interface{})["data"].(map[string]interface{})["card"].(map[string]interface{})["name"],dat["model"].(map[string]interface{})["name"],
		dat["action"].(map[string]interface{})["memberCreator"].(map[string]interface{})["username"])
	
		default:
			msg = ""
		}
		fmt.Println(msg)
		
	if msg != "" {
		_, err := fbBot("saig2015", msg)
		if err != nil {
			log.Println(err)
		}
	} else {
	}

	}else{
		log.Println("Home",string(r.URL.Path))
		fmt.Fprintf(w, "Home %s",string( r.URL.Path))
	}
}

func fbBot(group string, msg string) ([]byte, error) {
	fbHost := "http://172.17.42.1:8000" // TODO:Change this
	reqUrl := fbHost + "/" + group + "/" + msg
	log.Println(reqUrl)
	res, err := http.Get(reqUrl)
	if err != nil {
		log.Println(err)
	}

	bot, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return bot, err
}


func main() {
	// Go Web
	port := ":9000"

	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}
