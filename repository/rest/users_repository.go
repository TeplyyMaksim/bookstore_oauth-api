package rest

import (
	"fmt"
	"github.com/TeplyyMaksim/bookstore_oauth-api/domain/user"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/imroc/req"
	"time"
)

var (
	usersRestClient *req.Req
)

func init() {
	r := req.New()
	r.SetTimeout(100 * time.Millisecond)

	usersRestClient = r
}


type RestUsersRepository interface {
	LoginUser(string, string) (*user.User, *errors_utils.HttpError)
}

type usersRepository struct {}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (*usersRepository) LoginUser(email string, password string) (*user.User, *errors_utils.HttpError) {
	requestBody := user.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response, err := usersRestClient.Post("http://localhost:8000/users/login", req.BodyJSON(&requestBody))
	if err != nil {
		return nil, errors_utils.NewInternalServerError(err.Error())
	}


	if response.Response().StatusCode > 299 {
		var error errors_utils.HttpError

		err := response.ToJSON(&error)
		fmt.Println(err, error)
		if err != nil {
			return nil, errors_utils.NewInternalServerError(err.Error())
		}

		return nil, &error
	}

	var user user.User
	if err = response.ToJSON(&user); err != nil {
		return nil, errors_utils.NewInternalServerError(err.Error())
	}

	return &user, nil
}

