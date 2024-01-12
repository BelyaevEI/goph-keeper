package binarymodels

import (
	"bytes"
	"encoding/binary"
)

// Structure for work with bin data
type Binarydata struct {
	UserID  uint32 `json:"userid"`
	Bin     string `json:"bin"`
	Service string `json:"service"`
	Note    string `json:"note"`
}

// MarshalBinary implement interface encoding.BinaryMarshaler.
func (b *Binarydata) MarshalBinary() ([]byte, error) {

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, b.UserID)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, b.Bin)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, b.Service)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, b.Note)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalBinary implements interfce encoding.BinaryUnmarshaler.
func (b *Binarydata) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)

	err := binary.Read(buf, binary.LittleEndian, &b.UserID)
	if err != nil {
		return err
	}

	bin, err := readString(buf)
	if err != nil {
		return err
	}
	b.Bin = bin

	service, err := readString(buf)
	if err != nil {
		return err
	}
	b.Service = service

	note, err := readString(buf)
	if err != nil {
		return err
	}
	b.Note = note

	return nil
}

// writeString searializing data to binary
func writeString(buf *bytes.Buffer, str string) error {

	err := binary.Write(buf, binary.LittleEndian, uint8(len(str)))
	if err != nil {
		return err
	}

	err = binary.Write(buf, binary.LittleEndian, []byte(str))
	if err != nil {
		return err
	}

	return nil
}

// readString desiarilizing data from binary
func readString(buf *bytes.Reader) (string, error) {

	var length uint8
	err := binary.Read(buf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length)
	err = binary.Read(buf, binary.LittleEndian, &strBytes)
	if err != nil {
		return "", err
	}

	return string(strBytes), nil
}
