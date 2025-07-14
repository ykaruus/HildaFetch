package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	ColorPrimary   string   `json:"colorPrimary"`
	ColorSecundary string   `json:"colorSecundary"`
	BorderColor    string   `json:"borderColor"`
	BorderType     string   `json:"borderType"`
	UsernameColor  string   `json:"username_color"`
	Frases         []string `json:"frases"`
}

func OpenConfigJson(file string) (*Config, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}
	var config Config
	errJson := json.Unmarshal(data, &config)

	if errJson != nil {
		return nil, errJson
	}
	return &config, nil
}
