package access_token

import (
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
)


type Repository interface {
	GetById(string) (*AccessToken, *errors_utils.HttpError)
	Create(token AccessToken) *errors_utils.HttpError
	UpdateExpirationTime(token AccessToken) *errors_utils.HttpError
}

type Service interface {
	GetById(string) (*AccessToken, *errors_utils.HttpError)
	Create(token AccessToken) *errors_utils.HttpError
	UpdateExpirationTime(AccessToken) *errors_utils.HttpError
}

type service struct {
	repository Repository
}
func (s *service) GetById(accessTokenId string) (*AccessToken, *errors_utils.HttpError) {
	accessToken, err := s.repository.GetById(accessTokenId)

	if len(accessTokenId) == 0 {
		return nil, errors_utils.NewBadRequestError("invalid access token id")
	}

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

//func (s *service) Create(request *AccessTokenRequest) *errors_utils.HttpError {
//	if err := request.Validate(); err != nil {
//		return err
//	}
//
//	//s.repository
//	//user, err := Login
//
//	//
//	return s.repository.Create(at)
//}

func (s *service) Create(at AccessToken) *errors_utils.HttpError {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors_utils.HttpError {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(at)
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
