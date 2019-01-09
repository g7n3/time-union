package time_union

import (
	"math/rand"
	"testing"
	"time"
)

// Make @param:t pieces of fake time sets
func makeTimes(t int) TSlice {
	var set TSlice
	rand.Seed(time.Now().Unix())

	for i := 0; i < t; i++ {
		randStart := rand.Int63n(100)
		randEnd := randStart + rand.Int63n(15)
		set = append(set, T{Start: randStart, End: randEnd})
	}

	return set
}

func Test_Union(t *testing.T) {
	testSet := makeTimes(10)

	t.Logf("Input set: %+v\n", testSet)

	res, err := testSet.Union()
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Output set: %+v\n", res)
}
