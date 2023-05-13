package algo

import "math/rand"

var rnd *rand.Rand

func SetSeed(seed int64) {
	rnd = rand.New(rand.NewSource(seed))
}
