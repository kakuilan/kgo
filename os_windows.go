// +build windows

package kgo

import (
	"errors"
	"os"
)

// HomeDir 获取当前用户的主目录.
func (ko *LkkOS) HomeDir() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// Prefer standard environment variable USERPROFILE
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home, nil
	}

	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		return "", errors.New("[HomeDir] HOMEDRIVE, HOMEPATH, or USERPROFILE are blank")
	}

	return home, nil
}
