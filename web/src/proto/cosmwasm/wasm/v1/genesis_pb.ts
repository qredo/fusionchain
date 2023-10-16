// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file cosmwasm/wasm/v1/genesis.proto (package cosmwasm.wasm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { CodeInfo, ContractCodeHistoryEntry, ContractInfo, Model, Params } from "./types_pb.js";

/**
 * GenesisState - genesis state of x/wasm
 *
 * @generated from message cosmwasm.wasm.v1.GenesisState
 */
export class GenesisState extends Message<GenesisState> {
  /**
   * @generated from field: cosmwasm.wasm.v1.Params params = 1;
   */
  params?: Params;

  /**
   * @generated from field: repeated cosmwasm.wasm.v1.Code codes = 2;
   */
  codes: Code[] = [];

  /**
   * @generated from field: repeated cosmwasm.wasm.v1.Contract contracts = 3;
   */
  contracts: Contract[] = [];

  /**
   * @generated from field: repeated cosmwasm.wasm.v1.Sequence sequences = 4;
   */
  sequences: Sequence[] = [];

  constructor(data?: PartialMessage<GenesisState>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmwasm.wasm.v1.GenesisState";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "params", kind: "message", T: Params },
    { no: 2, name: "codes", kind: "message", T: Code, repeated: true },
    { no: 3, name: "contracts", kind: "message", T: Contract, repeated: true },
    { no: 4, name: "sequences", kind: "message", T: Sequence, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenesisState {
    return new GenesisState().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenesisState {
    return new GenesisState().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenesisState {
    return new GenesisState().fromJsonString(jsonString, options);
  }

  static equals(a: GenesisState | PlainMessage<GenesisState> | undefined, b: GenesisState | PlainMessage<GenesisState> | undefined): boolean {
    return proto3.util.equals(GenesisState, a, b);
  }
}

/**
 * Code struct encompasses CodeInfo and CodeBytes
 *
 * @generated from message cosmwasm.wasm.v1.Code
 */
export class Code extends Message<Code> {
  /**
   * @generated from field: uint64 code_id = 1;
   */
  codeId = protoInt64.zero;

  /**
   * @generated from field: cosmwasm.wasm.v1.CodeInfo code_info = 2;
   */
  codeInfo?: CodeInfo;

  /**
   * @generated from field: bytes code_bytes = 3;
   */
  codeBytes = new Uint8Array(0);

  /**
   * Pinned to wasmvm cache
   *
   * @generated from field: bool pinned = 4;
   */
  pinned = false;

  constructor(data?: PartialMessage<Code>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmwasm.wasm.v1.Code";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "code_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "code_info", kind: "message", T: CodeInfo },
    { no: 3, name: "code_bytes", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 4, name: "pinned", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Code {
    return new Code().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Code {
    return new Code().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Code {
    return new Code().fromJsonString(jsonString, options);
  }

  static equals(a: Code | PlainMessage<Code> | undefined, b: Code | PlainMessage<Code> | undefined): boolean {
    return proto3.util.equals(Code, a, b);
  }
}

/**
 * Contract struct encompasses ContractAddress, ContractInfo, and ContractState
 *
 * @generated from message cosmwasm.wasm.v1.Contract
 */
export class Contract extends Message<Contract> {
  /**
   * @generated from field: string contract_address = 1;
   */
  contractAddress = "";

  /**
   * @generated from field: cosmwasm.wasm.v1.ContractInfo contract_info = 2;
   */
  contractInfo?: ContractInfo;

  /**
   * @generated from field: repeated cosmwasm.wasm.v1.Model contract_state = 3;
   */
  contractState: Model[] = [];

  /**
   * @generated from field: repeated cosmwasm.wasm.v1.ContractCodeHistoryEntry contract_code_history = 4;
   */
  contractCodeHistory: ContractCodeHistoryEntry[] = [];

  constructor(data?: PartialMessage<Contract>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmwasm.wasm.v1.Contract";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "contract_address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "contract_info", kind: "message", T: ContractInfo },
    { no: 3, name: "contract_state", kind: "message", T: Model, repeated: true },
    { no: 4, name: "contract_code_history", kind: "message", T: ContractCodeHistoryEntry, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Contract {
    return new Contract().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Contract {
    return new Contract().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Contract {
    return new Contract().fromJsonString(jsonString, options);
  }

  static equals(a: Contract | PlainMessage<Contract> | undefined, b: Contract | PlainMessage<Contract> | undefined): boolean {
    return proto3.util.equals(Contract, a, b);
  }
}

/**
 * Sequence key and value of an id generation counter
 *
 * @generated from message cosmwasm.wasm.v1.Sequence
 */
export class Sequence extends Message<Sequence> {
  /**
   * @generated from field: bytes id_key = 1;
   */
  idKey = new Uint8Array(0);

  /**
   * @generated from field: uint64 value = 2;
   */
  value = protoInt64.zero;

  constructor(data?: PartialMessage<Sequence>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmwasm.wasm.v1.Sequence";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id_key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "value", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Sequence {
    return new Sequence().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Sequence {
    return new Sequence().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Sequence {
    return new Sequence().fromJsonString(jsonString, options);
  }

  static equals(a: Sequence | PlainMessage<Sequence> | undefined, b: Sequence | PlainMessage<Sequence> | undefined): boolean {
    return proto3.util.equals(Sequence, a, b);
  }
}

