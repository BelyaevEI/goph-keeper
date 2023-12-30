package textsmodels

import (
	"bytes"
	"encoding/binary"
)

// Structure for work with texts data
type Textsdata struct {
	UserID  uint32 `json:"userid"`
	Text    string `json:"text"`
	Service string `json:"service"`
	Note    string `json:"note"`
}

// MarshalBinary реализует интерфейс encoding.BinaryMarshaler.
func (t *Textsdata) MarshalBinary() ([]byte, error) {
	// Создаем буфер для хранения бинарных данных
	buf := new(bytes.Buffer)

	// Сериализуем UserID
	err := binary.Write(buf, binary.LittleEndian, t.UserID)
	if err != nil {
		return nil, err
	}

	// Сериализуем Text
	err = writeString(buf, t.Text)
	if err != nil {
		return nil, err
	}

	// Сериализуем Service
	err = writeString(buf, t.Service)
	if err != nil {
		return nil, err
	}

	// Сериализуем Note
	err = writeString(buf, t.Note)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalBinary implement interface encoding.BinaryUnmarshaler.
func (t *Textsdata) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)

	err := binary.Read(buf, binary.LittleEndian, &t.UserID)
	if err != nil {
		return err
	}

	text, err := readString(buf)
	if err != nil {
		return err
	}
	t.Text = text

	service, err := readString(buf)
	if err != nil {
		return err
	}
	t.Service = service

	note, err := readString(buf)
	if err != nil {
		return err
	}
	t.Note = note

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
	// Десериализуем длину строки
	var length uint8
	err := binary.Read(buf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// Десериализуем саму строку
	strBytes := make([]byte, length)
	err = binary.Read(buf, binary.LittleEndian, &strBytes)
	if err != nil {
		return "", err
	}

	return string(strBytes), nil
}
