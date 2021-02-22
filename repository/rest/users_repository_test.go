package rest

import (
	"github.com/TeplyyMaksim/bookstore_oauth-api/domain/user"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain (m *testing.M) {
	httpmock.ActivateNonDefault(usersRestClient.Client())
	defer httpmock.DeactivateAndReset()

	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	httpmock.RegisterResponder(
		"POST",
		"http://localhost:8000/users/login",
		func(request *http.Request) (*http.Response, error) {
			time.Sleep(1 * time.Second)
			return httpmock.NewJsonResponse(http.StatusOK, struct {}{})
		},
	)
	defer httpmock.Reset()


	repository := NewRepository()
	response, error := repository.LoginUser("email", "Password")


	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.True(
		t,
		strings.Contains(error.Message, "context deadline exceeded") ||
			strings.Contains(error.Message, "exceeded while awaiting headers"),
	)
}

func TestLoginErrorBadInterface(t *testing.T) {
	httpmock.RegisterResponder(
		"POST",
		"http://localhost:8000/users/login",
		httpmock.NewStringResponder(http.StatusInternalServerError, "Some str resp"),
	)


	repository := NewRepository()
	response, error := repository.LoginUser("email", "Password")


	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.Equal(t, http.StatusInternalServerError, error.Status)
}

func TestLoginBadCredentials(t *testing.T) {
	httpmock.RegisterResponder(
		"POST",
		"http://localhost:8000/users/login",
		httpmock.NewJsonResponderOrPanic(http.StatusNotFound, errors_utils.NewNotFoundError("Invalid credentials")),
	)


	repository := NewRepository()
	response, error := repository.LoginUser("email", "Password")


	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.Equal(t, http.StatusNotFound, error.Status)
	assert.Equal(t, "Invalid credentials", error.Message)
}

func TestLoginUserBadInterface(t *testing.T) {
	httpmock.RegisterResponder(
		"POST",
		"http://localhost:8000/users/login",
		httpmock.NewStringResponder(http.StatusOK, "Some str resp"),
	)


	repository := NewRepository()
	response, error := repository.LoginUser("email", "Password")


	assert.Nil(t, response)
	assert.NotNil(t, error)
	assert.Equal(t, http.StatusInternalServerError, error.Status)
}

func TestLoginSuccessful(t *testing.T) {
	httpmock.RegisterResponder(
		"POST",
		"http://localhost:8000/users/login",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, user.User{
			Id:        1,
			FirstName: "Maksym",
			LastName:  "Teplyy",
			Email:     "teplyy.maksim@gmail.com",
		}),
	)


	repository := NewRepository()
	response, error := repository.LoginUser("email", "Password")


	assert.Nil(t, error)
	assert.NotNil(t, response)
	assert.Equal(t, response.Id, 1)
	assert.Equal(t, response.FirstName, "Maksym")
	assert.Equal(t, response.LastName, "Teplyy")
	assert.Equal(t, response.Email, "teplyy.maksim@gmail.com")
}