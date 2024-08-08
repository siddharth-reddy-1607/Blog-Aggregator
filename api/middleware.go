package api

import (
	"net/http"
	"strings"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/utils"
)

//authedHandlers are those which need be passed as inputs to authMiddleware. They will only be executed if the authencation succeeds
type authedHandler func(http.ResponseWriter, *http.Request, *database.User)

func (apiConfig *APIConfig) AuthMiddleware(next authedHandler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
        apiKey,found := strings.CutPrefix(r.Header.Get("Authorization"),"ApiKey ")
        if !found{
            utils.RespondWithError(w,http.StatusUnauthorized,"Invalid Header Format : No Authorization Header Found / 'ApiKey ' Prefix is missing")
            return
        }
        user,err := apiConfig.DBQueries.GetUserByAPIKey(r.Context(),apiKey)
        if err != nil{
            utils.RespondWithError(w,http.StatusUnauthorized,"User not found in database")
            return
        }
        next(w,r,&user)
    })
}
