package main

import (
	"github.com/charmbracelet/huh"
)

func Match_input_form() error {
	var player_count string
	player_count_input := huh.NewInput().Title("How Many Players?").Value(&player_count)

	page_one := huh.NewGroup(player_count_input)

	form := huh.NewForm(page_one)

	return form.Run()
}
