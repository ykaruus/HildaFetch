package ui

type AsciiService struct {
	file   []byte
	layout *LayoutService
}

func NewAsciiService(file []byte, layout *LayoutService) *AsciiService {

	return &AsciiService{
		file:   file,
		layout: layout,
	}
}

func (ascii *AsciiService) SetAscii() string {
	return ascii.layout.CreateAscii(string(ascii.file))
}
