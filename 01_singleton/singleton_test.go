package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	got := GetInstance()
	fmt.Println(got)
}

func TestLazyGetInstance(t *testing.T) {
	got := GetLazyInstance()
	fmt.Println(got)
}
