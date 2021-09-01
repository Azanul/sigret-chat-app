package websocket

import (
	"bytes"
	"errors"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"io/ioutil"
)

var packetConfig = &packet.Config{
	DefaultCipher: packet.CipherAES256,
}

func Encrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {

	encbuf := bytes.NewBuffer(nil)

	w, _ := armor.Encode(encbuf, "PGP MESSAGE", nil)

	pt, _ := openpgp.SymmetricallyEncrypt(w, key, nil, packetConfig)

	_, err = pt.Write(plaintext)
	if err != nil {
		return
	}

	pt.Close()
	w.Close()
	ciphertext = encbuf.Bytes()

	return
}

func Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	decbuf := bytes.NewBuffer(ciphertext)

	armorBlock, _ := armor.Decode(decbuf)

	failed := false
	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if failed {
			return nil, errors.New("decryption failed")
		}
		failed = true
		return key, nil
	}

	md, err := openpgp.ReadMessage(armorBlock.Body, nil, prompt, packetConfig)
	if err != nil {
		return
	}

	plaintext, err = ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return
	}

	return
}
