package mongo

import (
	"github.com/Sirupsen/logrus"
	"github.com/rs/rest-layer/schema"
	"golang.org/x/net/context"
)

// File is the file struct
type File struct {
	filename string
}

var (
	// NewFile is a field hook handler that generates a new globally unique id if none exist,
	// to be used in schema with OnInit.
	//
	// The generated ID is a Mongo like base64 object id (mgo/bson code has been embedded
	// into this function to prevent dep)
	NewFile = func(ctx context.Context, value interface{}) interface{} {
		if value == nil {
			// value = newID()
		}
		return value
	}

	// MediaField is a common schema field configuration that generate an globally unique id
	// for new item id.
	MediaField = schema.Field{
		Description: "A media file stored in GridFS",
		// OnInit:      NewID,
		Filterable: true,
		Sortable:   true,
		Validator:  &schema.String{
		// This regexp matches a base32 id
		// Regexp: "[0-9a-z]{20}$",
		},
	}
)

// newID returns a new globally unique id using a copy of the mgo/bson algorithm.
// func newID() string {
// 	return xid.New().String()
// }

// Validate validates...
func (f File) Validate(value interface{}) (interface{}, error) {
	logrus.Info(value)
	return nil, nil
}
