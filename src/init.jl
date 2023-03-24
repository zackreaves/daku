using SQLite
function initdb(database_path)
  db = SQLite.DB(database_path)
  dbinit = String(read("init.sql"))
  SQLite.execute(db,dbinit)
end
