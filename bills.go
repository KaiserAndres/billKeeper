package main

import (
	"fmt"
	"sort"
	"time"
)

type Bill struct {
	Name       string
	Id         uint
	Expiration time.Time
	Payed      bool
}

func (b Bill) TimeToExpiration() time.Duration {
	// Returns the time until the bill expires
	// Negates the result as to reverse the calculation,
	// will return negative if the bill has expired, possitive if
	// it's not due.
	return -time.Since(b.Expiration)
}

func (b Bill) String() string {
	return fmt.Sprintf("%s - %d: payed=%b", b.Name, b.Id, b.Payed)
}

type ByDate []Bill

func (b ByDate) Len() int { return len(b) }

func (b ByDate) Less(i, j int) bool {
	return b[i].TimeToExpiration() < b[j].TimeToExpiration()
}

func (b ByDate) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func Sort(b []Bill) {
	sort.Sort(ByDate(b))
}
