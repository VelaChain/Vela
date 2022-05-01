/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "VelaChain.vela.market";

export interface MsgCreatePool {
  creator: string;
  amountA: string;
  denomA: string;
  amountB: string;
  denomB: string;
  minShares: string;
}

export interface MsgCreatePoolResponse {}

export interface MsgJoinPool {
  creator: string;
  amountA: string;
  denomA: string;
  amountB: string;
  denomB: string;
  minShares: string;
}

export interface MsgJoinPoolResponse {}

export interface MsgAddLiquidity {
  creator: string;
  amountA: string;
  denomA: string;
  amountB: string;
  denomB: string;
  minShares: string;
}

export interface MsgAddLiquidityResponse {}

export interface MsgExitPool {
  creator: string;
  denomA: string;
  denomB: string;
}

export interface MsgExitPoolResponse {}

export interface MsgRemoveLiquidity {
  creator: string;
  shares: string;
  denomA: string;
  denomB: string;
}

export interface MsgRemoveLiquidityResponse {}

export interface MsgSwap {
  creator: string;
  amountIn: string;
  denomIn: string;
  minAmountOut: string;
  denomOut: string;
}

export interface MsgSwapResponse {}

const baseMsgCreatePool: object = {
  creator: "",
  amountA: "",
  denomA: "",
  amountB: "",
  denomB: "",
  minShares: "",
};

