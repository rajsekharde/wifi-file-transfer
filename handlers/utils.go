package handlers

import (
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"fmt"
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

func convertBytes(bytes int) string {
	var bytesF float32 = float32(bytes)
	p := 0
	
	for bytesF >= 1024 && p < 4 {
		bytesF /= 1024
		p++
	}

	var unit string
	switch p {
	case 0:
		unit = "B"
	case 1:
		unit = "KB"
	case 2:
		unit = "MB"
	case 3:
		unit = "GB"
	default:
		unit = "TB"
	}

	return fmt.Sprintf("%.2f %s", bytesF, unit)
}