package register

import (
	"acme/internal/modules/data"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterer_Do(t *testing.T) {
	// inputs
	in := &data.Person{
		FullName: "Chang",
		Phone:    "11122233345",
		Currency: "CNY",
	}

	// call method
	registerer := &Registerer{}
	ID, err := registerer.Do(in)

	// validate expectations
	require.NoError(t, err)
	assert.True(t, ID > 0)
}
