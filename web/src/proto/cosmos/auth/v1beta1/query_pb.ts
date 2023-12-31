// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file cosmos/auth/v1beta1/query.proto (package cosmos.auth.v1beta1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3 } from "@bufbuild/protobuf";
import { PageRequest, PageResponse } from "../../base/query/v1beta1/pagination_pb.js";
import { Params } from "./auth_pb.js";

/**
 * QueryAccountsRequest is the request type for the Query/Accounts RPC method.
 *
 * @generated from message cosmos.auth.v1beta1.QueryAccountsRequest
 */
export class QueryAccountsRequest extends Message<QueryAccountsRequest> {
  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 1;
   */
  pagination?: PageRequest;

  constructor(data?: PartialMessage<QueryAccountsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmos.auth.v1beta1.QueryAccountsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PageRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAccountsRequest {
    return new QueryAccountsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAccountsRequest {
    return new QueryAccountsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAccountsRequest {
    return new QueryAccountsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryAccountsRequest | PlainMessage<QueryAccountsRequest> | undefined, b: QueryAccountsRequest | PlainMessage<QueryAccountsRequest> | undefined): boolean {
    return proto3.util.equals(QueryAccountsRequest, a, b);
  }
}

/**
 * QueryAccountsResponse is the response type for the Query/Accounts RPC method.
 *
 * @generated from message cosmos.auth.v1beta1.QueryAccountsResponse
 */
export class QueryAccountsResponse extends Message<QueryAccountsResponse> {
  /**
   * accounts are the existing accounts
   *
   * @generated from field: repeated google.protobuf.Any accounts = 1;
   */
  accounts: Any[] = [];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;

  constructor(data?: PartialMessage<QueryAccountsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmos.auth.v1beta1.QueryAccountsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "accounts", kind: "message", T: Any, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PageResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAccountsResponse {
    return new QueryAccountsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAccountsResponse {
    return new QueryAccountsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAccountsResponse {
    return new QueryAccountsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryAccountsResponse | PlainMessage<QueryAccountsResponse> | undefined, b: QueryAccountsResponse | PlainMessage<QueryAccountsResponse> | undefined): boolean {
    return proto3.util.equals(QueryAccountsResponse, a, b);
  }
}

/**
 * QueryAccountRequest is the request type for the Query/Account RPC method.
 *
 * @generated from message cosmos.auth.v1beta1.QueryAccountRequest
 */
export class QueryAccountRequest extends Message<QueryAccountRequest> {
  /**
   * address defines the address to query for.
   *
   * @generated from field: string address = 1;
   */
  address = "";

  constructor(data?: PartialMessage<QueryAccountRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmos.auth.v1beta1.QueryAccountRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAccountRequest {
    return new QueryAccountRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAccountRequest {
    return new QueryAccountRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAccountRequest {
    return new QueryAccountRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryAccountRequest | PlainMessage<QueryAccountRequest> | undefined, b: QueryAccountRequest | PlainMessage<QueryAccountRequest> | undefined): boolean {
    return proto3.util.equals(QueryAccountRequest, a, b);
  }
}

/**
 * QueryAccountResponse is the response type for the Query/Account RPC method.
 *
 * @generated from message cosmos.auth.v1beta1.QueryAccountResponse
 */
export class QueryAccountResponse extends Message<QueryAccountResponse> {
  /**
   * account defines the account of the corresponding address.
   *
   * @generated from field: google.protobuf.Any account = 1;
   */
  account?: Any;

  constructor(data?: PartialMessage<QueryAccountResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmos.auth.v1beta1.QueryAccountResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "message", T: Any },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAccountResponse {
    return new QueryAccountResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAccountResponse {
    return new QueryAccountResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAccountResponse {
    return new QueryAccountResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryAccountResponse | PlainMessage<QueryAccountResponse> | undefined, b: QueryAccountResponse | PlainMessage<QueryAccountResponse> | undefined): boolean {
    return proto3.util.equals(QueryAccountResponse, a, b);
  }
}

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method.
 *
 * @generated from message cosmos.auth.v1beta1.QueryParamsRequest
 */
export class QueryParamsRequest extends Message<QueryParamsRequest> {
  constructor(data?: PartialMessage<QueryParamsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmos.auth.v1beta1.QueryParamsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryParamsRequest {
    return new QueryParamsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryParamsRequest {
    return new QueryParamsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryParamsRequest {
    return new QueryParamsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryParamsRequest | PlainMessage<QueryParamsRequest> | undefined, b: QueryParamsRequest | PlainMessage<QueryParamsRequest> | undefined): boolean {
    return proto3.util.equals(QueryParamsRequest, a, b);
  }
}

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 *
 * @generated from message cosmos.auth.v1beta1.QueryParamsResponse
 */
export class QueryParamsResponse extends Message<QueryParamsResponse> {
  /**
   * params defines the parameters of the module.
   *
   * @generated from field: cosmos.auth.v1beta1.Params params = 1;
   */
  params?: Params;

  constructor(data?: PartialMessage<QueryParamsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmos.auth.v1beta1.QueryParamsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "params", kind: "message", T: Params },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryParamsResponse {
    return new QueryParamsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryParamsResponse {
    return new QueryParamsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryParamsResponse {
    return new QueryParamsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryParamsResponse | PlainMessage<QueryParamsResponse> | undefined, b: QueryParamsResponse | PlainMessage<QueryParamsResponse> | undefined): boolean {
    return proto3.util.equals(QueryParamsResponse, a, b);
  }
}

