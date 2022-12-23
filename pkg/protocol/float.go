package procotol

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Message struct {
	Version  string
	Resource string
	Headers  Headers
	Body     []byte
}

type Headers map[string]string

func ParseMessage(reader *bufio.Reader) (*Message, error) {
	message := &Message{
		Headers: make(map[string]string),
	}

	startLine, err := reader.ReadString('\n')
	if err != nil {
		panic("did not work at all" + err.Error())
	}

	if startLine == "" {
		panic("Not allowed empty start line")
	}

	startlineArray := strings.Split(startLine, " ")
	message.Version = startlineArray[1]

	fmt.Println(message.Version, startLine)

	for {
		headerLine, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("did not work at all" + err.Error())
		}

		if headerLine == "" {
			fmt.Println("EMPTY!")
			break
		}

		fmt.Println("Geadeler" + headerLine)

		headerLineArray := strings.Split(headerLine, ":")

		if len(headerLineArray) < 2 {
			break
		}

		headerKey := headerLineArray[0]
		headerValue := headerLineArray[1]

		message.Headers[headerKey] = headerValue
	}

	return message, nil
}

func (m *Message) String() string {
	startLine := strings.Join([]string{m.Resource, m.Version}, " ")
	headersLines := m.Headers.String()

	messageString := strings.Join([]string{startLine, headersLines}, "\n")

	return messageString
}

func (h *Headers) String() string {
	headerString := ""

	for key, value := range *h {
		headerString += key + ":" + value + "\n"
	}

	return headerString
}
