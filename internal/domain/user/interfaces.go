package user_domain

type CreateUserUseCase interface {
  execute(user *User) error
}
