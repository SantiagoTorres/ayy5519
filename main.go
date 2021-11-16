package main 

import (
    "crypto/ed25519"
    "crypto"
    "crypto/rand"
    "fmt"
)

func main() {
    var seed []byte

    fmt.Sscanf("23839fa830b636b6d1c0e1f6fffebd94c8827af761188ec1a5664c53e395ef6e", "%x", &seed)
    privkey := ed25519.NewKeyFromSeed(seed)
    signature, _ := privkey.Sign(rand.Reader, []uint8("ayylmao"), crypto.Hash(0))
    fmt.Printf("%x\n", signature)
}
