package structs

import "time"

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

type FeedItem struct {
	Id        int       `json:"id"`
	User      User      `json:"user"`
	Milestone Milestone `json:"milestone"`
	Likes     []int     `json:"likes-array"`
}

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	ProfilePic string `json:"profilepic"`
}

type Goal struct {
	Id          int       `json:"id"`
	User        User      `json:"user"`
	Deadline    time.Time `json:"deadline-date-time"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type Milestone struct {
	Id         int    `json:"id"`
	Goal       Goal   `json:"goal"`
	Title      string `json:"title"`
	Percentage int    `json:"percentage"`
}

type Comment struct {
	Id       int       `json:"id"`
	User     User      `json:"user"`
	Feeditem FeedItem  `json:"feeditem"`
	Date     time.Time `json:"date"`
	Content  string    `json:"content"`
}
