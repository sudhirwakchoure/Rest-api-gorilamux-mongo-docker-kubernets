// Package Model is a model

package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Student  ...
// swagger:model
type Student struct {
	Id              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name,omitempty"   bson:"name,omitempty"`
	City            string             `json:"city,omitempty"  bson:"city,omitempty"`
	Country         string             `json:"country,omitempty" bson:"country,omitempty" `
	Course          string             `json:"course,omitempty" bson:"course,omitempty"`
	YearOfAdmission string             `json:"year_of_admission,omitempty" bson:"yearofadmission,omitempty"`
}
