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
type book struct {
	mdu.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}
```

## Create

```go
testBook := &book{
    Name:  "Test Create",
    Pages: 1,
}
booksColl := mdu.Coll(testBook)
id, err := booksColl.Create(mdu.Ctx(), testBook)
```

## Update

```go
testBook.Name = "Test Update"
booksColl := mdu.Coll(testBook)
err := booksColl.Update(mdu.Ctx(), testBook)
```

## Select

```go
booksColl := mdu.Coll(testBook)
err := booksColl.FindByID(mdu.Ctx(), id, testBook)
```

## Delete

```go
booksColl := mdu.Coll(testBook)
err = booksColl.Delete(mdu.Ctx(), testBook)
```


## APIs
- `FindByID`: FindByID method finds a doc and decodes it to a model, otherwise returns an error.
- `First`: First method searches and returns the first document in the search results.
- `Create`: Create method inserts a new model into the database.
- `Update`: Update function persists the changes made to a model to the database using the specified context.
- `Delete`: Delete method deletes a model (doc) from a collection using the specified context.
- `FindAll`: 