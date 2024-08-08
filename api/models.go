package api

import (
	"time"
	"github.com/google/uuid"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
)
type User struct{
    ID uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Name string `json:"name"`
    APIKey string `json:"api_key"`
}

func databaseUserToUser(u *database.User) *User{
    return &User{
        ID : u.ID,
        CreatedAt: u.CreatedAt,
        UpdatedAt: u.UpdatedAt,
        Name : u.Name,
        APIKey: u.Apikey,
    }
}

type Feed struct{
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
    Url string `json:"Url"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    UserID uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(f *database.Feed) *Feed{
    return &Feed{
        ID: f.ID,
        Name : f.Name,
        Url : f.Url,
        CreatedAt : f.CreatedAt,
        UpdatedAt: f.UpdatedAt,
        UserID: f.UserID,
    }
}
