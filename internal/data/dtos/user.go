package dtos

type CreateUserInputDTO struct {
  ID string `json:"id" bson:"id"`
  Nickname string `json:"nickname" bson:"nickname"`
  Password string `json:"password" bson:"password"`
}
