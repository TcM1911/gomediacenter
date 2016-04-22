package gomediacenter

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTicksToMilliSeconds(t *testing.T) {
	assert.Equal(t, int64(2), TicksToMilliSeconds(20000))
}

func TestTicksToNanoSeconds(t *testing.T) {
	assert.Equal(t, int64(300000), TicksToNanoSeconds(30000))
}

func TestNanoSecondsToTicks(t *testing.T) {
	assert.Equal(t, int64(40000), NanoSecondsToTicks(400000))
}

func TestMilliSecondsToTicks(t *testing.T) {
	assert.Equal(t, int64(50000), MilliSecondsToTicks(5))
}