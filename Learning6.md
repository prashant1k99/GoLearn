### 61) SHA256 Hashes:
```go
package main

// Go implements several hash functions in various crypto/* packages.
import (
	"crypto/sha256"
	"fmt"
)

// SHA256 hashes are frequently used to compute short identities for binary or text blobs. For example, TLS/SSL certificates use SHA256 to compute a certificate's signature. Here's how to compute SHA256 hashes in Go.

func main() {
	s := "sha256 this is sgtring"

	// Here we start with a new hash
	h := sha256.New()

	// Write expects bytes. If you hace a string s, use []byte(s) to coerce it to bytes
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice. THe argument to Sum can be used to append to an exisitng byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	fmt.Println(s)
	// sha256 this is sgtring
	fmt.Printf("%x\n", bs)
	// 3a97e70165a808c6d867ecb3d250de8712822c619aeedd6dd7f7794117b37a16
}
```
