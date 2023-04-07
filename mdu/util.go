package mdu

import (
	"github.com/jinzhu/inflection"
	"github.com/softwok/mongo-util/internal/util"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

// Coll returns the collection associated with a model.
func Coll(m Model, opts ...*options.CollectionOptions) *Collection {

	if collGetter, ok := m.(CollectionGetter); ok {
		return collGetter.Collection()
	}

	return CollectionByName(CollName(m), opts...)
}

// CollName returns a model's collection name. The `CollectionNameGetter` will be used
// if the model implements this interface. Otherwise, the collection name is inferred
// based on the model's type using reflection.
func CollName(m Model) string {

	if collNameGetter, ok := m.(CollectionNameGetter); ok {
		return collNameGetter.CollectionName()
	}

	name := reflect.TypeOf(m).Elem().Name()

	return inflection.Plural(util.ToLowerCamelCase(name))
}

// UpsertTrueOption returns new instance of UpdateOptions with the upsert property set to true.
func UpsertTrueOption() *options.UpdateOptions {
	upsert := true
	return &options.UpdateOptions{Upsert: &upsert}
}
