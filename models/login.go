package models

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
type StandardResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
type LoginResponse struct {
	Token string `json:"token"`
	Admin Admin  `json:"admin"`
}
