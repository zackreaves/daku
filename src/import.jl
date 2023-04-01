using DataFrames, CSV, Dates, SQLite

function name2num(n::Int)
    if n == 3
        return 1
    elseif n == 4
        return 2
    elseif n == 5
        return 3
    elseif n == 6
        return 4
    elseif n == 7
        return 5
    end
end

function scoreimport(scorecsv::String) # Short term tool for importing data kept in a very specific format that shouldn't be expect to be generalized.
    df = DataFrame(CSV.File(scorecsv))

    # Aaron: 3
    # Will: 4
    # Bailey: 5
    # Zack: 6
    # Ryan: 7

    # zack = df[!,6]
    # bailey = df[!,5]
    # ryan = df[!,7]
    # will = df[!,4]
    # aaron = df[!,3]
    daytime = df[!,1] + df[!,2]

    

    for i in 1:length(df)

        dt = daytime[i]
        DBInterface.execute["INSERT "] # FIXME: finish this up.
        for k in 3:7
            if k != missing
                [ name2num(k) df[k,i] ]
            end
        end
    end
end
