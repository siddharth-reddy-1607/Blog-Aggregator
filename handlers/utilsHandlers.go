package handlers

import (
    "net/http"
    "github.com/siddharth-reddy-1607/Blog-Aggregator/utils"
)

func HealthHandler() http.Handler{
    return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
        response := struct {Status string `json:"status"`}{Status : "ok"}
        utils.RespondWithJSON(w,http.StatusOK,response)
    })
}

func ErrorHandler() http.Handler{
    return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
        utils.RespondWithError(w,http.StatusInternalServerError,"Internal Server Error")
    })
}
