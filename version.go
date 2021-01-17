// Copyright © 2020, 2021 Mark Salter
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Package linuxver parses and compares upstream linux kernel version strings.
package linuxver

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	// NoVersion is a special linux version for commits which don't precede
	// a version tag (those following newest version tag when HEAD is not
	// itself tagged with a version).
	NoVersion = &LinuxVersion{255, 0, 0, 0}
	// regex to parse an upstream Linux version string
	vre = regexp.MustCompile(`^v(([0-9]{1,2})[.]([0-9]+)(?:[.]([0-9]+))?(?:-rc([0-9]+))?)$`)
)

// LinuxVersion represents an upstream Linux version.
type LinuxVersion struct {
	// Major holds the major version number. A major value of 255 is special
	// and indicates "no version". This is useful when the linux kernel tree
	// has commits at HEAD which do not precede a version tag (HEAD itself
	// is not tagged with a version).
	Major uint8
	// Minor holds the minor version number.
	Minor uint8
	// Rel holds a release version number (only used in v2.6 kernels).
	Rel   uint8
	// RC holds the release candidate number or zero for the final release.
	RC    uint8
}

// String converts the LinuxVersion to a readable string
func (v *LinuxVersion) String() string {
	if v.Major == 0 {
		return ""
	} else if v.Major == 255 {
		return "Unversioned"
	}
	s := fmt.Sprintf("v%d.%d", v.Major, v.Minor)
	if v.Rel > 0 {
		s = fmt.Sprintf("%s.%d", s, v.Rel)
	}
	if v.RC > 0 {
		s = fmt.Sprintf("%s-rc%d", s, v.RC)
	}
	return s
}

// Equals returns true if LinuxVersion v1 == v2
func (v1 *LinuxVersion) Equals(v2 *LinuxVersion) bool {
	return v1.Major == v2.Major &&
		v1.Minor == v2.Minor &&
		v1.Rel == v2.Rel &&
		v1.RC == v2.RC
}

// ComesBefore returns true if v1 is older than v2
func (v1 *LinuxVersion) ComesBefore(v2 *LinuxVersion) bool {
	if v1.Major < v2.Major {
		return true
	} else if v1.Major > v2.Major {
		return false
	}
	if v1.Minor < v2.Minor {
		return true
	} else if v1.Minor > v2.Minor {
		return false
	}
	if v1.Rel < v2.Rel {
		return true
	} else if v1.Rel > v2.Rel {
		return false
	}
	if v2.RC == 0 {
		return v1.RC != 0
	} else if v1.RC == 0 {
		return false
	}
	return v1.RC < v2.RC
}

// ComesAfter returns true if v1 is newer than v2
func (v1 *LinuxVersion) ComesAfter(v2 *LinuxVersion) bool {
	if v1.Major > v2.Major {
		return true
	} else if v1.Major < v2.Major {
		return false
	}
	if v1.Minor > v2.Minor {
		return true
	} else if v1.Minor < v2.Minor {
		return false
	}
	if v1.Rel > v2.Rel {
		return true
	} else if v1.Rel < v2.Rel {
		return false
	}
	if v1.RC == 0 {
		return v2.RC != 0
	} else if v2.RC == 0 {
		return false
	}
	return v1.RC > v2.RC
}

// New tries to convert a given string into a LinuxVersion.
// If successful, a pointer to the LinuxVersion is returned, else
// nil is returned.
func New(v string) *LinuxVersion {
	var rc, rel uint8

	if v == "Unversioned" {
		return NoVersion
	}

	m := vre.FindStringSubmatch(v)
	if len(m) != 6 {
		return nil
	}

	i, _ := strconv.Atoi(m[2])
	maj := uint8(i)
	i, _ = strconv.Atoi(m[3])
	min := uint8(i)
	if len(m[4]) > 0 {
		i, _ = strconv.Atoi(m[4])
		rel = uint8(i)
	}
	if len(m[5]) > 0 {
		i, _ = strconv.Atoi(m[5])
		rc = uint8(i)
	}
	return &LinuxVersion{maj, min, rel, rc}
}
