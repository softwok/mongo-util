package crud

import (
	"github.com/google/uuid"
	"github.com/softwok/mongo-util/internal/util"
	"github.com/softwok/mongo-util/mdu"
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

func TestCreateWithID(t *testing.T) {
	resetCollection()

	uuidGenerated := uuid.NewString()
	productsColl := mdu.Coll(&product{})
	testProduct := newProduct("TestCreate", 124)
	testProduct.SetID(uuidGenerated)
	id, err := productsColl.Create(testProduct)
	util.PanicErr(err)

	err = productsColl.FindByID(id, testProduct)
	util.PanicErr(err)

	assert.Equal(t, uuidGenerated, testProduct.ID)
	assert.Equal(t, "TestCreate", testProduct.Name)
	assert.NotNil(t, id)
}

func TestCreate(t *testing.T) {
	resetCollection()

	productsColl := mdu.Coll(&product{})
	testProduct := newProduct("TestCreate", 124)
	id, err := productsColl.Create(testProduct)
	util.PanicErr(err)

	err = productsColl.FindByID(id, testProduct)
	util.PanicErr(err)

	assert.Equal(t, "TestCreate", testProduct.Name)
	assert.NotNil(t, testProduct.ID)
	assert.NotNil(t, id)
}

func TestFindById(t *testing.T) {
	resetCollection()

	productsColl := mdu.Coll(&product{})
	testProduct := insertProduct(newProduct("TestFind", 121))

	err := productsColl.FindByID(testProduct.ID, testProduct)
	util.PanicErr(err)

	assert.Equal(t, "TestFind", testProduct.Name)
	assert.NotNil(t, testProduct.ID)
}

func TestFindByIdWithInvalidId(t *testing.T) {
	productsColl := mdu.Coll(&product{})
	assert.NotNil(t, productsColl.FindByID("invalid id", &product{}))
}

func TestUpdate(t *testing.T) {
	productsColl := mdu.Coll(&product{})
	testProduct := insertProduct(newProduct("TestCreate", 122))
	testProduct.Name = "TestUpdate"
	err := productsColl.Update(testProduct)
	util.PanicErr(err)
	assert.Equal(t, "TestUpdate", testProduct.Name)

	err = productsColl.FindByID(testProduct.ID, testProduct)
	util.PanicErr(err)

	assert.Equal(t, "TestUpdate", testProduct.Name)
	assert.NotNil(t, testProduct.ID)
}

func TestPatch(t *testing.T) {
	productsColl := mdu.Coll(&product{})
	testProduct := insertProduct(newProduct("TestCreate", 122))
	fields := map[string]interface{}{"name": "TestPatch"}
	err := productsColl.Patch(testProduct, fields)
	util.PanicErr(err)

	err = productsColl.FindByID(testProduct.ID, testProduct)
	util.PanicErr(err)

	assert.Equal(t, "TestPatch", testProduct.Name)
	assert.NotNil(t, testProduct.ID)
}

func TestDelete(t *testing.T) {
	productsColl := mdu.Coll(&product{})
	testProduct := insertProduct(newProduct("TestDelete", 124))

	err := productsColl.Delete(testProduct)
	util.PanicErr(err)

	err = productsColl.FindByID(testProduct.ID, testProduct)
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
	err := mdu.Coll(&product{}).FindAll(&results, bson.D{})
	util.PanicErr(err)
	assert.Equal(t, 5, len(results))
}

// -----------------
// Helpers
// -----------------
func insertProduct(testProduct *product) *product {
	productsColl := mdu.Coll(&product{})
	_, err := productsColl.Create(testProduct)
	util.PanicErr(err)
	return testProduct
}

func createProduct(name string, price int) interface{} {
	testProduct := newProduct(name, price)
	productsColl := mdu.Coll(&product{})
	id, err := productsColl.Create(testProduct)
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
	resetCollection()
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
