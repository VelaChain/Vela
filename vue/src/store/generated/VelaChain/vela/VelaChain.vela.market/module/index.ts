// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgJoinPool } from "./types/market/tx";
import { MsgExitPool } from "./types/market/tx";
import { MsgSendShares } from "./types/market/tx";
import { MsgAddLiquidity } from "./types/market/tx";
import { MsgRemoveLiquidity } from "./types/market/tx";
import { MsgSwap } from "./types/market/tx";
import { MsgCreatePool } from "./types/market/tx";


const types = [
  ["/VelaChain.vela.market.MsgJoinPool", MsgJoinPool],
  ["/VelaChain.vela.market.MsgExitPool", MsgExitPool],
  ["/VelaChain.vela.market.MsgSendShares", MsgSendShares],
  ["/VelaChain.vela.market.MsgAddLiquidity", MsgAddLiquidity],
  ["/VelaChain.vela.market.MsgRemoveLiquidity", MsgRemoveLiquidity],
  ["/VelaChain.vela.market.MsgSwap", MsgSwap],
  ["/VelaChain.vela.market.MsgCreatePool", MsgCreatePool],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgJoinPool: (data: MsgJoinPool): EncodeObject => ({ typeUrl: "/VelaChain.vela.market.MsgJoinPool", value: MsgJoinPool.fromPartial( data ) }),
    msgExitPool: (data: MsgExitPool): EncodeObject => ({ typeUrl: "/VelaChain.vela.market.MsgExitPool", value: MsgExitPool.fromPartial( data ) }),
    msgSendShares: (data: MsgSendShares): EncodeObject => ({ typeUrl: "/VelaChain.vela.market.MsgSendShares", value: MsgSendShares.fromPartial( data ) }),
    msgAddLiquidity: (data: MsgAddLiquidity): EncodeObject => ({ typeUrl: "/VelaChain.vela.market.MsgAddLiquidity", value: MsgAddLiquidity.fromPartial( data ) }),
    msgRemoveLiquidity: (data: MsgRemoveLiquidity): EncodeObject => ({ typeUrl: "/VelaChain.vela.market.MsgRemoveLiquidity", value: MsgRemoveLiquidity.fromPartial( data ) }),
    msgSwap: (data: MsgSwap): EncodeObject => ({ typeUrl: "/VelaChain.vela.market.MsgSwap", value: MsgSwap.fromPartial( data ) }),
    msgCreatePool: (data: MsgCreatePool): EncodeObject => ({ typeUrl: "/VelaChain.vela.market.MsgCreatePool", value: MsgCreatePool.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
