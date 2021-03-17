package auth

import (
	"fmt"
	"github.com/damonchen/org-star/internal/config"
	"github.com/go-chi/chi"
	"net/http"
)

// Login login
func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("login\n"))
}

// Logout logout
func Logout(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("logout\n"))
}

// Redirect redirect
func Redirect(w http.ResponseWriter, req *http.Request) {
	// 返回一个跳转的链接
	// https://github.com/login/oauth/authorize?
	//  client_id=7e015d8ce32370079895&
	//  redirect_uri=http://localhost:8080/oauth/redirect
	code := chi.URLParam(req, "code")
	fmt.Printf("get code from %s\n", code)

	w.Write([]byte("redirect"))
}

func Authorize(w http.ResponseWriter, req *http.Request) {
	cfg := config.GetCfg()
	clientID := cfg.Github.AppID
	server := cfg.Server

	fmt.Printf("server %s\n", server)

	redirectURI := fmt.Sprintf("%s/oauth/redirect", server)
	uri := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		clientID, redirectURI)

	w.Header().Set("Location", uri)
	w.WriteHeader(301)
	return
}
