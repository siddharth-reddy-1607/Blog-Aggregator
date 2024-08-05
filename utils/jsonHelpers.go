package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var(
    DataMarshalError = errors.New("Error while marshalling data")
    DataUnmarshalError = errors.New("Error while unmarshalling JSON")
)

func RespondWithError(w http.ResponseWriter,StatusCode int,ErrMsg string){
    err := struct{ErrMsg string`json:"error"`}{ErrMsg : ErrMsg}
    RespondWithJSON(w,StatusCode,err)

}
func RespondWithJSON(w http.ResponseWriter,statusCode int,payload interface{}){
    data,err := json.Marshal(payload)
    if err != nil{
        log.Printf("%v for payload : %+v",err,payload)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.Header().Add("Content-Type","application/json")
    w.WriteHeader(statusCode)
    if _,err := w.Write(data); err != nil{
        log.Printf("%v",err)
        return 
    }
}
