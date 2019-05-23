package test

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestLockCount(t *testing.T) {
	LockCount()
}

func TestSyncMap(t *testing.T) {
	SyncMap()

	dir, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dir)
}
