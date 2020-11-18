package rc4

import (
	"crypto/cipher"
	"crypto/rc4"
	"crypto/sha256"
	"io"
	"net"
)

type RC4Cipher struct {
	*cipher.StreamReader
	*cipher.StreamWriter
}

type CipherConn struct {
	net.Conn
	Rwc io.ReadWriteCloser
}

func NewRC4Cipher(rwc io.ReadWriteCloser, key []byte) (*RC4Cipher, error) {
	decryptCipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encryptCipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &RC4Cipher{
		StreamReader: &cipher.StreamReader{
			S: decryptCipher,
			R: rwc,
		},
		StreamWriter: &cipher.StreamWriter{
			S: encryptCipher,
			W: rwc,
		},
	}, nil
}

func NewCipherConn(conn net.Conn, password string) (*CipherConn, error) {
	key := sha256.Sum256([]byte(password))
	rwc, err := NewRC4Cipher(conn, key[:])

	if err != nil {
		return nil, err
	}

	return &CipherConn{
		Conn: conn,
		Rwc:  rwc,
	}, nil
}