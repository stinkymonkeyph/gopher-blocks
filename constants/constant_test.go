package constants

import (
	"testing"
)

func TestBlockchainNameValue(t *testing.T) {
	if BLOCKCHAIN_NAME != "GopherBlocks" {
		t.Fatalf("Incorrect blockchain name, expected GopherBlocks received %q", BLOCKCHAIN_NAME)
	}
}

func TestHexPrefixValue(t *testing.T) {
	if HEX_PREFIX != "0x" {
		t.Fatalf("Incorrect hex prefix, expected 0x received %q", HEX_PREFIX)
	}
}
