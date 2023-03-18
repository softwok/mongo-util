package crud

import (
	mdu "github.com/softwok/mongo-util"
	"github.com/softwok/mongo-util/internal/util"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	testProduct := newProduct("TestCreate", 124)
	productsColl := mdu.Coll(testProduct)
	id, err := productsColl.Create(mdu.Ctx(), testProduct)
	util.PanicErr(err)

	err = productsColl.FindByID(mdu.Ctx(), id, testProduct)
	util.PanicErr(err)

	assert.Equal(t, "TestCreate", testProduct.Name)
	assert.NotNil(t, testProduct.ID)
	assert.NotNil(t, id)
}

func TestFindByIdWithValidId(t *testing.T) {
	testProduct := newProduct("TestFind", 124)
	productsColl := mdu.Coll(testProduct)
	id, err := productsColl.Create(mdu.Ctx(), testProduct)
	util.PanicErr(err)

	err = productsColl.FindByID(mdu.Ctx(), id, testProduct)
	util.PanicErr(err)

	assert.Equal(t, "TestFind", testProduct.Name)
	assert.NotNil(t, testProduct.ID)
	assert.NotNil(t, id)
}

func TestFindByIdWithInvalidId(t *testing.T) {
	assert.NotNil(t, mdu.Coll(&product{}).FindByID(mdu.Ctx(), "invalid id", &product{}))
}

func TestUpdate(t *testing.T) {
	testProduct := newProduct("TestCreate", 124)
	productsColl := mdu.Coll(testProduct)
	id, err := productsColl.Create(mdu.Ctx(), testProduct)
	util.PanicErr(err)
	assert.Equal(t, "TestCreate", testProduct.Name)

	testProduct.Name = "TestUpdate"
	err = productsColl.Update(mdu.Ctx(), testProduct)
	util.PanicErr(err)
	assert.Equal(t, "TestUpdate", testProduct.Name)

	err = productsColl.FindByID(mdu.Ctx(), id, testProduct)
	util.PanicErr(err)

	assert.Equal(t, "TestUpdate", testProduct.Name)
	assert.NotNil(t, testProduct.ID)
	assert.NotNil(t, id)
}

func TestDelete(t *testing.T) {
	testProduct := newProduct("TestDelete", 124)
	productsColl := mdu.Coll(testProduct)
	id, err := productsColl.Create(mdu.Ctx(), testProduct)
	util.PanicErr(err)
	assert.Equal(t, "TestDelete", testProduct.Name)

	err = productsColl.Delete(mdu.Ctx(), testProduct)
	util.PanicErr(err)

	err = productsColl.FindByID(mdu.Ctx(), id, testProduct)
	assert.Equal(t, "mongo: no documents in result", err.Error())
}

func TestFindAll(t *testing.T) {
	resetCollection()
	_ = createProduct("Product1", 100)
	_ = createProduct("Product2", 200)
	_ = createProduct("Product3", 300)
	_ = createProduct("Product4", 400)
	_ = createProduct("Product5", 500)

	var results []product
	err := mdu.Coll(&product{}).FindAll(mdu.Ctx(), &results, bson.D{})
	util.PanicErr(err)
	assert.Equal(t, 5, len(results))
}

// -----------------
// Helpers
//-----------------

func createProduct(name string, price int) interface{} {
	testProduct := newProduct(name, price)
	productsColl := mdu.Coll(testProduct)
	id, err := productsColl.Create(mdu.Ctx(), testProduct)
	util.PanicErr(err)
	return id
}

type product struct {
	mdu.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Price            int    `json:"price" bson:"price"`
}

func newProduct(name string, price int) *product {
	return &product{
		Name:  name,
		Price: price,
	}
}

func shutdown() {
	//resetCollection()
	mdu.Disconnect()
	mdu.ResetDefaultConfig()
}

func setup() {
	err := mdu.Init(
		&mdu.Config{CtxTimeout: 5 * time.Second},
		"mango_test_db",
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return
}

func resetCollection() {
	_, err := mdu.Coll(&product{}).DeleteMany(mdu.Ctx(), bson.M{})

	util.PanicErr(err)
}
