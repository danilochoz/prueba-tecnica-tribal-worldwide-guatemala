package models

type Config struct {
	URL string
}

type Result struct {
	Results []User `json:"results"`
}

type User struct {
	Gender   string   `json:"gender"`
	Name     UserName `json:"name"`
	Location Location `json:"location"`
	Email    string   `json:"email"`
	Login    Login    `json:"login"`
}

type UserName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Login struct {
	UUID string `json:"uuid"`
}
