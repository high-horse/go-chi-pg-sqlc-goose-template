package router

import (
	"net/http"
	// "server/sql/database"
	md "server/middleware"
	"server/handlers/util"
	"server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)


func InitRouter(apiCfg models.DBConfig) http.Handler {

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
	// registerUserRoutes(v1Router, apiCfg)
	// registerFeedRoutes(v1Router, apiCfg)

	router.Mount("/v1", v1Router)

	return router
}


func registerUtilRoutes(r chi.Router) {
	r.Get("/health", util.HandlerReady)
	r.Get("/err", util.HandleErr)
}

// func registerUserRoutes(r chi.Router, apiCfg config.ApiConfig) {
// 	r.Post("/user", apiCfg.handlerCreateUser)
// 	r.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
// 	r.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))
// }

// func registerFeedRoutes(r chi.Router, apiCfg config.ApiConfig) {
// 	r.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
// 	r.Get("/feed", apiCfg.handlerGetFeed)
// 	r.Post("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
// 	r.Get("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
// 	r.Delete("/feed_follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))
// }