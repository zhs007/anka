package base

import (
	"crypto/aes"
	"encoding/hex"
	"time"

	"github.com/seehuhn/fortuna"
)

// GenerateNameID -
func GenerateNameID() (string, error) {
	gen := fortuna.NewGenerator(aes.NewCipher)
	gen.Seed(time.Now().Unix())
	data := gen.PseudoRandomData(8)
	return hex.EncodeToString(data), nil
}
