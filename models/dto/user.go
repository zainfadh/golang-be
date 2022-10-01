package dto

type LoginRequestDto struct {
	Email    string `json:"email" example:"abc@gmail.com"`
	Password string `json:"password" example:"test123"`
}

type GetDataByID struct {
	ID int `json:"ID"`
}
