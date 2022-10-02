package main

import (
	"tes/handler"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// セッションCookieの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	// ログイン用のhandler
	router.POST("/login", handler.Login)
	// 認証済のみアクセス可能なグループ
	authUserGroup := router.Group("/auth")
	authUserGroup.Use(middleware.LoginCheckMiddleware())
	{
		authUserGroup.GET("/getSample", handler.getSample)
	}
	router.Run()
}
