package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.False(t, IsEmpty("text"))
	assert.False(t, IsEmpty("	"))
}

func TestIsNotEmpty(t *testing.T) {
	assert.False(t, IsNotEmpty(""))
	assert.True(t, IsNotEmpty("text"))
	assert.True(t, IsNotEmpty("	"))
}

func TestIsBlank(t *testing.T) {
	assert.True(t, IsBlank(""))
	assert.True(t, IsBlank("	"))
	assert.False(t, IsBlank("text"))
}

func TestIsNotBlank(t *testing.T) {
	assert.False(t, IsNotBlank(""))
	assert.False(t, IsNotBlank("	"))
	assert.True(t, IsNotBlank("text"))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse(""), "")
	assert.Equal(t, Reverse("X"), "X")
	assert.Equal(t, Reverse("b\u0301"), "b\u0301")
	assert.Equal(t, Reverse("😎⚽"), "⚽😎")
	assert.Equal(t, Reverse("Les Mise\u0301rables"), "selbare\u0301siM seL")
	assert.Equal(t, Reverse("ab\u0301cde"), "edcb\u0301a")
	assert.Equal(t, Reverse("This `\xc5` is an invalid UTF8 character"), "retcarahc 8FTU dilavni na si `�` sihT")
	assert.Equal(t, Reverse("The quick bròwn 狐 jumped over the lazy 犬"), "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT")
	assert.Equal(t, Reverse("رائد شوملي"), "يلموش دئار")
}
