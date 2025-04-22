# Basic Commands

`daku init` initializes the database.

`daku csv` imports data from csv files.

`-a` add flag to override database address.

# Configuration

`DAKU_DB_ADDRESS` holds the database location.

`DAKU_SQL_SERVICE` hold the database service, only supports 'postgres' and 'sqlite3' values.

# CSV Formatting

Any CSVs used to input data must use the proper names in their first column.
The names can be placed in any order relative to one another.
Many columns can be safely excluded.

## Players

Available Columns: name_first

'name_first' contains the names of players - string.

## Games

Available columns: name,ties_possible,tie_breakers,score_kept,round_extensions

'name' contains the names of games - string.

'ties_possible' communicates if the game has a built in concept of ties - boolean.

'tie_breakers' communicates if the game has a built in method for breaking ties - boolean.

'score_kept' communicates if the game has scoring or just wins and losses - boolean.

'round_extensions' communicates whether or not the game has round extensions, as in allows for continuing play without reseting the score - boolean.

## Match Data

Available columns: game_id,rounds,datetime,player_count,id

'game_id' contains the id number for the game - integer.

'rounds' contains the number of rounds or extensions in a given match - integer.

'datetime' contains the date and time information for when a match was ended - timestamp.

'player_count' contains the number of players in a match - integer.

'id' contains id for the match, pair 'match_id' from Player Data with this - integer.

## Player Data

Available columns: player_id,match_id,points,win,ties,round_number

'player_id' contains 'id' from Players corresponding to the player's name - integer.

'match_id' contains the 'id' from Match Data corresponding to the appropriate match - integer.

'points' contains the point count gathered at the end of the round or extension - integer.

'win' communicates if the player won the match - boolean.

'ties' contains the number of ties for cases where a game has built in tie breakers - integer.

'round_number' The round or extension number - integer.
