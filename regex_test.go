package gohtparser

import "testing"

func TestIsComment(t *testing.T) {
	if !IsComment("# test") {
		t.Error("コメントが認識されていません")
	}
}
