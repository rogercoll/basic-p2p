
package server

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
	assert.Empty(t, err)
}