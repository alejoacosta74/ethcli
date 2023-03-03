<img src=https://upload.wikimedia.org/wikipedia/commons/d/d0/Eth-diamond-rainbow.png width=50px>

# ethcli

`ethcli` is a golang based ethereum cli used to interact with an Ethereum node via the JSON RPC API

## Subcommands
`ethcli`  supports the following subcommands:
- `decoderawtx` : Decodes a ethereum RLP encoded transaction
   
    - usage
  
    ```
    ethcli decoderawtx <raw tx hex string>
    ```

    - example
  ```
  ❯ ethcli decoderawtx f867808504a817c8008252089471517f86711b4bff4d789ad6fee9a58d8af1c6bb830f42408027a044f69cb0e1417336ec4020225c33877064bac1fd306338caf460fea86d4c6993a05cabcb2707a1d9f4f58d38b177419dc9eec4abe98f9411b4376a01cde136c623
  ```

    (output)
  ```json
  {
    "type":"0x0",
    "nonce":"0x0",
    "gasPrice":"0x4a817c800",
    "maxPriorityFeePerGas":null,
    "maxFeePerGas":null,
    "gas":"0x5208",
    "value":"0xf4240",
    "input":"0x",
    "v":"0x27",
    "r":"0x44f69cb0e1417336ec4020225c33877064bac1fd306338caf460fea86d4c6993",
    "s":"0x5cabcb2707a1d9f4f58d38b177419dc9eec4abe98f9411b4376a01cde136c623",
    "to":"0x71517f86711b4bff4d789ad6fee9a58d8af1c6bb",
    "hash":"0xc6eac2f84d1f5c62312e2887b2dce43e557a57db0857a2d74c60a7f17e6467e1"
  }

  ```

- `getaddr`: Returns the ethereum address from the private key
  
  - usage
  ```
  ethcli getaddr <private key>
  ```

  - example
  ```
  ❯ ethcli getaddr 00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35
  ```

  (output)
  ```
  Address: 0x6Fd56E72373a34bA39Bf4167aF82e7A411BFED47
  ```

- `gettx`: Retrieves the details of an ethereum transaction by issuing *eth_getTransactionByHash* RPC call
  
  - usage
  ```
  ethcli gettx <tx hash>
  ```

  - example
  ```
  ❯ ethcli gettx  "0xfe36ef21c8d742fa1a6407b71dfa954dd1040543d52564be40d524e07e7bbb49"
  ```
  (output)
  ```json
  	  "hash": "0xfe36ef21c8d742fa1a6407b71dfa954dd1040543d52564be40d524e07e7bbb49",
	    "nonce": "0xf7",
	    "blockHash": "0xeaaa08ee92f22f5ab6cd081c04c7d65c3acd6da5bba7e00da435b383015b1665",
	    "blockNumber": "0x201",
	    "transactionIndex": "0x0",
	    "from": "0xa6d2799a4b465805421bd10247386a708f01db03",
	    "to": "0xb5883af3b4ecb3c3bb792cbcefbd1542ee30ba0a",
	    "value": "0x1",
	    "gas": "0x5208",
	    "gasPrice": "0x4a817c800",
	    "input": "0x",
	    "v": "0x2d45",
	    "r": "0x12d19b01932b608ca3614ed8b01254bf4acecfc849a1c1cc28e275b0142e37a8",
	    "s": "0x4b672b3cdc04bf992c295ebd8511ea1511e31517a2663293cce65a243ed93c1"
  ```
- `importkey`: Imports an ethereum private key (in hexadecimal format), into the keystore of the node
  
  - usage
  ```
  ethcli importkey <private key>
  ```

- `send`: send ether to an ethereum address. The `amount` argument is specified in units of wei. A private key must be provided with the `--key` flag (or alternatively via a keystore file with the `--keystore` flag), in order to sign the transaction.
Owner of the private key must have sufficient funds to cover the transaction fee.
  
  - usage 
  ```
  ethcli send <amount> <receiver address> [<sender private key>|<keystore>]
  ```

## Command line flags
To see a list of all available command line flags in `ethcli`, execute:
```sh
./ethcli -h
Use this CLI to interact with the Ethereum blockchain.
Use ethcli --help to see the list of available sub commands.
Must be used with a running Ethereum node with JSON-RPC API enabled.

Usage:
  ethcli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decoderawtx Decode a ethereum RLP encoded raw transaction
  getaddr     Extracts the ethereum address from the private key
  gettx       Get transaction details
  help        Help about any command
  importkey   imports a private key into the keystore of the node
  send        Send ether to an ethereum address

Flags:
      --config string     config file
  -d, --debug             Enable debug logging
  -h, --help              help for ethcli
  -l, --loglevel string   log level (trace, debug, info, warn, error, fatal, panic
  -n, --node string       Ethereum node URL (default "http://127.0.0.1:8545")
  ```

  You can also get help for a specific subcommand by running `ethcli <subcommand> -h`:

  ```bash
  ❯ ./ethcli send --help
Sends the amount of ether to the specified receiver address. 
Usage: ethcli send <amount> <receiver address> {<sender private key> | <keystore filename>}

- Amount shall be specified in units of wei. 
- A private key must be provided with the --key flag or alternatively via a keystore file
with the flag --keystore <filename> .

Usage:
  ethcli send [flags]

Examples:
ethcli send 100 0x71517f86711B4BFf4D789Ad6FEE9a58D8AF1c6bB --key 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef

Flags:
  -h, --help              help for send
  -k, --key string        Private key to sign the transaction
  -s, --keystore string   Keystore file that holds the private key

Global Flags:
      --config string     config file
  -d, --debug             Enable debug logging
  -l, --loglevel string   log level (trace, debug, info, warn, error, fatal, panic
  -n, --node string       Ethereum node URL (default "http://127.0.0.1:8080/proxy")
  ```

### - Reading flags from environment vars

`ethcli` also supports passing flags via environment variables. 

For example, you can set the `--node` flag by setting the `NODE` environment variable as:
```bash
export NODE=http://127.0.0.1:5777
```

### - Reading flags from a configuration file

Alternatively, you can use a configuration file to set the default values for the flags. 

The configuration file must be in YAML format and the default location is `$HOME/.ethcli/config.yaml`. 

You can override the default location by setting the `CONFIG` environment variable or use the `--config` flag.

For instance you can save a private key in the configuration file and use it to sign transactions:
```yaml
node: 'http://127.0.0.1:5777'
key: 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcde
```

## Building
To build the `ethcli` binary, execute the following command:
```bash
make build
```
The binary will be created in the `bin` directory.

## Testing

### Unit tests

To run the unit tests, execute the following command:
```
make unit-test
```

### Integration tests

To run the integration tests, execute the following command:
```
make integration-test
```

Integration tests are executed against a ganache-cli node running on a docker container.

(You must have docker installed on your machine to run the integration tests)


## To Do
- :soon: Implement subcommand `create` (creates a ethereum contract)
- :soon: Implement subcommand `call` (calls a contract function)
- :soon: Implement subcommand `getbalance` (retrieves the balance of an ethereum address)
- :soon: Implement subcommand `verifysig` (verifies a signature of an ethereum signed tx)
- :soon: Increase coverage for unit tests
- :soon: Increase coverage for integration tests