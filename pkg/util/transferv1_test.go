package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransferV1Builder(t *testing.T) {
	tr := NewTransferV1Builder().MustBuild()
	assert.Equal(t, "9ar4tAzhzw3gt6NPAG4hjb1Y3BeV85DAqfaeQPHmuiNG", tr.ID.String())
}
