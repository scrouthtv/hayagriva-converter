package common

import (
	"strconv"
	"time"
)

type DateInformation interface {
	When() time.Time
	Year() int
}

type Date struct {
	T time.Time
}

func (d Date) When() time.Time {
	return d.T
}

func (d Date) Year() int {
	return d.T.Year()
}

func (d Date) String() string {
	return "date: " + d.T.String()
}

type RetrievedDate struct {
	T time.Time
}

func (d RetrievedDate) When() time.Time {
	return d.T
}

func (d RetrievedDate) Year() int {
	return d.T.Year()
}

func (d RetrievedDate) String() string {
	return "retrieved date: " + d.T.String()
}

type WebDate struct {
	T time.Time
}

func (d WebDate) When() time.Time {
	return d.T
}

func (d WebDate) Year() int {
	return d.T.Year()
}

func (d WebDate) String() string {
	return "web date: " + d.T.String()
}

type PublicationYear struct {
	Yr int
}

func (d PublicationYear) When() time.Time {
	return time.Time{}
}

func (d PublicationYear) Year() int {
	return d.Yr
}

func (d PublicationYear) String() string {
	return "publication year: " + strconv.Itoa(d.Yr)
}
