// Copyright 2021 The life-file Authors
// This file is part of life-file.
//
// life-file is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// life-file is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with life-file. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"testing"

	"github.com/lifefile/life-file/common"
)

func TestFacebook(t *testing.T) {
	for _, tt := range []struct {
		url  string
		want common.Address
	}{
		{
			"https://www.facebook.com/fooz.gazonk/posts/2837228539847129",
			common.HexToAddress("0xDeadDeaDDeaDbEefbEeFbEEfBeeFBeefBeeFbEEF"),
		},
	} {
		_, _, gotAddress, err := authFacebook(tt.url)
		if err != nil {
			t.Fatal(err)
		}
		if gotAddress != tt.want {
			t.Fatalf("address wrong, have %v want %v", gotAddress, tt.want)
		}
	}
}
