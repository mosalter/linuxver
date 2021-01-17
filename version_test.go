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

package linuxver

import (
	"testing"
)

func TestNew(t *testing.T) {
	type vtest struct {
		v   string
		ver *LinuxVersion
	}
	var tests = []vtest{
		{"v1.0", &LinuxVersion{1, 0, 0, 0}},
		{"v1.2", &LinuxVersion{1, 2, 0, 0}},
		{"v1.20", &LinuxVersion{1, 20, 0, 0}},
		{"v10.20", &LinuxVersion{10, 20, 0, 0}},
		{"v1.2.3", &LinuxVersion{1, 2, 3, 0}},
		{"v1.2.3-rc4", &LinuxVersion{1, 2, 3, 4}},
		{"v1.2.13-rc4", &LinuxVersion{1, 2, 13, 4}},
		{"v1.0-rc3", &LinuxVersion{1, 0, 0, 3}},
		{"v1.0-rc3x", nil},
		{"v100.0", nil},
		{"v10.0.3-", nil},
		{"10.0.3", nil},
	}

	for _, test := range tests {
		v := New(test.v)
		if v == nil {
			if test.ver != nil {
				t.Errorf("Got nil! Expected: %s", test.v)
			}
			continue
		}
		if v.Major != test.ver.Major ||
			v.Minor != test.ver.Minor ||
			v.Rel != test.ver.Rel ||
			v.RC != test.ver.RC {
			t.Errorf("Parse of %s failed. Exp: %s  Got %s", test.v, test.ver.String(), v)
		}
		if test.v != test.ver.String() {
			t.Errorf("String() of %s failed:  %s", test.v, test.ver.String())
		}
	}
}

func TestVersionCompare(t *testing.T) {
	type vtest struct {
		v1     string
		v2     string
		op     string
		result bool
	}
	var tests = []vtest{
		// majors
		{"v1.0", "v1.0", "==", true},
		{"v1.0", "v2.0", "==", false},
		{"v1.0", "v2.0", "<", true},
		{"v2.0", "v1.0", "<", false},
		{"v1.0", "v1.0", "<", false},
		{"v2.0", "v1.0", ">", true},
		{"v1.0", "v2.0", ">", false},
		{"v1.0", "v1.0", ">", false},
		// minors
		{"v1.1", "v1.1", "==", true},
		{"v1.0", "v1.1", "==", false},
		{"v1.0", "v1.1", "<", true},
		{"v1.1", "v1.0", "<", false},
		{"v1.0", "v1.0", "<", false},
		{"v1.1", "v1.0", ">", true},
		{"v1.0", "v1.1", ">", false},
		{"v1.0", "v1.0", ">", false},
		// rcs
		{"v1.0-rc1", "v1.0-rc1", "==", true},
		{"v1.0-rc1", "v1.0-rc2", "==", false},
		{"v1.0", "v1.0-rc1", "==", false},
		{"v1.0-rc1", "v1.0-rc2", "<", true},
		{"v1.0-rc1", "v1.0-rc1", "<", false},
		{"v1.0-rc1", "v1.0", "<", true},
		{"v1.0", "v1.0-rc1", "<", false},
		{"v1.0-rc1", "v1.0-rc2", ">", false},
		{"v1.0-rc1", "v1.0-rc1", ">", false},
		{"v1.0-rc1", "v1.0", ">", false},
		{"v1.0", "v1.0-rc1", ">", true},
	}

	for _, test := range tests {
		v1 := New(test.v1)
		v2 := New(test.v2)

		if test.op == "==" {
			if v1.Equals(v2) != test.result {
				t.Errorf("%s == %s: failed!", test.v1, test.v2)
			}
		} else if test.op == "<" {
			if v1.ComesBefore(v2) != test.result {
				t.Errorf("%s < %s: failed!", test.v1, test.v2)
			}
		} else if test.op == ">" {
			if v1.ComesAfter(v2) != test.result {
				t.Errorf("%s > %s: failed!", test.v1, test.v2)
			}
		} else {
			t.Fatalf("Unknown op: \"%s\"", test.op)
		}

	}
}
