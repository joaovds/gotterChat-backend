package dtos

type CreateUserInputDTO struct {
  ID string `json:"id"`
  Nickname string `json:"nickname"`
  Password string `json:"password"`
}
