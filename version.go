package main

import (
	"fmt"
)

type Version struct {
	Major uint64
	Minor uint64
	Patch uint64
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v Version) Next() Version {
	return v.NextPatch()
}

func (v Version) NextPatch() Version {
	return Version{
		Major: v.Major,
		Minor: v.Minor,
		Patch: v.Patch + 1,
	}
}

func (v Version) NextMinor() Version {
	return Version{
		Major: v.Major,
		Minor: v.Minor + 1,
		Patch: 0,
	}
}

func (v Version) NextMajor() Version {
	return Version{
		Major: v.Major + 1,
		Minor: 0,
		Patch: 0,
	}
}
