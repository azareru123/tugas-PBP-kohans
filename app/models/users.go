package models

type User struct {
	ID       int    `form : "id" json:"id"`
	Name     string `form : "name" json:"name"`
	Age      int    `form : "age" json:"age"`
	Address  string `form : "address" json:"address"`
	Email    string `form : "email" json:"email"`
	Password string `form : "password" json:"password"`
}

type UserResponse struct {
	Status  int    `form : "status" json:"status"`
	Message string `form : "Message" json:"Message"`
	Data    []User `form : "Data" json:"Data"`
}
