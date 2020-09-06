package model

type User struct {
	ID       string `json:"_id" bson:"_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UseRespoonse struct {
	ID       string `json:"_id" bson:"_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
