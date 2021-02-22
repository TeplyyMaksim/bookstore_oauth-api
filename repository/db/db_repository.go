package db

import (
	"github.com/TeplyyMaksim/bookstore_oauth-api/clents/cassandra"
	"github.com/TeplyyMaksim/bookstore_oauth-api/domain/access_token"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)


func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.HttpError)
	Create(access_token.AccessToken)  *errors_utils.HttpError
	UpdateExpirationTime(access_token.AccessToken)  *errors_utils.HttpError
}

type dbRepository struct {

}

func (*dbRepository) GetById(id string) (*access_token.AccessToken, *errors_utils.HttpError) {
	var at access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&at.AccessToken,
		&at.UserId,
		&at.ClientId,
		&at.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors_utils.NewNotFoundError("No access token found")
		}

		return nil, errors_utils.NewInternalServerError(err.Error())
	}

	return &at, nil
}

func (*dbRepository) Create(token access_token.AccessToken) *errors_utils.HttpError {
	if err := cassandra.GetSession().Query(
		queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires,
	).Exec(); err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}

	return nil
}

func (*dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors_utils.HttpError {
	if err := cassandra.GetSession().Query(
		queryUpdateExpires,
		token.AccessToken,
		token.Expires,
	).Exec(); err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}

	return nil
}