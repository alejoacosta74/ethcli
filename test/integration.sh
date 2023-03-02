#! /bin/bash

# This script is used to run the integration tests for the ethcli application
# It requires docker to be installed and running
# A ganache-cli docker container is started and the tests are run against it

set -eu pipefail


trap onError ERR

onError(){
	echo "error: Script failed: see failed command above"
	cleanup
    	exit 1
}

cleanup(){
	if [ -n "$(docker ps -q -f name=ganache-cli)" ]; then
		docker stop ganache-cli > /dev/null
	fi
	if [ -n "$(docker ps -aq -f name=ganache-cli)" ]; then
		docker rm ganache-cli > /dev/null
	fi


}

# set environment variable with the path to the ethcli binary
export PATH=$PATH:$(pwd)/bin

# Start ganache-cli docker container
PRIVATEKEY="0x85cbc7b1adfe877051d746c3996a01c2bc3e7a6988490439b1f4b4c2b465322d"
BALANCE="100000000000000000000"
PLATFORM=$(uname -m)

## GANACHE options
PORT="--p 8545"
WALLET="--account=$PRIVATEKEY,$BALANCE"

## check if the docker image is available
if [ -z "$(docker images -q trufflesuite/ganache-cli:latest 2> /dev/null)" ]; then
	docker pull --platform $PLATFORM trufflesuite/ganache-cli:latest
	if [ $? -ne 0 ]; then
		echo "error: docker pull failed"
		exit 1
	fi
fi

## check if the docker container is already running
if [ -n "$(docker ps -q -f name=ganache-cli)" ]; then
	docker stop ganache-cli > /dev/null
	docker rm ganache-cli > /dev/null
fi
echo "Starting ganache-cli docker container"
docker run -d -p 8545:8545 --name ganache-cli trufflesuite/ganache-cli:latest $PORT $WALLET > /dev/null

echo "...waiting 20 seconds for Ganache to finish starting up"
sleep 20

# check that the ethcli command is available
if [ -z "$(command -v ethcli)" ]; then
	echo "error: ethcli command not found"
	cleanup
	exit 1
fi

# environment variables for ethcli
export KEY=$PRIVATEKEY
export NODE=http://localhost:8545

# run the integration tests

## test subcommand 'send'
AMOUNT=500000
RECEIVER=0xf17f52151EbEF6C7334FAD080c5704D77216b732
echo "Testing ethcli send command:"
echo "....sending $AMOUNT wei to $RECEIVER"
ethcli send $AMOUNT $RECEIVER > /dev/null

if [ $? -ne 0 ]; then
	echo "error: ethcli send failed"
	cleanup
	exit 1
fi
echo "...checking balance of receiver"
json=$(curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0xf17f52151EbEF6C7334FAD080c5704D77216b732", "latest"],"id":1}' $NODE)
result_hex=$(echo $json | grep -o '"result":"[^"]*"' | sed 's/"result":"//;s/"//')
result=$(printf "%d\n" $result_hex)
if [ "$result" != $AMOUNT ]; then
	echo "...balance of $RECEIVER is $result wei. Expected $AMOUNT wei." 
	echo "Testing 'send' failed!"
	cleanup
	exit 1
else
	echo "...balance of $RECEIVER is $result wei. Expected $AMOUNT wei." 
	echo "Testing 'send' passed!"
fi

