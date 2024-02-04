package requests

type CreateUser struct {
	Name         string ` json:"name" validate:"required"`
	Email        string ` json:"email" validate:"email,required"`
	Password     string ` json:"password" validate:"required"`
	ConfPassword string `json:"confPassword" validate:"required"`
}
