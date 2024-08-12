package api

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/utils"
)

func (apiConfig *APIConfig) CreateFeedHandler(w http.ResponseWriter,r *http.Request,user *database.User){
    requestJSON := struct{Name string `json:"name"`
                          Url string `json:"url"`}{}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&requestJSON); err != nil{
        utils.RespondWithError(w,http.StatusInternalServerError,jsonDecodeError.Error())
        return
    }
    curTime := time.Now()
    feed,err := apiConfig.DBQueries.CreateFeed(r.Context(),database.CreateFeedParams{ID: uuid.New(),
                                                                                     Name: requestJSON.Name,
                                                                                     Url: requestJSON.Url,
                                                                                     CreatedAt: curTime,
                                                                                     UpdatedAt: curTime,
                                                                                     UserID: (*user).ID})
    if err != nil{
        utils.RespondWithError(w,http.StatusInternalServerError,"Error while creating feed")
        return
    }
    curTime = time.Now()
    feedFollow,err := apiConfig.DBQueries.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{ID: uuid.New(),
                                                                                                     UserID: (*user).ID,
                                                                                                     FeedID: feed.ID,
                                                                                                     CreatedAt: curTime,
                                                                                                     UpdatedAt: curTime})
    responseJSON := struct{Feed *Feed `json:"feed"`
                           FeedFollow *FeedFollow `json:"feed_follow"`}{Feed : databaseFeedToFeed(&feed),
                                                                        FeedFollow: databaseFeedFollowToFeedFollow(&feedFollow)}
    utils.RespondWithJSON(w,http.StatusCreated,responseJSON)
}

func (apiConfig *APIConfig) GetFeedsHandler() http.Handler{
    return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
        databaseFeeds,err := apiConfig.DBQueries.GetFeeds(r.Context())
        if err != nil{
            utils.RespondWithError(w,http.StatusInternalServerError,"Error while getting feeds")
            return
        }
        feeds := []*Feed{}
        for _,dbFeed := range databaseFeeds{
            feeds = append(feeds,databaseFeedToFeed(&dbFeed))
        }
        utils.RespondWithJSON(w,http.StatusOK,feeds)
    })
}
