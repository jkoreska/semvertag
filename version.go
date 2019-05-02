package main

import (
	"fmt"
)

type Version struct {
	Major uint64
	Minor uint64
	Point uint64
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Point)
}

func (v Version) Next() Version {
	return v.NextPoint()
}

func (v Version) NextPoint() Version {
	return Version{
		Major: v.Major,
		Minor: v.Minor,
		Point: v.Point + 1,
	}
}

func (v Version) NextMinor() Version {
	return Version{
		Major: v.Major,
		Minor: v.Minor + 1,
		Point: 0,
	}
}

func (v Version) NextMajor() Version {
	return Version{
		Major: v.Major + 1,
		Minor: 0,
		Point: 0,
	}
}
