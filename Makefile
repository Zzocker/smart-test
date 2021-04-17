build:
	abigen --sol Ballot/Ballot.sol --pkg ballot --out Ballot/ballot.go
	abigen --sol OpenAuction/OpenAuction.sol --pkg openauction --out OpenAuction/OpenAuction.go
	abigen --sol  ERC20/contract/erc20.sol --pkg stub --out ERC20/stub/erc20.go
erc20:
	go build -o ERC20/bin/deploy ERC20/cmd/deploy.go