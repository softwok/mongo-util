package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestToLowerCamelCase(t *testing.T) {
	assert.Equal(t, "purchaseOrder", ToLowerCamelCase("PurchaseOrder"))
	assert.Equal(t, "purchaseOrder", ToLowerCamelCase("PurchaseOrder"))
	assert.Equal(t, "purchaseOrder", ToLowerCamelCase("PurchaseOrder"))
}
