package configs

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"sync"
)

type Config struct {
	NationalAccountNumber    string `json:"nationalAccountNumber" xml:"nationalAccountNumber"`
	LiquidationAccountNumber string `json:"liquidatitonAccountNumber" xml:"liquidatitonAccountNumber"`
}

var (
	config Config
	once   sync.Once
)

func (c *Config) NewConfig() Config {
	return Config{
		NationalAccountNumber:    "BY04CBDC36029110100040000000",
		LiquidationAccountNumber: "BY04CBDC36029110100040000001",
	}
}

func NewConfig() *Config {
	once.Do(func() {
		config = config.NewConfig()
		thisLocation, err := os.Getwd()
		fmt.Println(thisLocation)
		file, err := os.ReadFile(thisLocation + "/" + "settings.xml")
		if err != nil {
			log.Println("Error read file")

		} else {
			fmt.Println(string(file))
			err = xml.Unmarshal(file, &config)
			if err != nil {
				log.Println("Error unmarshal file")
			}
		}

	})

	return &config
}
