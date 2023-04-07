package mdu

import (
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

func TestToColl(t *testing.T) {
	resetCollection()

	assert.Equal(t, "purchaseOrders", Coll(&PurchaseOrder{}).Name())
	assert.Equal(t, "purchaseOrders", Coll(&purchaseOrder{}).Name())
	assert.Equal(t, "purchaseOrders", Coll(&purchase_Order{}).Name())
}

type PurchaseOrder struct {
	DefaultModel `bson:",inline"`
}

type purchaseOrder struct {
	DefaultModel `bson:",inline"`
}

type purchase_Order struct {
	DefaultModel `bson:",inline"`
}

func shutdown() {
	resetCollection()
	Disconnect()
	ResetDefaultConfig()
}

func setup() {
	err := Init(
		&Config{CtxTimeout: 5 * time.Second},
		"mangoTestDb",
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return
}

func resetCollection() {
	_, err := Coll(&PurchaseOrder{}).DeleteMany(Ctx(), bson.M{})
	_, err = Coll(&purchaseOrder{}).DeleteMany(Ctx(), bson.M{})
	_, err = Coll(&purchase_Order{}).DeleteMany(Ctx(), bson.M{})

	util.PanicErr(err)
}
