package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/api"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
)

func main(){
    if err := godotenv.Load(); err != nil{
        log.Fatalf("Could not load .env file : %v\n",err)
    }

    PORT := os.Getenv("PORT")
    DBSTRING := os.Getenv("DBSTRING")

    db,err := sql.Open("postgres",DBSTRING)
    if err != nil{
        log.Fatalf("Error while setting up DB Connection : %v\n",err)
    }
    dbQueries := database.New(db)
    apiConfig := api.APIConfig{DBQueries: dbQueries}

    serveMux := http.NewServeMux()
    server := http.Server{
        Handler: serveMux,
        Addr: ":"+PORT,
    }

    serveMux.Handle("GET /v1/healthz",api.HealthHandler())
    serveMux.Handle("GET /v1/error",api.ErrorHandler())

    serveMux.Handle("POST /v1/users",apiConfig.CreateUserHandler())
    serveMux.Handle("GET /v1/users",apiConfig.AuthMiddleware(apiConfig.GetUserHandler))

    serveMux.Handle("POST /v1/feeds",apiConfig.AuthMiddleware(apiConfig.CreateFeedHandler))
    serveMux.Handle("GET /v1/feeds",apiConfig.GetFeedsHandler())

    serveMux.Handle("POST /v1/feed_follows",apiConfig.AuthMiddleware(apiConfig.CreateFeedFollowHandler))
    serveMux.Handle("GET /v1/feed_follows",apiConfig.AuthMiddleware(apiConfig.GetFeedFollowsHandler))
    serveMux.Handle("DELETE /v1/feed_follows/{feedFollowID}",apiConfig.AuthMiddleware(apiConfig.DeleteFeedFollowHandler))

    log.Printf("Listening and Serving on port %v",PORT)
    server.ListenAndServe()
}
