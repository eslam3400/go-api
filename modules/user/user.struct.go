package user_module

type SUser struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
