package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func Encrypt(data, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	initializationVector := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, initializationVector); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, initializationVector)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encData string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data, err := base64.URLEncoding.DecodeString(encData)
	if err != nil {
		return nil, err
	}
	initializationVector := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, initializationVector)
	stream.XORKeyStream(data, data)
	return data, nil
}