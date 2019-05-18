package API

import (
	"encoding/json"
	"fmt"
	"go/types"
	"io/ioutil"
	"log"
	"net/http"
	"strive_backend/API/structs"
)

func APIHandler() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Users data below")
		fmt.Fprintln(w, getUsers())
		fmt.Fprintln(w, "Goals data below")
		fmt.Fprintln(w, getGoals())
		fmt.Fprintln(w, "Feed data below")
		fmt.Fprintln(w, getFeedItems())
		fmt.Fprintln(w, "Milestone data below")
		fmt.Fprintln(w, getMilestones())
		fmt.Fprintln(w, "Comments data below")
		fmt.Fprintln(w, getComments())
		fmt.Fprintln(w, readJsonById("goals.json", 0))
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
		return "Not found"
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
		return "Not found"
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
		return "Not found"
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
		return "Not found"
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
		return "Not found"
	default:
		panic(types.Interface{})
	}

}
