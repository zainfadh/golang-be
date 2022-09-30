package dto

type LoginRequestDto struct {
	Username string `json:"username" example:"deddy"`
	Password string `json:"password" example:"0c91a43f8e1ec5fcba28f8a5a34532679305ca131302ad2a06218b47f30ced88"`
	Response string `json:"response"`
}

type GetDataByID struct {
	ID int64 `json:"ID"`
}
