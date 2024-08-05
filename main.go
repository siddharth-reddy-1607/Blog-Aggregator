package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
    "github.com/siddharth-reddy-1607/Blog-Aggregator/handlers"
)

func main(){
    if err := godotenv.Load(); err != nil{
        log.Fatalf("Could not load .env file : %v\n",err)
    }
    PORT := os.Getenv("PORT")
    serveMux := http.NewServeMux()
    server := http.Server{
        Handler: serveMux,
        Addr: ":"+PORT,
    }

    serveMux.Handle("GET /v1/healthz",handlers.HealthHandler())
    serveMux.Handle("GET /v1/error",handlers.ErrorHandler())

    log.Printf("Listening and Serving on port %v",PORT)
    server.ListenAndServe()
}
