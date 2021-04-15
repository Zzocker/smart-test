build:
	abigen --sol Ballot/Ballot.sol --pkg ballot --out Ballot/ballot.go
	abigen --sol OpenAuction/OpenAuction.sol --pkg openauction --out OpenAuction/OpenAuction.go