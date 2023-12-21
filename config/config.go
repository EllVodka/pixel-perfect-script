package config

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"training.go/scriptPixelPerfect/models"
)

type (
	Module interface {
		Store() *models.StoreCfg
	}

	C struct {
		parser string
		cfg    *models.StoreCfg
	}
)

var _ Module = (*C)(nil)

// New create new config item.
func New(r io.Reader, parser string) *C {
	c := &C{
		parser: parser,
	}
	
	c.mustReadConfig(r)

	return c
}

// Store return the underlying store db configuration.
func (c *C) Store() *models.StoreCfg {
	return c.cfg
}

// mustReadConfig read config file.
func (c *C) mustReadConfig(reader io.Reader) {
	var err error

	switch c.parser {
	case "json", "yml":
		viper.SetConfigType(c.parser)
	default:
		log.Fatalf("Error, unsupported parsing format for config file %v\n", c.parser)
	}

	if err = viper.ReadConfig(reader); err != nil {
		log.Fatalf("mustReadConfig() - %v", err)
	}

	if err = viper.Unmarshal(&c.cfg); err != nil {
		log.Fatalf("mustReadConfig() - %v", err)
	}

	if err != nil {
		log.Fatalf("mustReadConfig() - error decrypting store encrypted password: %v\n", err)
	}
}

// GetParserFromFileExtension is a helper function to extract parser type from fileName.
func GetParserFromFileExtension(fname string) string {
	extension := filepath.Ext(fname)
	return strings.ReplaceAll(extension, ".", "")
}

// mustOpenConfigFile open config file.
func MustOpenConfigFile(fName string) io.Reader {
	f, err := os.Open(fName)
	if err != nil {
		log.Fatalf("mustOpenConfigFile() - Error, impossible to open config file %v : %v\n", fName, err)
	}

	return f
}
