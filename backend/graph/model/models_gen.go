// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	To        *User  `json:"to"`
	From      *User  `json:"from"`
	Timestamp string `json:"timestamp"`
}

type Mutation struct {
}

type NewMessage struct {
	Text       string `json:"text"`
	UserToID   string `json:"userToId"`
	UserFromID string `json:"userFromId"`
}

type Query struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
