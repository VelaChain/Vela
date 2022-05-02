# vela
**vela** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://starport.com).

## Velad Commands
```
velad tx market create-pool [amount-a] [denom-a] [amount-b] [denom-b] [min-shares]
```
`create-pool` command creates a new pool for the given denom pair, if no such pool already exists, and creates a liquidity provider with min-shares

```
velad tx market join-pool [amount-a] [denom-a] [amount-b] [denom-b] [min-shares]
```
`join-pool` command creates a new liquidity provider with at least min-shares for the pool associated with the given denom pair 

```
velad tx market add-liquidity [amount-a] [denom-a] [amount-b] [denom-b] [min-shares]
```
`add-liquidity` command adds at least min-shares to an existing liquidity provider for the pool associated with the given denom pair

```
velad tx market exit-pool [denom-a] [denom-b] 
```
`exit-pool` removes all shares from the creator's liquidity provider and removes the provider from the pool with the given denom pair

```
velad tx market remove-liquidity [shares] [denom-a] [denom-b]
```
`remove-liquidity` removes some, but not all, shares from the liquidity provider for pool associated with the given denom pair

```
velad tx market swap [amount-in] [denom-in] [min-amount-out] [denom-out]
```
`swap` adds amount-in to pool for at least min-amount-out using amount-out = amount-in x pool-amount-out/pool-amount-in

```
velad tx market send-shares [shares] [denom-a] [denom-b] [address]
```
`send-shares` adds shares to liq prov w/ address and removes shares from creator's liq prov

```
velad q market list-pool
```
`list-pool` lists the state of the current pools

```
velad q market list-liq-prov
``` 
`list-liq-prov` lists the state of the current liquidity providers

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.com).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.com/VelaChain/vela@latest! | sudo bash
```
`VelaChain/vela` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).
