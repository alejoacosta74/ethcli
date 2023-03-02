# ethcli

`ethcli` is a golang based ethereum CLI used to interact with an Ethereum node via the JSON RPC API

## Subcommands
The following subcommands are available:
- `decoderawtx` : Decodes a ethereum RLP encoded transaction
   
    - usage
    ```
    ethcli decoderawtx <raw tx hex string>
    ```

    - example
  ```
  ethcli decoderawtx f867808504a817c8008252089471517f86711b4bff4d789ad6fee9a58d8af1c6bb830f42408027a044f69cb0e1417336ec4020225c33877064bac1fd306338caf460fea86d4c6993a05cabcb2707a1d9f4f58d38b177419dc9eec4abe98f9411b4376a01cde136c623
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
  ethcli getaddr 00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35
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
  ethcli gettx  "0xfe36ef21c8d742fa1a6407b71dfa954dd1040543d52564be40d524e07e7bbb49"
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

- `send`: send ether to an ethereum address
  
  - usage 
  ```
  ethcli send <amount> <receiver address> [<sender private key>|<keystore>]
  ```

## Testing
To run the unit tests, execute the following command:
```
make unit-test
```
To run the integration tests, execute the following command:
```
make integration-test
```


## To Do
- Implement subcommand `create` (creates a ethereum contract)
- Implement subcommand `call` (calls a contract function)
- Add unit tests