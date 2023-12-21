package raja

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func GetPassword() ([]byte, error) {
	jsonXrs := new(xrs)
	resp, err := Client.Get(BASE_URL + "/assets/File/xrs.json")
	if err != nil {
		return nil, fmt.Errorf("get xrs error: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(body, jsonXrs)
	if err != nil {
		return nil, fmt.Errorf("xrs json decode error: %w", err)
	}
	return []byte(jsonXrs.Key), nil
}

func encrypt(input string, password []byte) (string, error) {
	randSalt, err := randomBytes(16)
	if err != nil {
		return "", err
	}
	randIV, err := randomBytes(16)
	if err != nil {
		return "", err
	}

	key := pbkdf2.Key(password, randSalt, 100, 32, sha1.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, randIV)
	inputPad := padPKCS7([]byte(input), mode.BlockSize())
	encrypted := make([]byte, len(inputPad))
	mode.CryptBlocks(encrypted, inputPad)

	result := appendByteSlices([][]byte{randSalt, randIV, encrypted})
	b64Encoded := base64.StdEncoding.EncodeToString(result)
	return b64Encoded, nil
}

func decrypt(input string, password []byte) (string, error) {
	b64Decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	randSalt := b64Decoded[:16]
	randIV := b64Decoded[16:32]
	encrypted := b64Decoded[32:]
	result := make([]byte, len(encrypted))

	key := pbkdf2.Key(password, randSalt, 100, 32, sha1.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, randIV)
	mode.CryptBlocks(result, encrypted)
	result = unPadPKCS7(result)
	return string(result), nil
}

func appendByteSlices(sl [][]byte) []byte {
	tmp := make([]byte, 0)
	for _, i := range sl {
		tmp = append(tmp, i...)
	}
	return tmp
}

func randomBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	return buf, err
}

func padPKCS7(input []byte, blockSize int) []byte {
	size := blockSize - (len(input) % blockSize)
	char := []byte{byte(size)}
	result := append(input, bytes.Repeat(char, size)...)
	return result
}

func unPadPKCS7(input []byte) []byte {
	char := input[len(input)-1]
	size := int(char)
	result := input[:len(input)-size]
	return result
}
