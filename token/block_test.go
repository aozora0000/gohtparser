package token

import "testing"

func TestBlock_GetName(t *testing.T) {
	block := NewBlock("TestBlock", []string{})
	if block.GetName() != "TestBlock" {
		t.Errorf("%s != %s", "TestBlock", block.GetName())
	}
	t.Log("TestBlock_GetName OK")
}

func TestBlock_ToString(t *testing.T) {
	block := NewBlock("files", []string{})

}
