package fcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/textproto"
)

type fcmResponse struct {
	Name string `json:"name"`
}

const multipartBoundary = "__END_OF_PART__"

func writeResponsePart(writer *multipart.Writer, data []byte, idx int) error {
	header := make(textproto.MIMEHeader)
	header.Add("Content-Type", "application/http")
	header.Add("Content-Id", fmt.Sprintf("%d", idx+1))
	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}

	_, err = part.Write(data)
	return err
}

func createMultipartResponse(success []fcmResponse, failure []string) ([]byte, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	_ = writer.SetBoundary(multipartBoundary)
	for idx, data := range success {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		var partBuffer bytes.Buffer
		partBuffer.WriteString("HTTP/1.1 200 OK\r\n")
		partBuffer.WriteString("Content-Type: application/json\r\n\r\n")
		partBuffer.Write(b)

		if err := writeResponsePart(writer, partBuffer.Bytes(), idx); err != nil {
			return nil, err
		}
	}

	for idx, data := range failure {
		var partBuffer bytes.Buffer
		partBuffer.WriteString("HTTP/1.1 500 Internal Server Error\r\n")
		partBuffer.WriteString("Content-Type: application/json\r\n\r\n")
		partBuffer.WriteString(data)

		if err := writeResponsePart(writer, partBuffer.Bytes(), idx+len(success)); err != nil {
			return nil, err
		}
	}

	writer.Close()
	return buffer.Bytes(), nil
}
