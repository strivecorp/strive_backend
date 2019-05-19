package API

import (
	"encoding/json"
	"fmt"
	"go/types"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strive_backend/API/structs"
)

func APIHandler() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		fmt.Fprintln(w, "use these!: /Users /User/id")
		fmt.Fprintln(w, "/Users /User/id")
		fmt.Fprintln(w, "/Goals /Goal/id")
		fmt.Fprintln(w, "/Feeds /Feed/id")
		fmt.Fprintln(w, "/Milestones /Milestone/id")
		fmt.Fprintln(w, "/Comments /Comment/id")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		//fmt.Fprintln(w, "Users data below")
		var idString = r.URL.Query().Get("id")
		if(idString!=""){
			idInt, err := strconv.Atoi(idString)
			if err != nil {
				fmt.Println("Invalid query")
			}
			fmt.Fprintln(w, readJsonById("users.json", idInt))
		}else{
			fmt.Fprintln(w,getUsers())
		}
	})
	http.HandleFunc("/goals", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		//fmt.Fprintln(w, "Goals data below")
		var idString = r.URL.Query().Get("id")
		if(idString!=""){
			idInt, err := strconv.Atoi(idString)
			if err != nil {
				fmt.Println("Invalid query")
			}
			fmt.Fprintln(w, readJsonById("goals.json", idInt))
		}else{
			fmt.Fprintln(w,getGoals())
		}
	})
	http.HandleFunc("/feeds", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		//fmt.Fprintln(w, "Feed data below")
		var idString = r.URL.Query().Get("id")
		if(idString!=""){
			idInt, err := strconv.Atoi(idString)
			if err != nil {
				fmt.Println("Invalid query")
			}
			fmt.Fprintln(w, readJsonById("feed_item.json", idInt))
		}else{
			fmt.Fprintln(w,getFeedItems())
		}
	})
	http.HandleFunc("/milestones", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		//fmt.Fprintln(w, "Milestone data below")
		var idString = r.URL.Query().Get("id")
		if(idString!=""){
			idInt, err := strconv.Atoi(idString)
			if err != nil {
				fmt.Println("Invalid query")
			}
			fmt.Fprintln(w, readJsonById("milestones.json", idInt))
		}else{
			fmt.Fprintln(w,getMilestones())
		}
	})
	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		//fmt.Fprintln(w, "Comments data below")
		var idString = r.URL.Query().Get("id")
		if(idString!=""){
			idInt, err := strconv.Atoi(idString)
			if err != nil {
				fmt.Println("Invalid query")
			}

			fmt.Fprintln(w, readJsonById("comments.json", idInt))
		}else{
			fmt.Fprintln(w,getComments())
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers() string {
	return readJson("users.json")
}

func getGoals() string {
	return readJson("goals.json")
}

func getFeedItems() string {
	return readJson("feed_item.json")
}

func getMilestones() string {
	return readJson("milestones.json")
}

func getComments() string {
	return readJson("comments.json")
}

func readJson(fileName string) string {
	b, err := ioutil.ReadFile("API/dummy_data/" + fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

/*
	var t []structs.Goal
	errRead := json.Unmarshal(b, &t)
	if errRead != nil {
		panic(errRead)
	}
	commentsJson, err := json.Marshal(t)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
*/

func readJsonById(fileName string, id int) string {
	bigBoyJson, err := ioutil.ReadFile("API/dummy_data/" + fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	switch fileName {
	case "comments.json":
		var t []structs.Comment
		errRead := json.Unmarshal(bigBoyJson, &t)
		if errRead != nil {
			panic(errRead)
		}
		for _, element := range t {
			if element.Id == id {
				json, err := json.MarshalIndent(element, "", "    ")
				if err != nil {
					log.Fatal("Cannot encode to JSON ", err)
				}
				return string(json)
			}
		}
		return "{ error: '404', description: 'Resource not found' }"
	case "feed_item.json":
		var t []structs.FeedItem
		errRead := json.Unmarshal(bigBoyJson, &t)
		if errRead != nil {
			panic(errRead)
		}
		for _, element := range t {
			if element.Id == id {
				json, err := json.MarshalIndent(element, "", "    ")
				if err != nil {
					log.Fatal("Cannot encode to JSON ", err)
				}
				return string(json)
			}
		}
		return "{ error: '404', description: 'Resource not found' }"
	case "goals.json":
		var t []structs.Goal
		errRead := json.Unmarshal(bigBoyJson, &t)
		if errRead != nil {
			panic(errRead)
		}
		for _, element := range t {
			if element.Id == id {
				json, err := json.MarshalIndent(element, "", "    ")
				if err != nil {
					log.Fatal("Cannot encode to JSON ", err)
				}
				return string(json)
			}
		}
		return "{ error: '404', description: 'Resource not found' }"
	case "milestones.json":
		var t []structs.Milestone
		errRead := json.Unmarshal(bigBoyJson, &t)
		if errRead != nil {
			panic(errRead)
		}
		for _, element := range t {
			if element.Id == id {
				json, err := json.MarshalIndent(element, "", "    ")
				if err != nil {
					log.Fatal("Cannot encode to JSON ", err)
				}
				return string(json)
			}
		}
		return "{ error: '404', description: 'Resource not found' }"
	case "users.json":
		var t []structs.User
		errRead := json.Unmarshal(bigBoyJson, &t)
		if errRead != nil {
			panic(errRead)
		}
		for _, element := range t {
			if element.Id == id {
				json, err := json.MarshalIndent(element, "", "    ")
				if err != nil {
					log.Fatal("Cannot encode to JSON ", err)
				}
				return string(json)
			}
		}
		return "{ error: '404', description: 'Resource not found' }"
	default:
		panic(types.Interface{})
	}

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}