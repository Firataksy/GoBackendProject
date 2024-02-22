package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() error {
	rc, err := ConnectRedis()
	if err != nil {
		return err
	}
	mux := mux.NewRouter()
	mux.HandleFunc("/signup", rc.SignUp)
	mux.HandleFunc("/login", rc.Login)
	mux.Handle("/userinfo", rc.tokenMiddleware(http.HandlerFunc(rc.GetUserData)))
	mux.Handle("/updateuser", rc.tokenMiddleware(http.HandlerFunc(rc.UpdateUserData)))
	mux.HandleFunc("/match", rc.Match)
	mux.Handle("/leaderboard", rc.tokenMiddleware(http.HandlerFunc(rc.ListLeaderBoard)))
	mux.HandleFunc("/simulation", rc.Simulation)
	mux.Handle("/usersearch", rc.tokenMiddleware(http.HandlerFunc(rc.UserSearch)))
	mux.Handle("/friendrequest", rc.tokenMiddleware(http.HandlerFunc(rc.FriendRequest)))
	mux.Handle("/friendrequestlist", rc.tokenMiddleware(http.HandlerFunc(rc.FriendRequestList)))
	mux.Handle("/friendacceptreject", rc.tokenMiddleware(http.HandlerFunc(rc.FriendAcceptReject)))
	mux.Handle("/friendlist", rc.tokenMiddleware(http.HandlerFunc(rc.FriendList)))
	http.Handle("/", mux)
	fmt.Println("http listen started")
	return nil
}
