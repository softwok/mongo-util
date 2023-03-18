package crud

import (
	mdu "github.com/softwok/mongo-util"
	"github.com/softwok/mongo-util/internal/util"
	"github.com/stretchr/testify/assert"
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
	testBook := newBook("TestCreate", 124)
	booksColl := mdu.Coll(testBook)
	id, err := booksColl.Create(mdu.Ctx(), testBook)
	util.PanicErr(err)

	err = booksColl.FindByID(mdu.Ctx(), id, testBook)
	util.PanicErr(err)

	assert.Equal(t, "TestCreate", testBook.Name)
	assert.NotNil(t, testBook.ID)
	assert.NotNil(t, id)
}

func TestFindByIdWithValidId(t *testing.T) {
	testBook := newBook("TestFind", 124)
	booksColl := mdu.Coll(testBook)
	id, err := booksColl.Create(mdu.Ctx(), testBook)
	util.PanicErr(err)

	err = booksColl.FindByID(mdu.Ctx(), id, testBook)
	util.PanicErr(err)

	assert.Equal(t, "TestFind", testBook.Name)
	assert.NotNil(t, testBook.ID)
	assert.NotNil(t, id)
}

func TestFindByIdWithInvalidId(t *testing.T) {
	assert.NotNil(t, mdu.Coll(&book{}).FindByID(mdu.Ctx(), "invalid id", &book{}))
}

func TestUpdate(t *testing.T) {
	testBook := newBook("TestCreate", 124)
	booksColl := mdu.Coll(testBook)
	id, err := booksColl.Create(mdu.Ctx(), testBook)
	util.PanicErr(err)
	assert.Equal(t, "TestCreate", testBook.Name)

	testBook.Name = "TestUpdate"
	err = booksColl.Update(mdu.Ctx(), testBook)
	util.PanicErr(err)
	assert.Equal(t, "TestUpdate", testBook.Name)

	err = booksColl.FindByID(mdu.Ctx(), id, testBook)
	util.PanicErr(err)

	assert.Equal(t, "TestUpdate", testBook.Name)
	assert.NotNil(t, testBook.ID)
	assert.NotNil(t, id)
}

func TestDelete(t *testing.T) {
	testBook := newBook("TestDelete", 124)
	booksColl := mdu.Coll(testBook)
	id, err := booksColl.Create(mdu.Ctx(), testBook)
	util.PanicErr(err)
	assert.Equal(t, "TestDelete", testBook.Name)

	err = booksColl.Delete(mdu.Ctx(), testBook)
	util.PanicErr(err)

	err = booksColl.FindByID(mdu.Ctx(), id, testBook)
	assert.Equal(t, "mongo: no documents in result", err.Error())
}

// -----------------
// Helpers
//-----------------

type book struct {
	mdu.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func newBook(name string, pages int) *book {
	return &book{
		Name:  name,
		Pages: pages,
	}
}

func shutdown() {
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
