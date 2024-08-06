package api

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/utils"
)

func (apiConfig *APIConfig) CreateUserHandler() http.Handler{
    return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
        requestJSON := struct{Name string `json:"name"`}{}
        decoder := json.NewDecoder(r.Body)
        if err := decoder.Decode(&requestJSON); err != nil{
            utils.RespondWithError(w,http.StatusInternalServerError,jsonDecodeError.Error())    
            return
        }

        uuidStruct := uuid.NullUUID{UUID : uuid.New(),
                                    Valid: true}
        time := time.Now()
        user,err := apiConfig.DBQueries.CreateUser(r.Context(),
                                                   database.CreateUserParams{
                                                        ID : uuidStruct,
                                                        Name: requestJSON.Name,
                                                        CreatedAt: time,
                                                        UpdatedAt: time})
        if err != nil{
            utils.RespondWithError(w,http.StatusInternalServerError,"Error while creating user")
            return
        }
        utils.RespondWithJSON(w,http.StatusCreated,databaseUserToUser(user))
    })
}
