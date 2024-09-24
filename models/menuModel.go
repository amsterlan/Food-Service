package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"price" validate:"required"`
	Start_Date *time.Time         `json:"food_image" validate:"required"`
	End_Date   *time.Time
	Created_at time.Time `json:"created_at"`
	Update_at  time.Time `json:"updated_at"`
	Food_id    string    `json:"food_id"`
	Menu_id    *string   `json:"menu_id" validate:"required"`
}
