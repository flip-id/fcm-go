package fcm

import "testing"

func fatalIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
