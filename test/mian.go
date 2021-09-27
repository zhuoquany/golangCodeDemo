package main

import (
	"encoding/hex"
	"fmt"

	"github.com/multiformats/go-multihash"
)

func main() {
	// ignores errors for simplicity.
	// don't do that at home.
	// Decode a SHA1 hash to a binary buffer
	buf, _ := hex.DecodeString("22a794a209ce515066ef6f193fcd7a17295a7005896e2cfbc7717655a24d5ada")

	// Create a new multihash with it.
	mHashBuf, _ := multihash.EncodeName(buf, "sha2-256")
	// Print the multihash as hex string
	fmt.Printf("hex: %s\n", hex.EncodeToString(mHashBuf))

	// Parse the binary multihash to a DecodedMultihash
	mHash, _ := multihash.Decode(mHashBuf)
	// Convert the sha1 value to hex string
	sha1hex := hex.EncodeToString(mHash.Digest)
	// Print all the information in the multihash
	fmt.Printf("obj: %v 0x%x %d %s\n", mHash.Name, mHash.Code, mHash.Length, sha1hex)
}
