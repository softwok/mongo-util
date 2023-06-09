package mdu

import (
	"context"
)

// CollectionGetter interface contains a method to return
// a model's custom collection.
type CollectionGetter interface {
	// Collection method return collection
	Collection() *Collection
}

// CollectionNameGetter interface contains a method to return
// the collection name of a model.
type CollectionNameGetter interface {
	// CollectionName method return model collection's name.
	CollectionName() string
}

// Model interface contains base methods that must be implemented by
// each model. If you're using the `DefaultModel` struct in your model,
// you don't need to implement any of these methods.
type Model interface {
	// PrepareID converts the id value if needed, then
	// returns it (e.g.convert string to objectId).
	PrepareID(id string) (string, error)

	GetID() string
	SetID(id string)
}

// DefaultModel struct contains a model's default fields.
type DefaultModel struct {
	IDField    `bson:",inline"`
	DateFields `bson:",inline"`
}

// DefaultTenantModel struct contains a model's default fields. This is useful for multi tenant systems.
type DefaultTenantModel struct {
	IDField       `bson:",inline"`
	DateFields    `bson:",inline"`
	TenantIdField `bson:",inline"`
}

// Creating function calls the inner fields' defined hooks
func (model *DefaultModel) Creating(ctx context.Context) error {
	return model.DateFields.Creating(ctx)
}

// Saving function calls the inner fields' defined hooks
func (model *DefaultModel) Saving(ctx context.Context) error {
	return model.DateFields.Saving(ctx)
}
