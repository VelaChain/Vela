/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "VelaChain.vela.market";

export interface LiqProv {
  poolName: string;
  address: string;
  shareAmount: string;
}

const baseLiqProv: object = { poolName: "", address: "", shareAmount: "" };

export const LiqProv = {
  encode(message: LiqProv, writer: Writer = Writer.create()): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.shareAmount !== "") {
      writer.uint32(26).string(message.shareAmount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LiqProv {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLiqProv } as LiqProv;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.shareAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LiqProv {
    const message = { ...baseLiqProv } as LiqProv;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.shareAmount !== undefined && object.shareAmount !== null) {
      message.shareAmount = String(object.shareAmount);
    } else {
      message.shareAmount = "";
    }
    return message;
  },

  toJSON(message: LiqProv): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.address !== undefined && (obj.address = message.address);
    message.shareAmount !== undefined &&
      (obj.shareAmount = message.shareAmount);
    return obj;
  },

  fromPartial(object: DeepPartial<LiqProv>): LiqProv {
    const message = { ...baseLiqProv } as LiqProv;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.shareAmount !== undefined && object.shareAmount !== null) {
      message.shareAmount = object.shareAmount;
    } else {
      message.shareAmount = "";
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
