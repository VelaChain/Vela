/* eslint-disable */
import { Params } from "../market/params";
import { Pool } from "../market/pool";
import { LiqProv } from "../market/liq_prov";
import { FeeMap } from "../market/fee_map";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "VelaChain.vela.market";

/** GenesisState defines the market module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  port_id: string;
  poolList: Pool[];
  liqProvList: LiqProv[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  feeMapList: FeeMap[];
}

const baseGenesisState: object = { port_id: "" };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.port_id !== "") {
      writer.uint32(18).string(message.port_id);
    }
    for (const v of message.poolList) {
      Pool.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.liqProvList) {
      LiqProv.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.feeMapList) {
      FeeMap.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.poolList = [];
    message.liqProvList = [];
    message.feeMapList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.port_id = reader.string();
          break;
        case 3:
          message.poolList.push(Pool.decode(reader, reader.uint32()));
          break;
        case 4:
          message.liqProvList.push(LiqProv.decode(reader, reader.uint32()));
          break;
        case 5:
          message.feeMapList.push(FeeMap.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.poolList = [];
    message.liqProvList = [];
    message.feeMapList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = String(object.port_id);
    } else {
      message.port_id = "";
    }
    if (object.poolList !== undefined && object.poolList !== null) {
      for (const e of object.poolList) {
        message.poolList.push(Pool.fromJSON(e));
      }
    }
    if (object.liqProvList !== undefined && object.liqProvList !== null) {
      for (const e of object.liqProvList) {
        message.liqProvList.push(LiqProv.fromJSON(e));
      }
    }
    if (object.feeMapList !== undefined && object.feeMapList !== null) {
      for (const e of object.feeMapList) {
        message.feeMapList.push(FeeMap.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.port_id !== undefined && (obj.port_id = message.port_id);
    if (message.poolList) {
      obj.poolList = message.poolList.map((e) =>
        e ? Pool.toJSON(e) : undefined
      );
    } else {
      obj.poolList = [];
    }
    if (message.liqProvList) {
      obj.liqProvList = message.liqProvList.map((e) =>
        e ? LiqProv.toJSON(e) : undefined
      );
    } else {
      obj.liqProvList = [];
    }
    if (message.feeMapList) {
      obj.feeMapList = message.feeMapList.map((e) =>
        e ? FeeMap.toJSON(e) : undefined
      );
    } else {
      obj.feeMapList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.poolList = [];
    message.liqProvList = [];
    message.feeMapList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = object.port_id;
    } else {
      message.port_id = "";
    }
    if (object.poolList !== undefined && object.poolList !== null) {
      for (const e of object.poolList) {
        message.poolList.push(Pool.fromPartial(e));
      }
    }
    if (object.liqProvList !== undefined && object.liqProvList !== null) {
      for (const e of object.liqProvList) {
        message.liqProvList.push(LiqProv.fromPartial(e));
      }
    }
    if (object.feeMapList !== undefined && object.feeMapList !== null) {
      for (const e of object.feeMapList) {
        message.feeMapList.push(FeeMap.fromPartial(e));
      }
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
