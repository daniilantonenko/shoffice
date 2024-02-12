package app

import "testing"

func Test_IsIPv4(t *testing.T) {
	name := "192.168.0.1"
	want := true

	if got := IsIPv4(name); got != want {
		t.Errorf("hello() = %t, want %t", got, want)
	}
}
