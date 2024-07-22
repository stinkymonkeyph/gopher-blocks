package blockchain

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	b := NewBlock("0x0", 0)

	if b.PrevHash != "0x0" {
		t.Fatalf("New block return incorrect prev hash, expected 0x0 but received %q", b.PrevHash)
	}
}

func TestToJSON(t *testing.T) {
	b := NewBlock("0x0", 0)
	s := b.ToJson()
	if s == "" {
		t.Fatalf("ToJson returned an empty string")
	}
}
