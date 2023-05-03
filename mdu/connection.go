package mdu

import (
	"context"
	"errors"
	"fmt"
	"github.com/softwok/mongo-util/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var config *Config
var client *mongo.Client
var db *mongo.Database

// Config struct contains extra configuration properties for the mdu package.
type Config struct {
	// Set to 10 second (10*time.Second) for example.
	CtxTimeout time.Duration
}

// NewCtx function creates and returns a new context with the specified timeout.
func NewCtx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}

// Ctx function creates and returns a new context with a default timeout value.
func Ctx() context.Context {
	return ctx()
}

func ctx() context.Context {
	return NewCtx(config.CtxTimeout)
}

// NewClient returns a new mongodb client.
func NewClient(opts ...*options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	ctx := Ctx()
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Println(err)
	}

	return client, nil
}

// NewCollection returns a new collection with the supplied database.
func NewCollection(db *mongo.Database, name string, opts ...*options.CollectionOptions) *Collection {
	coll := db.Collection(name, opts...)

	return &Collection{Collection: coll}
}

// ResetDefaultConfig resets the configuration values, client and database.
func ResetDefaultConfig() {
	config = nil
	client = nil
	db = nil
}

// Init initializes the client and database using the specified configuration values, or default.
func Init(conf *Config, dbName string, opts ...*options.ClientOptions) (err error) {

	// Use the predefined configuration values as default if the user
	// does not provide any.
	if conf == nil {
		conf = defaultConf()
	}
	config = conf
	if client, err = NewClient(opts...); err != nil {
		return err
	}

	db = client.Database(dbName)

	fmt.Println("database connected.")
	return nil
}

func Disconnect() {
	if client == nil {
		return
	}
	err := client.Disconnect(Ctx())
	if err != nil {
		fmt.Println("failed to disconnect from database.")
		return
	}
	fmt.Println("database disconnected.")
}

// CollectionByName returns a new collection using the current configuration values.
func CollectionByName(name string, opts ...*options.CollectionOptions) *Collection {
	return NewCollection(db, name, opts...)
}

// DefaultConfigs returns the current configuration values, client and database.
func DefaultConfigs() (*Config, *mongo.Client, *mongo.Database, error) {
	if util.AnyNil(config, client, db) {
		return nil, nil, nil, errors.New("please setup default config before acquiring it")
	}

	return config, client, db, nil
}

// defaultConf are the default configuration values when none are provided
// to the `Init` method.
func defaultConf() *Config {
	return &Config{CtxTimeout: 10 * time.Second}
}
