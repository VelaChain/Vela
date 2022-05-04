/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../market/params";
import { Pool } from "../market/pool";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { LiqProv } from "../market/liq_prov";
import { FeeMap } from "../market/fee_map";

export const protobufPackage = "VelaChain.vela.market";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetPoolRequest {
  denomA: string;
  denomB: string;
}

export interface QueryGetPoolResponse {
  pool: Pool | undefined;
}

export interface QueryAllPoolRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPoolResponse {
  pool: Pool[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLiqProvRequest {
  poolName: string;
  address: string;
}

export interface QueryGetLiqProvResponse {
  liqProv: LiqProv | undefined;
}

export interface QueryAllLiqProvRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLiqProvResponse {
  liqProv: LiqProv[];
  pagination: PageResponse | undefined;
}

export interface QueryGetFeeMapRequest {
  poolName: string;
}

export interface QueryGetFeeMapResponse {
  feeMap: FeeMap | undefined;
}

export interface QueryAllFeeMapRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllFeeMapResponse {
  feeMap: FeeMap[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetPoolRequest: object = { denomA: "", denomB: "" };

export const QueryGetPoolRequest = {
  encode(
    message: QueryGetPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.denomA !== "") {
      writer.uint32(10).string(message.denomA);
    }
    if (message.denomB !== "") {
      writer.uint32(18).string(message.denomB);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denomA = reader.string();
          break;
        case 2:
          message.denomB = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPoolRequest {
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
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

  toJSON(message: QueryGetPoolRequest): unknown {
    const obj: any = {};
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetPoolRequest>): QueryGetPoolRequest {
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
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

const baseQueryGetPoolResponse: object = {};

export const QueryGetPoolResponse = {
  encode(
    message: QueryGetPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pool !== undefined) {
      Pool.encode(message.pool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pool = Pool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPoolResponse {
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = Pool.fromJSON(object.pool);
    } else {
      message.pool = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPoolResponse): unknown {
    const obj: any = {};
    message.pool !== undefined &&
      (obj.pool = message.pool ? Pool.toJSON(message.pool) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetPoolResponse>): QueryGetPoolResponse {
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = Pool.fromPartial(object.pool);
    } else {
      message.pool = undefined;
    }
    return message;
  },
};

const baseQueryAllPoolRequest: object = {};

export const QueryAllPoolRequest = {
  encode(
    message: QueryAllPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPoolRequest } as QueryAllPoolRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPoolRequest {
    const message = { ...baseQueryAllPoolRequest } as QueryAllPoolRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPoolRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllPoolRequest>): QueryAllPoolRequest {
    const message = { ...baseQueryAllPoolRequest } as QueryAllPoolRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPoolResponse: object = {};

export const QueryAllPoolResponse = {
  encode(
    message: QueryAllPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.pool) {
      Pool.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPoolResponse } as QueryAllPoolResponse;
    message.pool = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pool.push(Pool.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPoolResponse {
    const message = { ...baseQueryAllPoolResponse } as QueryAllPoolResponse;
    message.pool = [];
    if (object.pool !== undefined && object.pool !== null) {
      for (const e of object.pool) {
        message.pool.push(Pool.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPoolResponse): unknown {
    const obj: any = {};
    if (message.pool) {
      obj.pool = message.pool.map((e) => (e ? Pool.toJSON(e) : undefined));
    } else {
      obj.pool = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllPoolResponse>): QueryAllPoolResponse {
    const message = { ...baseQueryAllPoolResponse } as QueryAllPoolResponse;
    message.pool = [];
    if (object.pool !== undefined && object.pool !== null) {
      for (const e of object.pool) {
        message.pool.push(Pool.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetLiqProvRequest: object = { poolName: "", address: "" };

export const QueryGetLiqProvRequest = {
  encode(
    message: QueryGetLiqProvRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLiqProvRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetLiqProvRequest } as QueryGetLiqProvRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLiqProvRequest {
    const message = { ...baseQueryGetLiqProvRequest } as QueryGetLiqProvRequest;
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
    return message;
  },

  toJSON(message: QueryGetLiqProvRequest): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLiqProvRequest>
  ): QueryGetLiqProvRequest {
    const message = { ...baseQueryGetLiqProvRequest } as QueryGetLiqProvRequest;
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
    return message;
  },
};

const baseQueryGetLiqProvResponse: object = {};

export const QueryGetLiqProvResponse = {
  encode(
    message: QueryGetLiqProvResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.liqProv !== undefined) {
      LiqProv.encode(message.liqProv, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLiqProvResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLiqProvResponse,
    } as QueryGetLiqProvResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.liqProv = LiqProv.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLiqProvResponse {
    const message = {
      ...baseQueryGetLiqProvResponse,
    } as QueryGetLiqProvResponse;
    if (object.liqProv !== undefined && object.liqProv !== null) {
      message.liqProv = LiqProv.fromJSON(object.liqProv);
    } else {
      message.liqProv = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLiqProvResponse): unknown {
    const obj: any = {};
    message.liqProv !== undefined &&
      (obj.liqProv = message.liqProv
        ? LiqProv.toJSON(message.liqProv)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLiqProvResponse>
  ): QueryGetLiqProvResponse {
    const message = {
      ...baseQueryGetLiqProvResponse,
    } as QueryGetLiqProvResponse;
    if (object.liqProv !== undefined && object.liqProv !== null) {
      message.liqProv = LiqProv.fromPartial(object.liqProv);
    } else {
      message.liqProv = undefined;
    }
    return message;
  },
};

const baseQueryAllLiqProvRequest: object = {};

export const QueryAllLiqProvRequest = {
  encode(
    message: QueryAllLiqProvRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllLiqProvRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllLiqProvRequest } as QueryAllLiqProvRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLiqProvRequest {
    const message = { ...baseQueryAllLiqProvRequest } as QueryAllLiqProvRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLiqProvRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLiqProvRequest>
  ): QueryAllLiqProvRequest {
    const message = { ...baseQueryAllLiqProvRequest } as QueryAllLiqProvRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLiqProvResponse: object = {};

export const QueryAllLiqProvResponse = {
  encode(
    message: QueryAllLiqProvResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.liqProv) {
      LiqProv.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllLiqProvResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLiqProvResponse,
    } as QueryAllLiqProvResponse;
    message.liqProv = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.liqProv.push(LiqProv.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLiqProvResponse {
    const message = {
      ...baseQueryAllLiqProvResponse,
    } as QueryAllLiqProvResponse;
    message.liqProv = [];
    if (object.liqProv !== undefined && object.liqProv !== null) {
      for (const e of object.liqProv) {
        message.liqProv.push(LiqProv.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLiqProvResponse): unknown {
    const obj: any = {};
    if (message.liqProv) {
      obj.liqProv = message.liqProv.map((e) =>
        e ? LiqProv.toJSON(e) : undefined
      );
    } else {
      obj.liqProv = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLiqProvResponse>
  ): QueryAllLiqProvResponse {
    const message = {
      ...baseQueryAllLiqProvResponse,
    } as QueryAllLiqProvResponse;
    message.liqProv = [];
    if (object.liqProv !== undefined && object.liqProv !== null) {
      for (const e of object.liqProv) {
        message.liqProv.push(LiqProv.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetFeeMapRequest: object = { poolName: "" };

export const QueryGetFeeMapRequest = {
  encode(
    message: QueryGetFeeMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFeeMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFeeMapRequest } as QueryGetFeeMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeMapRequest {
    const message = { ...baseQueryGetFeeMapRequest } as QueryGetFeeMapRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    return message;
  },

  toJSON(message: QueryGetFeeMapRequest): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFeeMapRequest>
  ): QueryGetFeeMapRequest {
    const message = { ...baseQueryGetFeeMapRequest } as QueryGetFeeMapRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    return message;
  },
};

const baseQueryGetFeeMapResponse: object = {};

export const QueryGetFeeMapResponse = {
  encode(
    message: QueryGetFeeMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.feeMap !== undefined) {
      FeeMap.encode(message.feeMap, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFeeMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFeeMapResponse } as QueryGetFeeMapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.feeMap = FeeMap.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeMapResponse {
    const message = { ...baseQueryGetFeeMapResponse } as QueryGetFeeMapResponse;
    if (object.feeMap !== undefined && object.feeMap !== null) {
      message.feeMap = FeeMap.fromJSON(object.feeMap);
    } else {
      message.feeMap = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFeeMapResponse): unknown {
    const obj: any = {};
    message.feeMap !== undefined &&
      (obj.feeMap = message.feeMap ? FeeMap.toJSON(message.feeMap) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFeeMapResponse>
  ): QueryGetFeeMapResponse {
    const message = { ...baseQueryGetFeeMapResponse } as QueryGetFeeMapResponse;
    if (object.feeMap !== undefined && object.feeMap !== null) {
      message.feeMap = FeeMap.fromPartial(object.feeMap);
    } else {
      message.feeMap = undefined;
    }
    return message;
  },
};

const baseQueryAllFeeMapRequest: object = {};

export const QueryAllFeeMapRequest = {
  encode(
    message: QueryAllFeeMapRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFeeMapRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFeeMapRequest } as QueryAllFeeMapRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFeeMapRequest {
    const message = { ...baseQueryAllFeeMapRequest } as QueryAllFeeMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFeeMapRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFeeMapRequest>
  ): QueryAllFeeMapRequest {
    const message = { ...baseQueryAllFeeMapRequest } as QueryAllFeeMapRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllFeeMapResponse: object = {};

export const QueryAllFeeMapResponse = {
  encode(
    message: QueryAllFeeMapResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.feeMap) {
      FeeMap.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFeeMapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFeeMapResponse } as QueryAllFeeMapResponse;
    message.feeMap = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.feeMap.push(FeeMap.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFeeMapResponse {
    const message = { ...baseQueryAllFeeMapResponse } as QueryAllFeeMapResponse;
    message.feeMap = [];
    if (object.feeMap !== undefined && object.feeMap !== null) {
      for (const e of object.feeMap) {
        message.feeMap.push(FeeMap.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFeeMapResponse): unknown {
    const obj: any = {};
    if (message.feeMap) {
      obj.feeMap = message.feeMap.map((e) =>
        e ? FeeMap.toJSON(e) : undefined
      );
    } else {
      obj.feeMap = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFeeMapResponse>
  ): QueryAllFeeMapResponse {
    const message = { ...baseQueryAllFeeMapResponse } as QueryAllFeeMapResponse;
    message.feeMap = [];
    if (object.feeMap !== undefined && object.feeMap !== null) {
      for (const e of object.feeMap) {
        message.feeMap.push(FeeMap.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Pool by index. */
  Pool(request: QueryGetPoolRequest): Promise<QueryGetPoolResponse>;
  /** Queries a list of Pool items. */
  PoolAll(request: QueryAllPoolRequest): Promise<QueryAllPoolResponse>;
  /** Queries a LiqProv by index. */
  LiqProv(request: QueryGetLiqProvRequest): Promise<QueryGetLiqProvResponse>;
  /** Queries a list of LiqProv items. */
  LiqProvAll(request: QueryAllLiqProvRequest): Promise<QueryAllLiqProvResponse>;
  /** Queries a FeeMap by index. */
  FeeMap(request: QueryGetFeeMapRequest): Promise<QueryGetFeeMapResponse>;
  /** Queries a list of FeeMap items. */
  FeeMapAll(request: QueryAllFeeMapRequest): Promise<QueryAllFeeMapResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Pool(request: QueryGetPoolRequest): Promise<QueryGetPoolResponse> {
    const data = QueryGetPoolRequest.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Query",
      "Pool",
      data
    );
    return promise.then((data) =>
      QueryGetPoolResponse.decode(new Reader(data))
    );
  }

  PoolAll(request: QueryAllPoolRequest): Promise<QueryAllPoolResponse> {
    const data = QueryAllPoolRequest.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Query",
      "PoolAll",
      data
    );
    return promise.then((data) =>
      QueryAllPoolResponse.decode(new Reader(data))
    );
  }

  LiqProv(request: QueryGetLiqProvRequest): Promise<QueryGetLiqProvResponse> {
    const data = QueryGetLiqProvRequest.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Query",
      "LiqProv",
      data
    );
    return promise.then((data) =>
      QueryGetLiqProvResponse.decode(new Reader(data))
    );
  }

  LiqProvAll(
    request: QueryAllLiqProvRequest
  ): Promise<QueryAllLiqProvResponse> {
    const data = QueryAllLiqProvRequest.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Query",
      "LiqProvAll",
      data
    );
    return promise.then((data) =>
      QueryAllLiqProvResponse.decode(new Reader(data))
    );
  }

  FeeMap(request: QueryGetFeeMapRequest): Promise<QueryGetFeeMapResponse> {
    const data = QueryGetFeeMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Query",
      "FeeMap",
      data
    );
    return promise.then((data) =>
      QueryGetFeeMapResponse.decode(new Reader(data))
    );
  }

  FeeMapAll(request: QueryAllFeeMapRequest): Promise<QueryAllFeeMapResponse> {
    const data = QueryAllFeeMapRequest.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.vela.market.Query",
      "FeeMapAll",
      data
    );
    return promise.then((data) =>
      QueryAllFeeMapResponse.decode(new Reader(data))
    );
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
