package builder

import (
	"github.com/softwok/mongo-util/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

// appendNotNull appends the provided key and value to the map if the value is not nil.
func appendNotNull(m bson.M, key string, val interface{}) {
	if !util.IsNil(val) {
		m[key] = val
	}
}
