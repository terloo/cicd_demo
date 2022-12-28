package main_test

import "testing"

func Test(t *testing.T) {
	t.Run("demo_test", func(t *testing.T) {
		got := "hello"
		want := "hello"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}