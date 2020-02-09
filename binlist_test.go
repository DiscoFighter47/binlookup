package binlookup_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DiscoFighter47/binlookup"
)

func TestLookUp(t *testing.T) {
	c := binlookup.NewClient(10. * time.Second)
	t.Run("valid card", func(t *testing.T) {
		resp, err := c.LookUp("45717360")
		assert.NoError(t, err)
		log.Println(resp)
	})

	t.Run("invalid card", func(t *testing.T) {
		_, err := c.LookUp("12345678")
		assert.Error(t, err)
	})
}
