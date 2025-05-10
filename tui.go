package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

func Match_input_form(config Settings) (Match_data) {
	var (
		player_count string
		games_options []huh.Option[uint]
		players_options []huh.Option[uint]
		match Match_data
		player_id []uint
	)
	player_count_input := huh.NewInput().Title("How Many Players?").Value(&player_count).Validate(func (player_count string) error {
			_,err := strconv.ParseUint(player_count,10,64)
			if err != nil {
				return fmt.Errorf("Requires Natural/Positive Whole Number.")
			}

			return nil
	})

	player_count_input.Run()

	games := Query_games_all(config)
	for i := 0; i < len(games); i++ {
		option := huh.NewOption(games[i].name,games[i].id)
		games_options = append(games_options, option)
	}

	game_select := huh.NewSelect[uint]().Title("Select Game:").Options(games_options...).Value(&match.game_id)

	// page1 := huh.NewGroup(player_count_input,game_select)
	// match_form := huh.NewForm(page1)

	// err := match_form.Run()
	err := game_select.Run()
	Error_check(err)

	player_count_int,_ := strconv.ParseUint(player_count,10,64)
	match.player_count = uint(player_count_int)

	players := Query_players_all(config)
	for i := 0; i < len(players); i++ {
		option := huh.NewOption(players[i].name_first,players[i].id)
		players_options = append(players_options, option)
	}

	player_select := huh.NewMultiSelect[uint]().Title("Select Players:").Options(players_options...).Limit(int(player_count_int)).Value(&player_id)
	page2 := huh.NewGroup(player_select)
	player_form := huh.NewForm(page2)

	err = player_form.Run()
	Error_check(err)

	player_data := make([]Player_data,len(player_id))

	for i := 0; i < len(player_data); i++ {
		player_data[i].player_id = player_id[i]
	}

	return match
}
