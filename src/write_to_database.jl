using DataFrames, CSV, Query

function choosename(table_column::DataFrame,selection_name::String)
  println("Which $selection_name\?\n",table_column)
  return readline()
end

function daku_write_csv_prompt(database::String="daku/")
  cd() # FIXME: This is a short term solution for confirming the directory where the command is being run.
  players = DataFrame(CSV.File(database*"players.csv"))
  games = DataFrame(CSV.File(database*"games.csv"))
  game_data = DataFrame(CSV.File(database*"game_data.csv"))
  player_data = DataFrame(CSV.File(database*"player_data.csv"))

  println("Which game?\n",game_data)
  println("How many players?")
  
  player_count::Int = readline()

  for i in 1:player_count
    
  end
  
end
