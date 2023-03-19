## Introduction
Wrapper and Util for MongoDB driver to perform CRUD operations.

https://www.mongodb.com/docs/manual/reference/operator/


## Initialization
```go
err := mdu.Init(
		&mdu.Config{CtxTimeout: 5 * time.Second},
		"mango_test_db",
		options.Client().ApplyURI("mongodb://localhost:27017"))
if err != nil {
    panic(err)
}
```

## MongoDB Collection Definition
By adding `mdu.DefaultModel` in model will include following attributes are added and generated automatically:
- Mongo Object ID: 
  - ```ID primitive.ObjectID `json:"id" bson:"_id,omitempty"```
- Created Date: 
  - ```CreatedAt time.Time `json:"created_at" bson:"created_at"```
- Updated Date: 
  - ```UpdatedAt time.Time `json:"updated_at" bson:"updated_at"```

Example Model:
```go
type product struct {
	mdu.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Price            int    `json:"price" bson:"price"`
}
```

## Create

```go
testProduct := &product{
    Name:  "Test Create",
    Price: 100,
}
productsColl := mdu.Coll(testProduct)
id, err := productsColl.Create(mdu.Ctx(), testProduct)
```

## Update

```go
testProduct.Name = "Test Update"
productsColl := mdu.Coll(testProduct)
err := productsColl.Update(mdu.Ctx(), testProduct)
```

## Select

```go
productsColl := mdu.Coll(&product{})
err := productsColl.FindByID(mdu.Ctx(), id, testProduct)
```

## Delete

```go
productsColl := mdu.Coll(&product{})
err = productsColl.Delete(mdu.Ctx(), testProduct)
```


## APIs
- `FindByID`: FindByID method finds a doc and decodes it to a model, otherwise returns an error.
- `First`: First method searches and returns the first document in the search results.
- `Create`: Create method inserts a new model into the database.
- `Update`: Update function persists the changes made to a model to the database using the specified context.
- `Delete`: Delete method deletes a model (doc) from a collection using the specified context.
- `FindAll`: FindAll finds, decodes and returns the results using the specified context.