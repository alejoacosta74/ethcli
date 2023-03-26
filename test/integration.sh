#! /bin/bash

# This script is used to run the integration tests for the ethcli application
# It requires docker to be installed and running
# A ganache-cli docker container is started and the tests are run against it


# Global variables
PRIVATEKEY="0x85cbc7b1adfe877051d746c3996a01c2bc3e7a6988490439b1f4b4c2b465322d"
SIGNER_ADDR="0xA6d2799a4b465805421bd10247386a708F01DB03"
BALANCE="100000000000000000000"
PLATFORM=$(uname -m)

## Ganache-cli variables
PORT="--p 8545"
WALLET="--account=$PRIVATEKEY,$BALANCE"


# ethcli variables
export KEY=$PRIVATEKEY
export NODE=http://localhost:8545
export PATH=$PATH:$(pwd)/bin


set -eEu pipefail
trap onError ERR

onError(){
	echo "error: Script failed: see failed command above"
	cleanup
    	exit 1
}

cleanup(){
	# check if environment CLEANUP is set to true
	if [ "${CLEANUP:-false}" = true ]; then
		echo "Cleaning up"
		if [ -n "$(docker ps -aq -f name=ganache-cli)" ]; then
			docker rm -f ganache-cli > /dev/null
		fi
	fi
}

## convert hex to decimal the first argument is the hex number
hex2dec(){
    local dec=$(echo "ibase=16; $(echo "$1" | tr '[:lower:]' '[:upper:]')" | bc)
    echo "$dec"
}

# read balance with curl
read_balance(){
	json=$(curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["'$1'", "latest"],"id":1}' $NODE)
	balance_hex=$(echo $json | grep -o '"result":"[^"]*"' | sed 's/"result":"0x//;s/"//')
	balance=$(hex2dec $balance_hex)
	echo $balance
}

start_ganache(){
	# check GANACHE_RESTART environment variable
	if [ "${GANACHE_RESTART:-true}" = true ]; then
		# check if the docker image is available
		if [ -z "$(docker images -q trufflesuite/ganache-cli:latest 2> /dev/null)" ]; then
			docker pull --platform $PLATFORM trufflesuite/ganache-cli:latest
			if [ $? -ne 0 ]; then
				echo "error: docker pull failed"
				exit 1
			fi
		fi

		# check if the docker container is already running
		if [ -n "$(docker ps -q -f name=ganache-cli)" ]; then
			docker stop ganache-cli > /dev/null
			docker rm ganache-cli > /dev/null
		fi
		echo "Starting ganache-cli docker container"
		docker run -d -p 8545:8545 --name ganache-cli trufflesuite/ganache-cli:latest $PORT $WALLET > /dev/null

		echo "...waiting 20 seconds for Ganache to finish starting up"
		sleep 20
	fi
}

# build ethcli if it is not already available
build_ethcli(){
	if [ ! -f $(pwd)/bin/ethcli ]; then
		echo "Building ethcli"
		../make build > /dev/null
	fi
}

test_getbalance(){
	# test subcommand 'getbalance'
	echo ""
	echo "Testing ethcli getbalance command:"
	local output=$(ethcli getbalance $SIGNER_ADDR)
	local result=$(echo $output | sed 's/[^0-9]*//g')

	local balance=$(read_balance $SIGNER_ADDR)

	echo "...obtained balance of signer $SIGNER_ADDR is $result wei. Expected $balance wei." 
	if [ "$result" != $balance ]; then
		echo "Testing 'getbalance' failed!"
		PASSED=false
	else 
		echo "Testing 'getbalance' passed!"
	fi
}

test_send(){
	# test subcommand 'send'
	echo ""
	AMOUNT=5000000
	RECEIVER=0xf17f52151EbEF6C7334FAD080c5704D77216b732
	local balance=$(read_balance $RECEIVER)

	echo "Testing ethcli send command:"
	# echo "... current balance is $balance wei"
	echo "...sending $AMOUNT wei to $RECEIVER"
	ethcli send $AMOUNT $RECEIVER > /dev/null

	local newbalance=$(read_balance $RECEIVER)
	# echo "...new balance is $newbalance wei"

	expected="$((balance+AMOUNT))"
	echo "...balance of $RECEIVER is $newbalance wei. Expected $expected wei." 
	if [ "$newbalance" != $expected ]; then
		echo "Testing 'send' failed!"
		PASSED=false
	else
		echo "Testing 'send' passed!"
	fi

}

# run the integration tests
run_tests(){
	PASSED=true
	test_getbalance
	test_send

	if [ "$PASSED" = false ]; then
		exit 1
	fi

}

main(){
	build_ethcli
	start_ganache
	run_tests
	cleanup
}

main