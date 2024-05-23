module github.com/gcsequino/GoOrch/GoOrch

go 1.18

require internal/node v1.0.0
replace internal/node => ./internal/node
require internal/registry v1.0.0
replace internal/registry => ./internal/node