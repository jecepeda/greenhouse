package gutil

import (
	"math/rand"
	"time"
)

func RandFloat(min, max float64) float64 {
	rand.Seed(time.Now().Unix())
	return min + rand.Float64()*(max-min)
}

func RandBool() bool {
	rand.Seed(time.Now().Unix())
	return rand.Int63()%2 == 1
}
