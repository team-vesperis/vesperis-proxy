package utils

import "go.minekube.com/common/minecraft/color"

func GetColorOrange() color.Color {
	color, _ := color.Hex("#ff8c00")
	return color
}

func GetColorTitle() color.Color {
	color, _ := color.Hex("#ffb108")
	return color
}

func GetColorUnderTitle() color.Color {
	color, _ := color.Hex("#41b2e3")
	return color
}
