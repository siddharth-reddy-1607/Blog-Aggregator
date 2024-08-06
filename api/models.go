package api

import (
	"time"
	"github.com/google/uuid"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
)
type User struct{
    ID uuid.NullUUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Name string `json:"name"`
}

func databaseUserToUser(u database.User) *User{
    return &User{
        ID : u.ID,
        CreatedAt: u.CreatedAt,
        UpdatedAt: u.UpdatedAt,
        Name : u.Name,
    }
}
