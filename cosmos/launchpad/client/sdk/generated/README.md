# Go API client for openapi

A REST interface for state queries, transaction generation and broadcasting.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://stargate.cosmos.network*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthApi* | [**AuthAccountsAddressGet**](docs/AuthApi.md#authaccountsaddressget) | **Get** /auth/accounts/{address} | Get the account information on blockchain
*BankApi* | [**BankBalancesAddressGet**](docs/BankApi.md#bankbalancesaddressget) | **Get** /bank/balances/{address} | Get the account balances
*TendermintRPCApi* | [**BlocksHeightGet**](docs/TendermintRPCApi.md#blocksheightget) | **Get** /blocks/{height} | Get a block at a certain height
*TendermintRPCApi* | [**BlocksLatestGet**](docs/TendermintRPCApi.md#blockslatestget) | **Get** /blocks/latest | Get the latest block
*TendermintRPCApi* | [**NodeInfoGet**](docs/TendermintRPCApi.md#nodeinfoget) | **Get** /node_info | The properties of the connected node
*TendermintRPCApi* | [**SyncingGet**](docs/TendermintRPCApi.md#syncingget) | **Get** /syncing | Syncing state of node
*TendermintRPCApi* | [**ValidatorsetsHeightGet**](docs/TendermintRPCApi.md#validatorsetsheightget) | **Get** /validatorsets/{height} | Get a validator set a certain height
*TendermintRPCApi* | [**ValidatorsetsLatestGet**](docs/TendermintRPCApi.md#validatorsetslatestget) | **Get** /validatorsets/latest | Get the latest validator set
*TransactionsApi* | [**TxsDecodePost**](docs/TransactionsApi.md#txsdecodepost) | **Post** /txs/decode | Decode a transaction from the Amino wire format
*TransactionsApi* | [**TxsEncodePost**](docs/TransactionsApi.md#txsencodepost) | **Post** /txs/encode | Encode a transaction to the Amino wire format
*TransactionsApi* | [**TxsGet**](docs/TransactionsApi.md#txsget) | **Get** /txs | Search transactions
*TransactionsApi* | [**TxsHashGet**](docs/TransactionsApi.md#txshashget) | **Get** /txs/{hash} | Get a Tx by hash
*TransactionsApi* | [**TxsPost**](docs/TransactionsApi.md#txspost) | **Post** /txs | Broadcast a signed tx


## Documentation For Models

 - [Block](docs/Block.md)
 - [BlockHeader](docs/BlockHeader.md)
 - [BlockHeaderVersion](docs/BlockHeaderVersion.md)
 - [BlockId](docs/BlockId.md)
 - [BlockIdParts](docs/BlockIdParts.md)
 - [BlockLastCommit](docs/BlockLastCommit.md)
 - [BlockLastCommitPrecommits](docs/BlockLastCommitPrecommits.md)
 - [BlockQuery](docs/BlockQuery.md)
 - [BlockQueryBlockMeta](docs/BlockQueryBlockMeta.md)
 - [BroadcastTxCommitResult](docs/BroadcastTxCommitResult.md)
 - [CheckTxResult](docs/CheckTxResult.md)
 - [Coin](docs/Coin.md)
 - [DeliverTxResult](docs/DeliverTxResult.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2006Result](docs/InlineResponse2006Result.md)
 - [InlineResponse2006ResultValue](docs/InlineResponse2006ResultValue.md)
 - [InlineResponse200ApplicationVersion](docs/InlineResponse200ApplicationVersion.md)
 - [InlineResponse200NodeInfo](docs/InlineResponse200NodeInfo.md)
 - [InlineResponse200NodeInfoOther](docs/InlineResponse200NodeInfoOther.md)
 - [InlineResponse200NodeInfoProtocolVersion](docs/InlineResponse200NodeInfoProtocolVersion.md)
 - [KvPair](docs/KvPair.md)
 - [Msg](docs/Msg.md)
 - [MsgValue](docs/MsgValue.md)
 - [PaginatedQueryTxs](docs/PaginatedQueryTxs.md)
 - [PublicKey](docs/PublicKey.md)
 - [StdTx](docs/StdTx.md)
 - [StdTxValue](docs/StdTxValue.md)
 - [StdTxValueFee](docs/StdTxValueFee.md)
 - [StdTxValuePubKey](docs/StdTxValuePubKey.md)
 - [StdTxValueSignatures](docs/StdTxValueSignatures.md)
 - [TendermintValidator](docs/TendermintValidator.md)
 - [TxQuery](docs/TxQuery.md)
 - [TxQueryResult](docs/TxQueryResult.md)


## Documentation For Authorization



## kms

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```



## Author



