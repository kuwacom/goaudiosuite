module example.com/example

go 1.22.5

replace local.packages/goaudiosuite => ../

require (
	local.packages/goaudiosuite v0.0.0-00010101000000-000000000000
	gopkg.in/hraban/opus.v2 v2.0.0-20230925203106-0188a62cb302 // indirect
)
