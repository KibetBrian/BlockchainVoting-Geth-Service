run:
	go run main.go
tidy:
	go mod tidy

#Compile to get bin and abi
compile:
	solc --bin --abi ./contracts/Election.sol -o ./build

#Generate golang code from the abi
abigengo:
	abigen --abi ./build/Election.abi --pkg election --out=election/Election.go --bin ./build/Election.bin

#Deploy contract to the network
deploy:
	go run deployment/main.go