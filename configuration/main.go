package configuration

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	HTTP struct {
		Port string `json:"port"`
	} `json:"http"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Sslmode  string `json:"sslmode"`
		Timezone string `json:"timezone"`
		Timeout  int    `json:"timeout"`
	} `json:"database"`
	Oauth struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RedirectUrl  string `json:"redirect_url"`
	}
	Authentication struct {
		Secret        string `json:"secret"`
		RefreshSecret string `json:"refresh_secret"`
	}
	Application struct {
		Url string `json:"url"`
	}
}

func Load(path string) (*Configuration, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configuration Configuration

	err = json.Unmarshal(bytes, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
