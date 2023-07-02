# CosmWasm support

This package allows for custom queries and custom messages sends from contract.


### What is supported 

- Queries:
  - interqueryResult - Get the result of a registered interchain query by query_id
  - InterchainAccountAddress - Get the interchain account address by owner_id and connection_id
  - RegisteredInterchainQueries - all set of registered interchain queries.
  - Registeredinterquery - registered interchain query with specified query_id
- Messages:
  - RegisterInterchainAccount - register an interchain account
  - SubmitTx - submit a transaction for execution on a remote chain
  - Registerinterquery - register an interchain query
  - Updateinterquery - update an interchain query
  - Removeinterquery - remove an interchain query


## Command line interface (CLI)

- Commands

```sh
  quadrated tx wasm -h
```

- Query

```sh
  quadrated q wasm -h
```

## Tests
