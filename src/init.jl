function initdb(database_folder_path::String="daku/")
  checkdir = readdir(database_folder_path)
  players = "name_first,name_second"
  games = "name,ties_possible,scoring"
  game_data = "rounds,game_id,ties,date_time"
  player_data = "player_id,wins"
  println(database_folder_path,checkdir,"\nIs this directory correct? [y/n]\n")
  if readline() == "y"
    cd() # FIXME: Slapdash solution until I figure out how to confirm directory without it.
    mkdir(database_folder_path)
    write(database_folder_path*"players.csv",players)
    write(database_folder_path*"games.csv",games)
    write(database_folder_path*"player_data.csv",player_data)
    write(database_folder_path*"game_data.csv",game_data)
  else
    @info "Directory - hopefully - unchanged."
  end
  return pwd()
end
