package domain

type User struct {
	Nickname    string   `json:"nickname"`
	Password    string   `json:"password"`
	Gender      string   `json:"gender"`
	PhoneNumber string   `json:"phoneNumber"`
	Interests   []string `json:"interests"`
}
