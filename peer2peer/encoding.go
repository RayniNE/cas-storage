package peer2peer

import (
	"encoding/gob"
	"io"

	"github.com/RayniNE/cas-storage/models"
)

type Decoder interface {
	Decode(io.Reader, *models.Message) error
}

// GOBDecoder is the decoder that will be used in GOB supported data streams.
type GOBDecoder struct{}

// Decode receives an io.Reader and a *models.Message struct.
// It decodes the io.Reader into the *models.Message struct.
func (dec GOBDecoder) Decode(r io.Reader, message *models.Message) error {
	return gob.NewDecoder(r).Decode(message)
}

// DefaultDecoder is the default decoder that the application will use.
type DefaultDecoder struct{}

// Decode receives an io.Reader and a *models.Message struct. It reads the data
// from the io.Reader and stores it in the *models.Message.Payload property
func (dec DefaultDecoder) Decode(r io.Reader, message *models.Message) error {
	buf := make([]byte, 1024)

	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	message.Payload = buf[:n]

	return nil
}
