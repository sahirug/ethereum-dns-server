@echo off
title compiling contracts
solc --abi DDNS.sol -o build
abigen --abi=./build/DDNS.abi --pkg=eddns --out=DDNS.go
solc --bin DDNS.sol -o build
abigen --bin=./build/DDNS.bin --abi=./build/DDNS.abi --pkg=eddns --out=DDNS.go