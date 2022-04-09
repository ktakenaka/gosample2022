package environment

import (
	"embed"
	"fmt"
	"path/filepath"
)

//go:embed *
var configs embed.FS

func GetConfig(env string) ([]byte, error) {
	fileName := filepath.Join(".", fmt.Sprintf("%s.yml", env))
	return configs.ReadFile(fileName)
}
