/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "VelaChain.vela.market";

export interface FeeMap {
  poolName: string;
  swap: string;
  exit: string;
}

const baseFeeMap: object = { poolName: "", swap: "", exit: "" };

export const FeeMap = {
  encode(message: FeeMap, writer: Writer = Writer.create()): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    if (message.swap !== "") {
      writer.uint32(18).string(message.swap);
    }
    if (message.exit !== "") {
      writer.uint32(26).string(message.exit);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): FeeMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFeeMap } as FeeMap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        case 2:
          message.swap = reader.string();
          break;
        case 3:
          message.exit = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeeMap {
    const message = { ...baseFeeMap } as FeeMap;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.swap !== undefined && object.swap !== null) {
      message.swap = String(object.swap);
    } else {
      message.swap = "";
    }
    if (object.exit !== undefined && object.exit !== null) {
      message.exit = String(object.exit);
    } else {
      message.exit = "";
    }
    return message;
  },

  toJSON(message: FeeMap): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.swap !== undefined && (obj.swap = message.swap);
    message.exit !== undefined && (obj.exit = message.exit);
    return obj;
  },

  fromPartial(object: DeepPartial<FeeMap>): FeeMap {
    const message = { ...baseFeeMap } as FeeMap;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.swap !== undefined && object.swap !== null) {
      message.swap = object.swap;
    } else {
      message.swap = "";
    }
    if (object.exit !== undefined && object.exit !== null) {
      message.exit = object.exit;
    } else {
      message.exit = "";
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
