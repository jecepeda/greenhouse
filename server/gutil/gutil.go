package gutil

import (
	"math/rand"
	"time"
)

// RandFloat generates a random float
func RandFloat(min, max float64) float64 {
	rand.Seed(time.Now().Unix())
	return min + rand.Float64()*(max-min)
}

// RandBool generates a random boolean
func RandBool() bool {
	rand.Seed(time.Now().Unix())
	return rand.Int63()%2 == 1
}
