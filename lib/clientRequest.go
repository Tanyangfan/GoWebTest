package lib

type BaseRequest struct {
	Token string `json:"token"`
}

type QueryUserRequest struct {
	BaseRequest
	ID int `json:"id"`
}

type RegistUserRequest struct {
	BaseRequest
	Name string `json:"name"`
}
