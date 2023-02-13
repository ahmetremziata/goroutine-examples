package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but god %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero_delay", time.Second * 0},
		{"quater_second_delay", time.Millisecond * 250},
		{"half_second_delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		orderFinished = []string{}
		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s, incorrect length of slice; expected 5 but god %d", e.name, len(orderFinished))
		}
	}
}
