package api

import (
	"errors"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
)

type APIConfig struct{
    DBQueries *database.Queries
}

var (
    jsonDecodeError = errors.New("Error while decoding JSON")
)
