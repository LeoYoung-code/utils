package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenUid(t *testing.T) {
	uid := GenUid()
	assert.NotZero(t, uid, "Uid should not be zero")
}
