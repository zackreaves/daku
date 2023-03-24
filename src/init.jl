using SQLite, DataFrames
function initdb(database_path::String, db_type::String="sqlite")
  if db_type == "sqlite"
    db = SQLite.DB(database_path)
    db_init = String(read("initlite.sql"))
    # players = DataFrame(names = ["id", "namef"], types = ["INTEGER PRIMARY KEY", "CHAR(80)"])
    DBInterface.execute(db,db_init)
  else
    println("Invalid Database Type")
  end
end
