package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

func Match_input_form(config Settings) (Match_data, error) {
	var (
		player_count string
		games_options []huh.Option[uint]
		match Match_data
	)
	player_count_input := huh.NewInput().Title("How Many Players?").Value(&player_count).Validate(func (player_count string) error {
			_,err := strconv.ParseUint(player_count,10,64)
			if err != nil {
				return fmt.Errorf("Requires Natural/Positive Whole Number.")
			}

			return nil
	})

	games := Query_games_all(config)
	for i := 0; i < len(games); i++ {
		option := huh.NewOption(games[i].name,games[i].id)
		games_options = append(games_options, option)
	}

	game_select := huh.NewSelect[uint]().Title("Select Game:").Options(games_options...).Value(&match.game_id)

	page_one := huh.NewGroup(player_count_input,game_select)

	form := huh.NewForm(page_one)

	err := form.Run()

	player_count_int,_ := strconv.ParseUint(player_count,10,64)
	match.player_count = uint(player_count_int)

	return match, err
}
