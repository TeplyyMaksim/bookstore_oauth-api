package access_token

import (
	"fmt"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"strings"
	"time"
)

const (
	expirationTime = 24 * time.Hour
	grandTypePassword = "PASSWORD"
	grandTypeClientCredentials = "CLIENT_CREDENTIALS"
)

type AccessToken struct {
	AccessToken 	string 		`json:"access_token"`
	UserId 			int64 		`json:"user_id"`
	ClientId		int64		`json:"client_id"`
	Expires 		int64		`json:"expires"`
}

type AccessTokenRequest struct {
	GrandType 		string 		`json:"grand_type"`
	Scope			string 		`json:"scope"`

	// Used for password grand_type
	Username 		string 		`json:"username"`
	Password 		string 		`json:"password"`

	// Used for client_credentials grand_type
	ClientId 		string 		`json:"client_id"`
	ClientSecret 	string 		`json:"client_secret"`
}

func (request *AccessTokenRequest) Validate() *errors_utils.HttpError {
	switch request.GrandType {
	case grandTypeClientCredentials:
		return nil
	case grandTypePassword:
		return nil
	default:
		return errors_utils.NewBadRequestError("Invalid grand_type parameter")
	}
}



func (aT *AccessToken) Validate() *errors_utils.HttpError {
	aT.AccessToken = strings.TrimSpace(aT.AccessToken)

	if aT.AccessToken == ""  {
		return errors_utils.NewBadRequestError("invalid access token id")
	}

	if aT.UserId <= 0 {
		return errors_utils.NewBadRequestError("invalid user id")
	}

	if aT.ClientId <= 0 {
		return errors_utils.NewBadRequestError("invalid client id")
	}

	if aT.Expires <= 0 {
		return errors_utils.NewBadRequestError("invalid expiraTion time")
	}

	return nil
}

func GetNewAccessToken () *AccessToken {
	return &AccessToken{
		AccessToken: "",
		UserId:      0,
		ClientId:    0,
		Expires:     time.Now().UTC().Add(expirationTime).Unix(),
	}
}

func (aT AccessToken) IsExpired() bool {
	fmt.Println(time.Now().UTC().Unix(), aT.Expires)
	return time.Now().UTC().Unix() > aT.Expires
}