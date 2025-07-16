package utils

type Config struct {
	ColorPrimary   string   `json:"colorPrimary"`
	ColorSecundary string   `json:"colorSecundary"`
	BorderColor    string   `json:"borderColor"`
	BorderType     string   `json:"borderType"`
	UsernameColor  string   `json:"username_color"`
	Frases         []string `json:"frases"`
}

func SetConfigDefaults(config *Config) {

	config.ColorPrimary = "#80B9B9"
	config.ColorSecundary = "#D46A6A"
	config.UsernameColor = "#80B9B9"
	config.BorderColor = "80B9B9"
	config.BorderType = "rounded"
	config.Frases = []string{
		"Well, that was traumatic... but that’s the life of an adventurer.",
		"The wilderness is calling.",
		"Woffs are just misunderstood.",
	}

}

func SetTemplateDefaults() string {

	return "[+] *OS*: {{OS}}\n[+] *PROCESSADOR* : {{CPU}}\n[+] *Memoria* : {{Memory usage}} / {{Memory}} GiB\n[+] *Armazenamento* : {{Disk Used}} / {{Disk}} GiB\n[+] *Tempo de boot* : {{Uptime}}"
}
