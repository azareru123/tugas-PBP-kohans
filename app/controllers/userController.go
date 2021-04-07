package controllers

import (
	"revelkohans/app"
	"revelkohans/app/models"
	"strconv"

	"github.com/revel/revel"
)

type UserController struct {
	*revel.Controller
}

func (u UserController) GetAllUsers() revel.Result {
	var users []models.User

	query := "SELECT id, email, name, age, address FROM users"
	rows, err := app.DB.Query(query)

	if err != nil {
		return u.RenderJSON(sendResponse(400, err.Error(), nil))
	}

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Age, &user.Address); err != nil {
			revel.AppLog.Error(err.Error())
		} else {
			users = append(users, user)
		}
	}

	return u.RenderJSON(sendResponse(200, "Success", users))
}

func (u UserController) InsertUser() revel.Result {
	err := u.Request.ParseForm()
	if err != nil {
		return u.RenderJSON(sendResponse(400, err.Error(), nil))
	}

	name := u.Request.Form.Get("name")
	email := u.Request.Form.Get("email")
	password := u.Request.Form.Get("password")
	age, _ := strconv.Atoi(u.Request.Form.Get("age"))
	address := u.Request.Form.Get("address")

	_, errQuery := app.DB.Exec("INSERT INTO users(name, password, age, address, email) VALUES (?, ?, ?, ?, ?)",
		name,
		password,
		age,
		address,
		email,
	)

	if errQuery != nil {
		return u.RenderJSON(sendResponse(400, errQuery.Error(), nil))
	} else {
		return u.RenderJSON(sendResponse(200, "Success", nil))
	}
}

func (u UserController) UpdateUser(id int) revel.Result {
	err := u.Request.ParseForm()
	if err != nil {
		return u.RenderJSON(sendResponse(400, err.Error(), nil))
	}

	name := u.Request.Form.Get("name")
	email := u.Request.Form.Get("email")
	password := u.Request.Form.Get("password")
	age, _ := strconv.Atoi(u.Request.Form.Get("age"))
	address := u.Request.Form.Get("address")

	_, errQuery := app.DB.Exec("UPDATE users SET name=?, password=?, age=?, address=?, email=? WHERE id=?",
		name,
		password,
		age,
		address,
		email,
		id,
	)

	if errQuery != nil {
		return u.RenderJSON(sendResponse(400, errQuery.Error(), nil))
	} else {
		return u.RenderJSON(sendResponse(200, "Success", nil))
	}
}

func (u UserController) DeleteUser(id int) revel.Result {
	_, errQuery := app.DB.Exec("DELETE FROM users WHERE id=?", id)

	if errQuery != nil {
		return u.RenderJSON(sendResponse(400, errQuery.Error(), nil))
	} else {
		return u.RenderJSON(sendResponse(200, "Success", nil))
	}
}

func sendResponse(status int, message string, data []models.User) models.UserResponse {
	var response models.UserResponse
	response.Status = status
	response.Message = message
	response.Data = data
	return response
}

func (u UserController) GetAllUsersHtml() revel.Result {
	var users []models.User

	query := "SELECT id, email, name, age, address FROM users"
	rows, err := app.DB.Query(query)

	if err != nil {
		return u.RenderJSON(sendResponse(400, err.Error(), nil))
	}

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Age, &user.Address); err != nil {
			revel.AppLog.Error(err.Error())
		} else {
			users = append(users, user)
		}
	}

	return u.Render(users)
}
