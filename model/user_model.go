package model

type UserModel struct {
	Id       uint   `json:"id"`
	HashId   string `json:"hash_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
