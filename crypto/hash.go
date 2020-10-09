// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// HashFileMD5 returns the MD5 hash of a given file
func HashFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	b := hash.Sum(nil)[:16]
	return hex.EncodeToString(b), nil
}
