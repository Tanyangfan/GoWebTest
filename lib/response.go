package lib

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserResponse struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type RegistResponse struct {
	UserResponse
}
