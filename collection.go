package mdu

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/softwok/mongo-util/builder"
	"github.com/softwok/mongo-util/field"
)

// Collection performs operations on models and the given Mongodb collection
type Collection struct {
	*mongo.Collection
}

// FindByID method finds a doc and decodes it to a model, otherwise returns an error.
// The id field can be any value that if passed to the `PrepareID` method, it returns
// a valid ID (e.g.string, bson.ObjectId).
func (c *Collection) FindByID(ctx context.Context, id interface{}, model Model, opts ...*options.FindOneOptions) error {
	return first(ctx, c, bson.M{field.ID: id}, model, opts...)
}

// First method searches and returns the first document in the search results.
func (c *Collection) First(ctx context.Context, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return first(ctx, c, filter, model, opts...)
}

// Create method inserts a new model into the database.
func (c *Collection) Create(ctx context.Context, model Model, opts ...*options.InsertOneOptions) (interface{}, error) {
	id, err := model.PrepareID(model.GetID())

	if err != nil {
		return nil, err
	}
	model.SetID(id)
	return create(ctx, c, model, opts...)
}

// Update function persists the changes made to a model to the database using the specified context.
// Calling this method also invokes the model's mgm updating, updated,
// saving, and saved hooks.
func (c *Collection) Update(ctx context.Context, model Model, opts ...*options.UpdateOptions) error {
	return update(ctx, c, model, opts...)
}

// Delete method deletes a model (doc) from a collection using the specified context.
// To perform additional operations when deleting a model
// you should use hooks rather than overriding this method.
func (c *Collection) Delete(ctx context.Context, model Model) error {
	return deleteByID(ctx, c, model)
}

// FindAll finds, decodes and returns the results using the specified context.
func (c *Collection) FindAll(ctx context.Context, results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cur, err := c.Find(ctx, filter, opts...)

	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}

//--------------------------------
// Aggregation methods
//--------------------------------

// SimpleAggregateFirst performs a simple aggregation, decodes the first aggregate result and returns it using the provided result parameter.
// The value of `stages` can be Operator|bson.M
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (c *Collection) SimpleAggregateFirst(ctx context.Context, result interface{}, stages ...interface{}) (bool, error) {
	cur, err := c.SimpleAggregateCursor(ctx, stages...)
	if err != nil {
		return false, err
	}
	if cur.Next(ctx) {
		return true, cur.Decode(result)
	}
	return false, nil
}

// SimpleAggregate performs a simple aggregation, decodes the aggregate result and returns the list using the provided result parameter.
// The value of `stages` can be Operator|bson.M
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (c *Collection) SimpleAggregate(ctx context.Context, results interface{}, stages ...interface{}) error {
	cur, err := c.SimpleAggregateCursor(ctx, stages...)
	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}

// SimpleAggregateCursor performs a simple aggregation and returns a cursor over the resulting documents.
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (c *Collection) SimpleAggregateCursor(ctx context.Context, stages ...interface{}) (*mongo.Cursor, error) {
	pipeline := bson.A{}

	for _, stage := range stages {
		if operator, ok := stage.(builder.Operator); ok {
			pipeline = append(pipeline, builder.S(operator))
		} else {
			pipeline = append(pipeline, stage)
		}
	}

	return c.Aggregate(ctx, pipeline, nil)
}
