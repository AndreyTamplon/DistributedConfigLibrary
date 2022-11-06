package entity

import (
	"time"
)

type Config struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Data      map[string]string `json:"data"`
	CreatedAt time.Time         `json:"created_at"`
	Version   int64             `json:"version"`
}
