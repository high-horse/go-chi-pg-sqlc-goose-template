package router

import (
	"net/http"
	userhandler "server/http/handlers/user_handler"
	"server/http/handlers/util"
	md "server/http/middleware"
	// "server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)


func InitRouter() http.Handler {

	router := chi.NewRouter()
	router.Use(md.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	registerUtilRoutes(v1Router)
	registerUserRoutes(v1Router)
	// registerFeedRoutes(v1Router, apiCfg)

	router.Mount("/v1", v1Router)

	return router
}


func registerUtilRoutes(r chi.Router) {
	r.Get("/health", util.HandlerReady)
	r.Get("/err", util.HandleErr)
}

func registerUserRoutes(r chi.Router) {

	r.Post("/register", userhandler.HandlerCreateUser)
	r.Post("/login", userhandler.HandlerLogin)
	r.With(md.JWTMiddleware).Post("/logout", userhandler.LogOut)
	r.With(md.JWTMiddleware).Get("/status", userhandler.CheckStatus)
	// r.Get("/user", dbCfg.middlewareAuth(dbCfg.handlerGetUser))
	// r.Get("/posts", dbCfg.middlewareAuth(dbCfg.handlerGetPostsForUser))
}

// func registerFeedRoutes(r chi.Router, apiCfg config.ApiConfig) {
// 	r.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
// 	r.Get("/feed", apiCfg.handlerGetFeed)
// 	r.Post("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
// 	r.Get("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
// 	r.Delete("/feed_follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))
// }