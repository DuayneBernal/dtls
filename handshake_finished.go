package dtls

import (
	"fmt"
	"reflect"
)

type finished struct {
	data []byte
}

func (h *finished) Init(masterSecret []byte, hash []byte, label string) {
	h.data = generatePrf(masterSecret, []byte(" finished"), hash, label, 12)
}

func (h *finished) Parse(rdr *byteReader) error {
	h.data = rdr.GetBytes(12)

	return nil
}

func (h *finished) Match(masterSecret []byte, hash []byte, label string) bool {
	mac := generatePrf(masterSecret, []byte(" finished"), hash, label, 12)
	if reflect.DeepEqual(mac, h.data) {
		return true
	} else {
		return false
	}
}

func (h *finished) Bytes() []byte {
	w := newByteWriter()
	w.PutBytes(h.data)

	return w.Bytes()
}

func (h *finished) Print() string {
	return fmt.Sprintf("data[%X]", h.data)
}
