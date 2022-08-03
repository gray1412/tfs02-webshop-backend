package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project-backend/database"

	"project-backend/model"

	bcrypt "project-backend/util/bcrypt"
	jwt "project-backend/util/jwt"
	response "project-backend/util/response"
	"time"
)

const (
	UserExist      = "User already exists"
	NewUser        = "New user added"
	NoUser         = "User doesn't exist"
	UserNotActive  = "Account has been deactivated"
	WrongPassword  = "Wrong password"
	LoginSuccess   = "Logged in"
	UserReactivate = "User account reactivated"
)

var db = database.ConnectDB()

func SignUp(w http.ResponseWriter, r *http.Request) {
	newUser := model.User{}

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		response.RespondWithJSON(w, 400, 0, err.Error(), nil)
		return
	}


	foundUser := model.User{}
	result := db.First(&foundUser, "email = ?", newUser.Email)


	if result.RowsAffected != 0 {
		switch foundUser.Active {
		case 0: // Account has been deleted -> update db with request body
			// Encrypt password before saving to db (will be moved to frontend later)
			hashPassword, err := bcrypt.HashPassword(newUser.Password)
			if err != nil {
				fmt.Println("error hasing password", err)
			}
			foundUser.Password = hashPassword
			foundUser.Name = newUser.Name
			foundUser.Phone = newUser.Phone
			foundUser.Address = newUser.Address
			foundUser.Active = 2
			foundUser.UpdatedAt = time.Now()
			db.Save(&foundUser)
			response.RespondWithJSON(w, 201, 1, UserReactivate, nil)
			return
		case 1: // User has been deactivated
			response.RespondWithJSON(w, 400, 0, UserNotActive, nil)
			return
		default: // User exists
			response.RespondWithJSON(w, 400, 0, UserExist, nil)
			return
		}
	}

	newUser.Active = 2

	if result := db.Create(&newUser); result.Error != nil {
		response.RespondWithJSON(w, 400, 0, result.Error.Error(), nil)
		return
	}
	newUser.Password = ""
	response.RespondWithJSON(w, 200, 1, NewUser, &newUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	credentials := model.User{} //hold user login credentials from request body
	user := model.User{}        // hold user data from db

	// Get login data from request
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response.RespondWithJSON(w, 400, 0, err.Error(), nil)
		return
	}

	// Check if requested email exists in db
	result := db.Where("email = ?", credentials.Email).First(&user)
	if result.Error != nil {
		response.RespondWithJSON(w, 400, 0, NoUser, nil)
		return
	}

	// Check if account is active
	if user.Active != 2 {
		response.RespondWithJSON(w, 400, 0, UserNotActive, nil)
		return
	}

	match := bcrypt.CheckPasswordHash(credentials.Password, user.Password)

	if !match {
		response.RespondWithJSON(w, 401, 0, WrongPassword, nil)
		return
	}

	//Generate jwt
	tokenString, err := jwt.CreateToken(w, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.RespondWithJSON(w, 500, 0, err.Error(), nil)
	}

	response.RespondWithJSON(w, 200, 1, LoginSuccess, map[string]string{"token": tokenString})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("email")
	user := model.User{}
	db := database.ConnectDB()
	result := db.Where("email = ?", email).Omit("password").First(&user)
	if result.Error != nil {
		response.RespondWithJSON(w, 400, 0, NoUser, nil)
	}
	response.RespondWithJSON(w, 200, 1, "", user)
}
