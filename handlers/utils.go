package handlers

import (
	"github.com/fatih/color"
)

const uploadDir = "./uploads"

var (
	red = color.RGB(255, 0, 0).SprintFunc()
	green = color.RGB(0, 255, 0).SprintFunc()
	blue = color.RGB(200, 200, 255).SprintFunc()
	grey = color.RGB(200, 200, 200).SprintFunc()
)