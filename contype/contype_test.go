package contype

import (
	"testing"
)

func Test_url_1(t *testing.T) {
	if FileType("build.golang.org") != "text/html" {
		t.Error("Error with Test #1")
	} else {
		t.Log("Test #1 passed!")
	}
}

func Test_url_2(t *testing.T) {
	if FileType(".mkv") != "video/x-matroska" {
		t.Error("Error with Test #2")
	} else {
		t.Log("Test #2 passed!")
	}
}

func Test_url_3(t *testing.T) {
	if FileType("https://github.com/kampsy/Nodejs/contype.js") != "application/javascript" {
		t.Error("Error with Test #3")
	} else {
		t.Log("Test #3 passed!")
	}
}
