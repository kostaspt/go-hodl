package ticker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateUrlWithOnlyLimit(t *testing.T) {
	ticker := New(42, "")
	assert.Equal(t, "https://api.coinmarketcap.com/v1/ticker?convert=USD&limit=42", ticker.ApiUrl.String())
}
