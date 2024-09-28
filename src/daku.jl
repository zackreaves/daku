module daku

export initdb
include("init.jl")

export scoreimport, name2num
include("import.jl")

export db_ins_prompt, db_insert
include("daku_write.jl")

end
