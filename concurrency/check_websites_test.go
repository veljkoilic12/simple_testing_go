package concurrency

import (
	"reflect"
	"testing"
)

const googleString = "http://google.com"
const myblogString = "http://myblog.com"
const faultySiteString = "waat://furhurterwe.geds"

func mockWebsiteChecker(url string) bool {
	if url == faultySiteString {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		googleString,
		myblogString,
		faultySiteString,
	}

	want := map[string]bool{
		googleString:     true,
		myblogString:     true,
		faultySiteString: false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
