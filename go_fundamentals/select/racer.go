package racer

import (
	"net/http"
	"time"
)

func Racer(firstWebsite, secondWebsite string) (winner string) {
	a := measureResponseTime(firstWebsite)
	b := measureResponseTime(secondWebsite)

	if a < b {
		return firstWebsite
	}
	return secondWebsite
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)

}
