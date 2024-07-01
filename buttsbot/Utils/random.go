package Utils

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func SelectRandomElement[V any](l []V) V {
	return l[r.Intn(len(l))]
}

func RandomizeChance(n int) bool {
	return r.Intn(n) == 0
}
