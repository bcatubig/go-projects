package urlshortener

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const urlLength = 8

// Redirect is a shortened URL
type Redirect struct {
	ID           string    `json:"id"`
	ShortUrl     string    `json:"short_url"`
	TargetUrl    string    `json:"target_url"`
	CreationDate time.Time `json:"creation_date"`
}

// NewRedirect creates a new Redirect url
func NewRedirect(url string) *Redirect {
	return &Redirect{
		ID:           uuid.New().String(),
		ShortUrl:     randomString(urlLength),
		TargetUrl:    url,
		CreationDate: time.Now(),
	}
}

// EncodeRedirect serializes the Redirect struck to a json bytes array
func EncodeRedirect(r *Redirect) []byte {
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(r)

	return buf.Bytes()
}

// DecodeRedirect converts a raw Request into a Request
func DecodeRedirect(data []byte) (*Redirect, error) {
	r := &Redirect{}
	err := json.Unmarshal(data, r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// From the lovely Jon Calhoun
// https://www.calhoun.io/creating-random-strings-in-go/
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	b := make([]byte, length)

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
