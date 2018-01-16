package HighFive

import (
	"fmt"
	"go/types"
	"google.golang.org/api/appengine/v1"
	"time"
)

type Competitor struct {
	PosNum int
	Type   string
}

type Race struct {
	Location string
	Competitors []Competitor
	Close time.Time
}

type Meeting struct {
	Location string
	Race []Race
}

func
