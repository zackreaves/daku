using SQLite, DataFrames
function initdb(database_path::String, db_type::String="sqlite")
  if db_type == "sqlite"
    db = SQLite.DB(database_path)
    db_init = String(read("init.sql"))
    DBInterface.execute(db,db_init)
  else
    @error "Invalid Database Type"
  end
end
