package handlers

import (
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

var uploadDir = getUploadDir()

var (
	red = color.RGB(255, 0, 0).SprintFunc()
	green = color.RGB(0, 255, 0).SprintFunc()
	blue = color.RGB(70, 120, 255).SprintFunc()
	grey = color.RGB(200, 200, 200).SprintFunc()
)

func getUploadDir() string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	realExe, err := filepath.EvalSymlinks(exe)
	if err != nil {
		panic(err)
	}

	baseDir := filepath.Dir(realExe)
	return filepath.Join(baseDir, "uploads")
}
