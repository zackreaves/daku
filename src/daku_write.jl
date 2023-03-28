using DataFrames, SQLite, Dates

function choose_name(table_column::DataFrame,selection_name::String)
	println("Which $selection_name\?\n",table_column,"\nEnter associated number")
	return parse(Int,readline())
end

function db_insert(database::SQLite.DB,table::String,column_name,data_placed)
	DBInterface.execute(database,"INSERT INTO $table ($column_name)\nVALUES ($data_placed)")
end

function db_ins_prompt(db_path::String)
	db = SQLite.DB(db_path)
	player_names = DataFrame(DBInterface.execute(db,"SELECT name_first FROM players"))
	game_names = DataFrame(DBInterface.execute(db,"SELECT name FROM games"))
	game = choose_name(game_names,"game")

	println("How many players?")

	player_count = parse(Int,readline())

	player_name_arr = zeros(player_count)
	player_wins_arr = zeros(player_count)

	for i in 1:player_count
		player_name_arr[i] = choose_name(player_names,"player")
		println("How many wins?\n")
		player_wins_arr[i] = parse(Int,readline())
	end

	round_count = sum(player_wins_arr)
	todays_date = now()

	DBInterface("INSERT INTO \"round_data\" (round_count,player_count,date_time,game_id)\nVALUES ($round_count,$player_count,$todays_date,$game")
end
