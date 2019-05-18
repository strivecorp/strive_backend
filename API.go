package main

import (
	"net/http"
	"time"
)

func APIHandler() {
	http.ListenAndServe(":8080", nil)
}

type FeedItem struct {
	id int
	person Person
	goal Goal
	likes []Person
	comments []Comment
}

type Person struct {
	id int
	firstName string
	lastName string
	email string
	goals []Goal
}

type Goal struct {
	id int
	deadline time.Time
	title string
	description string
	milestones []Milestone
}

type Milestone struct {
	id int
	title string
	percentage int
}

type Comment struct {
	id int
	person Person
	date time.Time
	content string
}