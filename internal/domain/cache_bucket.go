package domain

import "time"

type Bucket struct {
	Tokens     float64
	Rate       float64
	Capacity   float64
	LastRefill time.Time
}

func NewBucket(Tokens, Rate, Capacity float64, LastRefill time.Time) *Bucket {
	return &Bucket{
		Tokens:     Tokens,
		Rate:       Rate,
		Capacity:   Capacity,
		LastRefill: LastRefill,
	}
}
