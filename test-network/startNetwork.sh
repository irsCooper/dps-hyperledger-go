#!/bin/bash

echo "[-] Clearing network..."
./network.sh down

sleep 3;

clear

echo "
[+] Networg is cleared
"

sleep 3;

echo "
[+] Starting networg and create channel...
"
./network.sh up createChannel -ca

echo "
[+] Networg is started and channel is created
"

sleep 3;

echo "
[+] Deploy chaincode...
"
./network.sh  deployCC -ccn basic -ccp ../chaincode-go/ -ccv 1 -ccl go -c  mychannel

sleep 3;

echo "
[+] Chaincode is deployed
"

sleep 3;

./network.sh cc invoke -c mychannel -ccic '{"Args":["Init"]}'

echo "[+] Successfully initialized"
