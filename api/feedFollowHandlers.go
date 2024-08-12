package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/google/uuid"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/utils"
)

func (apiConfig *APIConfig) CreateFeedFollowHandler(w http.ResponseWriter,r *http.Request,user *database.User){
    requestJSON := struct{FeedID uuid.UUID `json:"feed_id"`}{}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&requestJSON); err != nil{
        utils.RespondWithError(w,http.StatusInternalServerError,jsonDecodeError.Error())
        return
    }
    curTime := time.Now()
    feedFollow,err := apiConfig.DBQueries.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{ID: uuid.New(),
                                                                                                       UserID: (*user).ID,
                                                                                                       FeedID: requestJSON.FeedID,
                                                                                                       CreatedAt: curTime,
                                                                                                       UpdatedAt: curTime})
    if err != nil{
        errMsg := fmt.Sprintf("Error while creating feed follow :%v",err.Error())
        utils.RespondWithError(w,http.StatusConflict,errMsg)
        return
    }
    utils.RespondWithJSON(w,http.StatusCreated,databaseFeedFollowToFeedFollow(&feedFollow))
}

func (apiConfig *APIConfig) GetFeedFollowsHandler(w http.ResponseWriter,r *http.Request,user *database.User){
    databaseFeedFollows,err := apiConfig.DBQueries.GetFeedFollows(r.Context(),(*user).ID)
    if err != nil{
        utils.RespondWithError(w,http.StatusInternalServerError,"Error while getting feed follows") 
        return
    }
    feedFollows := []*FeedFollow{}
    for _,dbFeedFollow := range(databaseFeedFollows){
        feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(&dbFeedFollow))
    }
    utils.RespondWithJSON(w,http.StatusOK,feedFollows)
}

func (apiConfig *APIConfig) DeleteFeedFollowHandler(w http.ResponseWriter,r *http.Request,user *database.User){
    ffID := r.PathValue("feedFollowID")
    feedFollowID,err := uuid.Parse(ffID)
    if err != nil{
        errMsg := fmt.Sprintf("Invalid Feed Follow ID: %v",err)
        utils.RespondWithError(w,http.StatusBadRequest,errMsg)
        return
    }
    if err := apiConfig.DBQueries.DeleteFeedFollow(r.Context(),feedFollowID); err != nil{
        utils.RespondWithError(w,http.StatusInternalServerError,"Error while deleting feed follow")
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
