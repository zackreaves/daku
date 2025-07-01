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
		players_options []huh.Option[uint]
		match Match_data
		player_id []uint
		score_inputs []huh.Input
	)
	player_count_input := huh.NewInput().Title("How Many Players?").Value(&player_count).Validate(func (player_count string) error {
			_,err := strconv.ParseUint(player_count,10,64)
			if err != nil {
				return fmt.Errorf("Requires Natural/Positive Whole Number.")
			}

			return nil
	})

	err := player_count_input.Run()
	if err != nil {
		return match, err
	}

	games, err := Query_games_all(config)
	if err != nil {
		return match, err
	}

	for i := range len(games) {
		option := huh.NewOption(games[i].name,games[i].id)
		games_options = append(games_options, option)
	}

	game_select := huh.NewSelect[uint]().Title("Select Game:").Options(games_options...).Value(&match.game_id)

	err = game_select.Run()
	if err != nil {
		return match, err
	}

	player_count_int,_ := strconv.ParseUint(player_count,10,64)
	match.player_count = uint(player_count_int)

	players, err := Query_players_all(config)
	if err != nil {
		return match, err
	}
	for i := range len(players) {
		option := huh.NewOption(players[i].name_first,players[i].id)
		players_options = append(players_options, option)
	}

	players_select := huh.NewMultiSelect[uint]().Title("Select Players:").Options(players_options...).Limit(int(player_count_int)).Value(&player_id)
	err = players_select.Run()
	if err != nil {
		return match, err
	}

	player_data := make([]Player_data,len(player_id))

	for i := range len(player_data) {
		player_data[i].player_id = player_id[i]

		score_input := huh.NewInput().Validate(func (player_count string) error {
			_,err := strconv.ParseUint(player_count,10,64)
			if err != nil {
				return fmt.Errorf("Requires Natural/Positive Whole Number.")
			}

			return nil
	 	})

		score_inputs = append(score_inputs, *score_input)
	}

	return match, nil
}
