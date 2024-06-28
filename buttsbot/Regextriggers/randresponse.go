package regextriggers

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func selectRandomResponse(l []string) string {
	return l[r.Intn(len(l))]
}

func randomizeChance(i int) bool {
	return r.Intn(i) == 0
}
