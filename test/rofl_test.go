package test

import (
	"math/rand"
	"testing"
	"time"
)

func TestRofl(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	if r.Int63()&1 == 1 {
		t.FailNow()
	}
}
