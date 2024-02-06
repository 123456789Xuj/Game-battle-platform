package models

type NewUser struct {
	ID          int
	Username    string `json:"Username"`
	PhoneNumber int    `json:"phone_number"`
	Password    string `json:"password"`
}
type GameFriend struct {
	ID       int
	Username string `json:"username"`
}
type GameRoom struct {
	ID           int
	Username     string `json:"username"`
	MainUsername string `json:"mainUsername"`
	Ready        bool   `json:"ready"`
}
type Chess struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type League struct {
	ID       int
	UserName string
	Integral int
}
