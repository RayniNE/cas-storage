package peer2peer

import (
	"encoding/gob"
	"io"

	"github.com/RayniNE/cas-storage/models"
)

type Decoder interface {
	Decode(io.Reader, *models.Message) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, message *models.Message) error {
	return gob.NewDecoder(r).Decode(message)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, message *models.Message) error {
	buf := make([]byte, 1024)

	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	message.Payload = buf[:n]

	return nil
}
