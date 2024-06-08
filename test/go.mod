module github.com/UedaTakeyuki/erapse/test

go 1.21

replace local.packages/erapse => ../../erapse

replace local.packages/package1 => ./package1

replace local.packages/package2 => ./package2

require (
	github.com/UedaTakeyuki/compare v0.0.0-20240525012659-6bb59cb148f3
	local.packages/erapse v0.0.0-00010101000000-000000000000
)

require (
	local.packages/package1 v0.0.0-00010101000000-000000000000 // indirect
	local.packages/package2 v0.0.0-00010101000000-000000000000 // indirect
)
