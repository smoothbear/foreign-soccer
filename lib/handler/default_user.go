package handler

import (
	"encoding/json"
	"log"
	"net/http"

	mysqlcode "github.com/VividCortex/mysqlerr"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"smooth-bear.live/lib/model"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (h _default) CreateNewUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var request CreateUserRequest
	err := decoder.Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400: Bad Request. Please check your request body."))

		log.Println("JSON Parse Error: ", err)

		return
	}

	accessor, err := h.accessManage.BeginTx()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Something went wrong."))

		log.Println("Accessor Error: ", err)

		return
	}

	pwHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		accessor.Rollback()

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Unable to hash password."))

		return
	}

	_, err = accessor.CreateUser(&model.User{
		Email:    request.Email,
		Password: string(pwHash),
		Name:     request.Password,
	})

	switch assertedError := err.(type) {
	case nil:
		break

	case validator.ValidationErrors:
		accessor.Rollback()

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400: Unable to insert data. (Invalid Data)"))

		return

	case *mysql.MySQLError:
		accessor.Rollback()

		switch assertedError.Number {
		case mysqlcode.ER_DUP_ENTRY:
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("409: User is already exist."))

			return

		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Unexpected server error."))

			return
		}
	}

	accessor.Commit()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201: Register user is successed."))

	return
}
