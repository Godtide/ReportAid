module github.com/ReportAid/app/src/server

go 1.14

require (
	github.com/ethereum/go-ethereum v1.9.13 // indirect
	github.com/ReportAid/app/src/server/internal/abi v0.0.1
)

replace  internal/abi => ../abi
