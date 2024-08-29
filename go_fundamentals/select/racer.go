package racer

import (
	"net/http"
	"time"
)

func Racer(firstWebsite, secondWebsite string) (winner string) {
	startA := time.Now()
	http.Get(firstWebsite)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(secondWebsite)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return firstWebsite
	}
	return secondWebsite
}
