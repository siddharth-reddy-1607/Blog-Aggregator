package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/utils"
)

func (apiConfig *APIConfig) GetPosts(w http.ResponseWriter,r *http.Request,user *database.User){
    queries := r.URL.Query()
    limit,err := strconv.Atoi(queries.Get("limit"))
    if err != nil{
        limit = 5
    }
    dbPosts,err := apiConfig.DBQueries.GetPosts(r.Context(),database.GetPostsParams{UserID: (*user).ID,
                                                                                    Limit: int32(limit)})
    if err != nil{
        errMsg := fmt.Sprintf("Error while getting user's posts: %v\n",err)
        utils.RespondWithError(w,http.StatusInternalServerError,errMsg)
        return
    }
    posts := []*Post{}
    for _,dbPost := range(dbPosts){
        posts = append(posts,databasePostToPost(&dbPost))
    }
    utils.RespondWithJSON(w,http.StatusOK,posts)
}
