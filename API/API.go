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

		fmt.Fprintln(w, "use these!: ")
		fmt.Fprintln(w, "/users /users?id=1")
		fmt.Fprintln(w, "/goals /goals?id=1")
		fmt.Fprintln(w, "/feeds /feeds?id=1")
		fmt.Fprintln(w, "/milestones /milestones?id=1")
		fmt.Fprintln(w, "/comments /comments?id=1")
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

	http.HandleFunc("/add/feed", addFeedItem)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addFeedItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var t []structs.FeedItem
	jsonFile, _ := ioutil.ReadFile("API/dummy_data/feed_item.json")
	errRead := json.Unmarshal(jsonFile, &t)
	if errRead != nil {
		panic(errRead)
	}

	body, _ := ioutil.ReadAll(r.Body)
	log.Println(r.Body)
	var tOne structs.FeedItem
	_ = json.Unmarshal(body, &tOne)
	if tOne.Likes == nil {
		tOne.Likes = []int{}
	}
	t = append(t, tOne)

	js, _ := json.Marshal(t)
	_ = ioutil.WriteFile("API/dummy_data/feed_item.json", js, 0644)

	w.Write([]byte(""))
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