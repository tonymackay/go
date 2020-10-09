// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package crypto

import "testing"

func TestHashFileMD5(t *testing.T) {
	want := "6f5902ac237024bdd0c176cb93063dc4"
	got, err := HashFileMD5("test_file.txt")
	if err != nil || want != got {
		t.Fatalf("want: %v got: %v, err: %v", want, got, err)
	}
}
