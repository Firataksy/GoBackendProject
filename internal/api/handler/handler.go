package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/my/repo/internal/api"

	mid "github.com/my/repo/internal/api/middleware"
)

func Handler() error {
	rc, err := api.ConnectRedis()
	if err != nil {
		return err
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/signup", rc.SignUp)
	mux.HandleFunc("/login", rc.Login)
	mux.Handle("/userinfo", mid.TokenMiddleware(http.HandlerFunc(rc.GetUserData)))
	mux.Handle("/updateuser", mid.TokenMiddleware(http.HandlerFunc(rc.UpdateUserData)))
	mux.HandleFunc("/match", rc.Match)
	mux.Handle("/leaderboard", mid.TokenMiddleware(http.HandlerFunc(rc.ListLeaderBoard)))
	mux.HandleFunc("/simulation", rc.Simulation)
	mux.Handle("/usersearch", mid.TokenMiddleware(http.HandlerFunc(rc.UserSearch)))
	mux.Handle("/friendrequest", mid.TokenMiddleware(http.HandlerFunc(rc.FriendRequest)))
	mux.Handle("/friendrequestlist", mid.TokenMiddleware(http.HandlerFunc(rc.FriendRequestList)))
	mux.Handle("/friendacceptreject", mid.TokenMiddleware(http.HandlerFunc(rc.FriendAcceptReject)))
	mux.Handle("/friendlist", mid.TokenMiddleware(http.HandlerFunc(rc.FriendList)))
	http.Handle("/", mux)
	fmt.Println("http listen started")
	return nil
}
