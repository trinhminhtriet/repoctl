package core

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed repoctl.1
var CONFIG_MAN []byte

func GenManPages(dir string) error {
	manPath := filepath.Join(dir, "repoctl.1")
	err := os.WriteFile(manPath, CONFIG_MAN, 0644)
	CheckIfError(err)

	fmt.Printf("Created %s\n", manPath)

	return nil
}