export const MsgCreatePool = {
  encode(message: MsgCreatePool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amountA !== "") {
      writer.uint32(18).string(message.amountA);
    }
    if (message.denomA !== "") {
      writer.uint32(26).string(message.denomA);
    }
    if (message.amountB !== "") {
      writer.uint32(34).string(message.amountB);
    }
    if (message.denomB !== "") {
      writer.uint32(42).string(message.denomB);
    }
    if (message.minShares !== "") {
      writer.uint32(50).string(message.minShares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amountA = reader.string();
          break;
        case 3:
          message.denomA = reader.string();
          break;
        case 4:
          message.amountB = reader.string();
          break;
        case 5:
          message.denomB = reader.string();
          break;
        case 6:
          message.minShares = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePool {
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = String(object.amountA);
    } else {
      message.amountA = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = String(object.amountB);
    } else {
      message.amountB = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    if (object.minShares !== undefined && object.minShares !== null) {
      message.minShares = String(object.minShares);
    } else {
      message.minShares = "";
    }
    return message;
  },

  toJSON(message: MsgCreatePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amountA !== undefined && (obj.amountA = message.amountA);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.amountB !== undefined && (obj.amountB = message.amountB);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    message.minShares !== undefined && (obj.minShares = message.minShares);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreatePool>): MsgCreatePool {
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = object.amountA;
    } else {
      message.amountA = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = object.amountB;
    } else {
      message.amountB = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    if (object.minShares !== undefined && object.minShares !== null) {
      message.minShares = object.minShares;
    } else {
      message.minShares = "";
    }
    return message;
  },
};

const baseMsgCreatePoolResponse: object = {};

export const MsgCreatePoolResponse = {
  encode(_: MsgCreatePoolResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreatePoolResponse {
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
    return message;
  },

  toJSON(_: MsgCreatePoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreatePoolResponse>): MsgCreatePoolResponse {
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
    return message;
  },
};

const baseMsgJoinPool: object = {
  creator: "",
  amountA: "",
  denomA: "",
  amountB: "",
  denomB: "",
  minShares: "",
};

export const MsgJoinPool = {
  encode(message: MsgJoinPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amountA !== "") {
      writer.uint32(18).string(message.amountA);
    }
    if (message.denomA !== "") {
      writer.uint32(26).string(message.denomA);
    }
    if (message.amountB !== "") {
      writer.uint32(34).string(message.amountB);
    }
    if (message.denomB !== "") {
      writer.uint32(42).string(message.denomB);
    }
    if (message.minShares !== "") {
      writer.uint32(50).string(message.minShares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgJoinPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgJoinPool } as MsgJoinPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amountA = reader.string();
          break;
        case 3:
          message.denomA = reader.string();
          break;
        case 4:
          message.amountB = reader.string();
          break;
        case 5:
          message.denomB = reader.string();
          break;
        case 6:
          message.minShares = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgJoinPool {
    const message = { ...baseMsgJoinPool } as MsgJoinPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = String(object.amountA);
    } else {
      message.amountA = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = String(object.amountB);
    } else {
      message.amountB = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    if (object.minShares !== undefined && object.minShares !== null) {
      message.minShares = String(object.minShares);
    } else {
      message.minShares = "";
    }
    return message;
  },

  toJSON(message: MsgJoinPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amountA !== undefined && (obj.amountA = message.amountA);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.amountB !== undefined && (obj.amountB = message.amountB);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    message.minShares !== undefined && (obj.minShares = message.minShares);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgJoinPool>): MsgJoinPool {
    const message = { ...baseMsgJoinPool } as MsgJoinPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = object.amountA;
    } else {
      message.amountA = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = object.amountB;
    } else {
      message.amountB = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    if (object.minShares !== undefined && object.minShares !== null) {
      message.minShares = object.minShares;
    } else {
      message.minShares = "";
    }
    return message;
  },
};

const baseMsgJoinPoolResponse: object = {};

export const MsgJoinPoolResponse = {
  encode(_: MsgJoinPoolResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgJoinPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgJoinPoolResponse } as MsgJoinPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgJoinPoolResponse {
    const message = { ...baseMsgJoinPoolResponse } as MsgJoinPoolResponse;
    return message;
  },

  toJSON(_: MsgJoinPoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgJoinPoolResponse>): MsgJoinPoolResponse {
    const message = { ...baseMsgJoinPoolResponse } as MsgJoinPoolResponse;
    return message;
  },
};

const baseMsgAddLiquidity: object = {
  creator: "",
  amountA: "",
  denomA: "",
  amountB: "",
  denomB: "",
  minShares: "",
};

export const MsgAddLiquidity = {
  encode(message: MsgAddLiquidity, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amountA !== "") {
      writer.uint32(18).string(message.amountA);
    }
    if (message.denomA !== "") {
      writer.uint32(26).string(message.denomA);
    }
    if (message.amountB !== "") {
      writer.uint32(34).string(message.amountB);
    }
    if (message.denomB !== "") {
      writer.uint32(42).string(message.denomB);
    }
    if (message.minShares !== "") {
      writer.uint32(50).string(message.minShares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddLiquidity {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddLiquidity } as MsgAddLiquidity;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amountA = reader.string();
          break;
        case 3:
          message.denomA = reader.string();
          break;
        case 4:
          message.amountB = reader.string();
          break;
        case 5:
          message.denomB = reader.string();
          break;
        case 6:
          message.minShares = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddLiquidity {
    const message = { ...baseMsgAddLiquidity } as MsgAddLiquidity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = String(object.amountA);
    } else {
      message.amountA = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = String(object.amountB);
    } else {
      message.amountB = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    if (object.minShares !== undefined && object.minShares !== null) {
      message.minShares = String(object.minShares);
    } else {
      message.minShares = "";
    }
    return message;
  },

  toJSON(message: MsgAddLiquidity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amountA !== undefined && (obj.amountA = message.amountA);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.amountB !== undefined && (obj.amountB = message.amountB);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    message.minShares !== undefined && (obj.minShares = message.minShares);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddLiquidity>): MsgAddLiquidity {
    const message = { ...baseMsgAddLiquidity } as MsgAddLiquidity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = object.amountA;
    } else {
      message.amountA = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = object.amountB;
    } else {
      message.amountB = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    if (object.minShares !== undefined && object.minShares !== null) {
      message.minShares = object.minShares;
    } else {
      message.minShares = "";
    }
    return message;
  },
};

const baseMsgAddLiquidityResponse: object = {};

export const MsgAddLiquidityResponse = {
  encode(_: MsgAddLiquidityResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddLiquidityResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddLiquidityResponse,
    } as MsgAddLiquidityResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddLiquidityResponse {
    const message = {
      ...baseMsgAddLiquidityResponse,
    } as MsgAddLiquidityResponse;
    return message;
  },

  toJSON(_: MsgAddLiquidityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddLiquidityResponse>
  ): MsgAddLiquidityResponse {
    const message = {
      ...baseMsgAddLiquidityResponse,
    } as MsgAddLiquidityResponse;
    return message;
  },
};

const baseMsgExitPool: object = { creator: "", denomA: "", denomB: "" };

export const MsgExitPool = {
  encode(message: MsgExitPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denomA !== "") {
      writer.uint32(18).string(message.denomA);
    }
    if (message.denomB !== "") {
      writer.uint32(26).string(message.denomB);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgExitPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgExitPool } as MsgExitPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denomA = reader.string();
          break;
        case 3:
          message.denomB = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgExitPool {
    const message = { ...baseMsgExitPool } as MsgExitPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    return message;
  },

  toJSON(message: MsgExitPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgExitPool>): MsgExitPool {
    const message = { ...baseMsgExitPool } as MsgExitPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    return message;
  },
};

const baseMsgExitPoolResponse: object = {};

export const MsgExitPoolResponse = {
  encode(_: MsgExitPoolResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgExitPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgExitPoolResponse } as MsgExitPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgExitPoolResponse {
    const message = { ...baseMsgExitPoolResponse } as MsgExitPoolResponse;
    return message;
  },

  toJSON(_: MsgExitPoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgExitPoolResponse>): MsgExitPoolResponse {
    const message = { ...baseMsgExitPoolResponse } as MsgExitPoolResponse;
    return message;
  },
};

const baseMsgRemoveLiquidity: object = {
  creator: "",
  shares: "",
  denomA: "",
  denomB: "",
};

export const MsgRemoveLiquidity = {
  encode(
    message: MsgRemoveLiquidity,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.shares !== "") {
      writer.uint32(18).string(message.shares);
    }
    if (message.denomA !== "") {
      writer.uint32(26).string(message.denomA);
    }
    if (message.denomB !== "") {
      writer.uint32(34).string(message.denomB);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRemoveLiquidity {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRemoveLiquidity } as MsgRemoveLiquidity;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.shares = reader.string();
          break;
        case 3:
          message.denomA = reader.string();
          break;
        case 4:
          message.denomB = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveLiquidity {
    const message = { ...baseMsgRemoveLiquidity } as MsgRemoveLiquidity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = String(object.shares);
    } else {
      message.shares = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    return message;
  },

  toJSON(message: MsgRemoveLiquidity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.shares !== undefined && (obj.shares = message.shares);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRemoveLiquidity>): MsgRemoveLiquidity {
    const message = { ...baseMsgRemoveLiquidity } as MsgRemoveLiquidity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = object.shares;
    } else {
      message.shares = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    return message;
  },
};

const baseMsgRemoveLiquidityResponse: object = {};

export const MsgRemoveLiquidityResponse = {
  encode(
    _: MsgRemoveLiquidityResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveLiquidityResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveLiquidityResponse,
    } as MsgRemoveLiquidityResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRemoveLiquidityResponse {
    const message = {
      ...baseMsgRemoveLiquidityResponse,
    } as MsgRemoveLiquidityResponse;
    return message;
  },

  toJSON(_: MsgRemoveLiquidityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRemoveLiquidityResponse>
  ): MsgRemoveLiquidityResponse {
    const message = {
      ...baseMsgRemoveLiquidityResponse,
    } as MsgRemoveLiquidityResponse;
    return message;
  },
};

const baseMsgSwap: object = {
  creator: "",
  amountIn: "",
  denomIn: "",
  minAmountOut: "",
  denomOut: "",
};

export const MsgSwap = {
  encode(message: MsgSwap, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amountIn !== "") {
      writer.uint32(18).string(message.amountIn);
    }
    if (message.denomIn !== "") {
      writer.uint32(26).string(message.denomIn);
    }
    if (message.minAmountOut !== "") {
      writer.uint32(34).string(message.minAmountOut);
    }
    if (message.denomOut !== "") {
      writer.uint32(42).string(message.denomOut);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwap } as MsgSwap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amountIn = reader.string();
          break;
        case 3:
          message.denomIn = reader.string();
          break;
        case 4:
          message.minAmountOut = reader.string();
          break;
        case 5:
          message.denomOut = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = String(object.amountIn);
    } else {
      message.amountIn = "";
    }
    if (object.denomIn !== undefined && object.denomIn !== null) {
      message.denomIn = String(object.denomIn);
    } else {
      message.denomIn = "";
    }
    if (object.minAmountOut !== undefined && object.minAmountOut !== null) {
      message.minAmountOut = String(object.minAmountOut);
    } else {
      message.minAmountOut = "";
    }
    if (object.denomOut !== undefined && object.denomOut !== null) {
      message.denomOut = String(object.denomOut);
    } else {
      message.denomOut = "";
    }
    return message;
  },

  toJSON(message: MsgSwap): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amountIn !== undefined && (obj.amountIn = message.amountIn);
    message.denomIn !== undefined && (obj.denomIn = message.denomIn);
    message.minAmountOut !== undefined &&
      (obj.minAmountOut = message.minAmountOut);
    message.denomOut !== undefined && (obj.denomOut = message.denomOut);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwap>): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = object.amountIn;
    } else {
      message.amountIn = "";
    }
    if (object.denomIn !== undefined && object.denomIn !== null) {
      message.denomIn = object.denomIn;
    } else {
      message.denomIn = "";
    }
    if (object.minAmountOut !== undefined && object.minAmountOut !== null) {
      message.minAmountOut = object.minAmountOut;
    } else {
      message.minAmountOut = "";
    }
    if (object.denomOut !== undefined && object.denomOut !== null) {
      message.denomOut = object.denomOut;
    } else {
      message.denomOut = "";
    }
    return message;
  },
};

const baseMsgSwapResponse: object = {};

export const MsgSwapResponse = {
  encode(_: MsgSwapResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    return message;
  },

  toJSON(_: MsgSwapResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSwapResponse>): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse>;
  JoinPool(request: MsgJoinPool): Promise<MsgJoinPoolResponse>;
  AddLiquidity(request: MsgAddLiquidity): Promise<MsgAddLiquidityResponse>;
  ExitPool(request: MsgExitPool): Promise<MsgExitPoolResponse>;
  RemoveLiquidity(
    request: MsgRemoveLiquidity
  ): Promise<MsgRemoveLiquidityResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Swap(request: MsgSwap): Promise<MsgSwapResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse> {
    const data = MsgCreatePool.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Msg",
      "CreatePool",
      data
    );
    return promise.then((data) =>
      MsgCreatePoolResponse.decode(new Reader(data))
    );
  }

  JoinPool(request: MsgJoinPool): Promise<MsgJoinPoolResponse> {
    const data = MsgJoinPool.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Msg",
      "JoinPool",
      data
    );
    return promise.then((data) => MsgJoinPoolResponse.decode(new Reader(data)));
  }

  AddLiquidity(request: MsgAddLiquidity): Promise<MsgAddLiquidityResponse> {
    const data = MsgAddLiquidity.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Msg",
      "AddLiquidity",
      data
    );
    return promise.then((data) =>
      MsgAddLiquidityResponse.decode(new Reader(data))
    );
  }

  ExitPool(request: MsgExitPool): Promise<MsgExitPoolResponse> {
    const data = MsgExitPool.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Msg",
      "ExitPool",
      data
    );
    return promise.then((data) => MsgExitPoolResponse.decode(new Reader(data)));
  }

  RemoveLiquidity(
    request: MsgRemoveLiquidity
  ): Promise<MsgRemoveLiquidityResponse> {
    const data = MsgRemoveLiquidity.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Msg",
      "RemoveLiquidity",
      data
    );
    return promise.then((data) =>
      MsgRemoveLiquidityResponse.decode(new Reader(data))
    );
  }

  Swap(request: MsgSwap): Promise<MsgSwapResponse> {
    const data = MsgSwap.encode(request).finish();
    const promise = this.rpc.request("VelaChain.vela.market.Msg", "Swap", data);
    return promise.then((data) => MsgSwapResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
