package webserver_dtos

type CreateUserRequestDTO struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
