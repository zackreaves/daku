package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

func Match_input_form() error {
	var player_count string
	player_count_input := huh.NewInput().Title("How Many Players?").Value(&player_count).Validate(func (player_count string) error {
			_,err := strconv.ParseUint(player_count,10,64)
			if err != nil {
				return fmt.Errorf("Requires Natural/Positive Whole Number.")
			}

			return nil
	})

	page_one := huh.NewGroup(player_count_input)

	form := huh.NewForm(page_one)

	return form.Run()
}
