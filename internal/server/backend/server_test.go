//go:build integration

package backend

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/itimofeev/social-network/internal/app/backend"
	"github.com/itimofeev/social-network/internal/repository/pg"
	"github.com/itimofeev/social-network/internal/server/backend/gen/api"
)

func TestCreateGetLoginUser(t *testing.T) {
	ctx := context.Background()

	client, clearFn := prepareClient(t)
	defer clearFn()

	var userRegistered *api.UserRegisterPostOK
	t.Run("register_user", func(t *testing.T) {
		registerResp, err := client.UserRegisterPost(ctx, api.OptUserRegisterPostReq{
			Value: api.UserRegisterPostReq{
				FirstName:  api.NewOptString("Macaulay"),
				SecondName: api.NewOptString("Culkin"),
				Birthdate:  api.NewOptBirthDate(api.BirthDate(time.Now())),
				Biography:  api.NewOptString("hello, there"),
				City:       api.NewOptString("Москва"),
				Password:   api.NewOptString("123456"),
			},
			Set: true,
		})
		require.NoError(t, err)
		require.IsType(t, registerResp, new(api.UserRegisterPostOK))
		userRegistered = registerResp.(*api.UserRegisterPostOK)
	})

	var userGot *api.User
	t.Run("get_user", func(t *testing.T) {
		gotUserResp, err := client.UserGetIDGet(ctx, api.UserGetIDGetParams{
			ID: api.UserId(userRegistered.UserID.Value),
		})
		require.NoError(t, err)
		require.IsType(t, gotUserResp, new(api.User))
		userGot = gotUserResp.(*api.User)

		require.Equal(t, userRegistered.UserID, api.NewOptString(string(userGot.ID.Value)))
	})

	var loginToken string
	t.Run("login", func(t *testing.T) {
		loginResp, err := client.LoginPost(ctx, api.OptLoginPostReq{
			Value: api.LoginPostReq{
				ID:       api.NewOptUserId(api.UserId(userRegistered.UserID.Value)),
				Password: api.NewOptString("123456"),
			},
			Set: true,
		})
		require.NoError(t, err)
		require.IsType(t, loginResp, new(api.LoginPostOK))
		loginToken = loginResp.(*api.LoginPostOK).Token.Value
		require.NotEmpty(t, loginToken)
	})

	t.Run("invalid_user_login", func(t *testing.T) {
		loginResp, err := client.LoginPost(ctx, api.OptLoginPostReq{
			Value: api.LoginPostReq{
				ID:       api.NewOptUserId(api.UserId("invalid_user_id")),
				Password: api.NewOptString("incorrect_password"),
			},
			Set: true,
		})
		require.NoError(t, err)
		require.IsType(t, loginResp, new(api.LoginPostNotFound))
	})

	t.Run("incorrect_password_on_login", func(t *testing.T) {
		loginResp, err := client.LoginPost(ctx, api.OptLoginPostReq{
			Value: api.LoginPostReq{
				ID:       api.NewOptUserId(api.UserId(userRegistered.UserID.Value)),
				Password: api.NewOptString("incorrect_password"),
			},
			Set: true,
		})
		require.NoError(t, err)
		require.IsType(t, loginResp, new(api.LoginPostBadRequest))
	})

	t.Run("get_not_existing_user", func(t *testing.T) {
		gotUserResp, err := client.UserGetIDGet(ctx, api.UserGetIDGetParams{
			ID: api.UserId("not_exists"),
		})
		require.NoError(t, err)
		require.IsType(t, gotUserResp, new(api.UserGetIDGetNotFound))
	})

	t.Run("search_user", func(t *testing.T) {
		searchUserResp, err := client.UserSearchGet(ctx, api.UserSearchGetParams{
			FirstName: "Maca",
			LastName:  "Culk",
		})
		require.NoError(t, err)

		require.IsType(t, searchUserResp, new(api.UserSearchGetOKApplicationJSON))
		users := searchUserResp.(*api.UserSearchGetOKApplicationJSON)
		require.Contains(t, *users, *userGot)
	})

}

func prepareClient(t *testing.T) (*api.Client, func()) {
	ctx := context.Background()
	repo, err := pg.New(ctx, pg.Config{DSN: os.Getenv("PG_REPOSITORY_DSN")})
	require.NoError(t, err)

	_app, err := backend.New(backend.Config{
		PGRepository:    repo,
		PasetoSecretKey: "5468ac74e23ea5c297413a3020af91601f22c82e77aa89cca4e8fb4ec28fb300",
	})
	require.NoError(t, err)

	server, err := NewServer(Config{
		Domain:          "http://localhost:8080",
		Version:         "1.0.0",
		Port:            "8080",
		ReadTimeout:     time.Second,
		WriteTimeout:    time.Second,
		ShutdownTimeout: time.Second * 10,
		App:             _app,
	})

	serverHandler := server.srv.Handler

	client, err := api.NewClient("http://localhost:8080/api/v1", nil, api.WithClient(httpClientMock{handler: serverHandler}))
	require.NoError(t, err)

	return client, func() {
		repo.Close()
	}
}

type httpClientMock struct {
	handler http.Handler
}

func (h httpClientMock) Do(r *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	h.handler.ServeHTTP(recorder, r)
	return recorder.Result(), nil
}
