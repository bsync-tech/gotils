package arrary

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestArraryShuffle(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano())

	emptyArray := []interface{}{}
	ArrarysShuffle(emptyArray, source)
	assert.Equal(t, emptyArray, []interface{}{})

	oneElementArray := []interface{}{11}
	ArrarysShuffle(oneElementArray, source)
	assert.Equal(t, oneElementArray, []interface{}{11})

	array := []interface{}{"a", "b", "c"}
	ArrarysShuffle(array, source)
	assert.Contains(t, array, "a")
	assert.Contains(t, array, "b")
	assert.Contains(t, array, "c")
}
