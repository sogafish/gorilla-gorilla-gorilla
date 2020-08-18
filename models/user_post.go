package models

type UserPost struct {
	ID    int    `json:"id"`
	TITLE int    `json:"title"`
	Body  string `json:"body"`
}

func GetPost() UserPost {
	var userPost UserPost

	return userPost
}

func GetPosts() []UserPost {
	var userPosts []UserPost

	return userPosts
}
