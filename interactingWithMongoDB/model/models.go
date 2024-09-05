package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

/* primitive.ObjectID is the type used to represent
MongoDB’s _id field, which is a unique identifier for documents.

bson:"_id,omitempty": Similar to json, but for BSON. MongoDB uses BSON format
 for documents, and this tag ensures that the ID field maps to MongoDB's _id field.
*/
