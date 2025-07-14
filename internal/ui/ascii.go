package ui

import "os"

type AsciiService struct {
	file   []byte
	layout *LayoutService
}

func NewAsciiService(filename string, layout *LayoutService) (*AsciiService, error) {
	file, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &AsciiService{
		file:   file,
		layout: layout,
	}, nil
}

func (ascii *AsciiService) SetAscii() string {
	return ascii.layout.CreateAscii(string(ascii.file))
}
