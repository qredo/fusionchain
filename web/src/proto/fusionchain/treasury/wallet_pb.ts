// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file fusionchain/treasury/wallet.proto (package fusionchain.treasury, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * WalletType specifies the Layer 1 blockchain that this wallet will be used
 * for.
 *
 * @generated from enum fusionchain.treasury.WalletType
 */
export enum WalletType {
  /**
   * The wallet type is missing
   *
   * @generated from enum value: WALLET_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * The wallet type for native Fusion chain cosmos accounts (not ERC-20 QRDO
   * tokens)
   *
   * @generated from enum value: WALLET_TYPE_QRDO = 1;
   */
  QRDO = 1,

  /**
   * The wallet type for mainnet ETH and its ERC-20 tokens (including non-native
   * QRDO)
   *
   * @generated from enum value: WALLET_TYPE_ETH = 2;
   */
  ETH = 2,

  /**
   * The wallet type for Sepolia testnet ETH and its ERC-20 tokens
   *
   * @generated from enum value: WALLET_TYPE_ETH_SEPOLIA = 3;
   */
  ETH_SEPOLIA = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(WalletType)
proto3.util.setEnumType(WalletType, "fusionchain.treasury.WalletType", [
  { no: 0, name: "WALLET_TYPE_UNSPECIFIED" },
  { no: 1, name: "WALLET_TYPE_QRDO" },
  { no: 2, name: "WALLET_TYPE_ETH" },
  { no: 3, name: "WALLET_TYPE_ETH_SEPOLIA" },
]);

/**
 * WalletRequestType used at the request level for query_keys
 *
 * @generated from enum fusionchain.treasury.WalletRequestType
 */
export enum WalletRequestType {
  /**
   * The wallet type is missing
   *
   * @generated from enum value: WALLET_REQUEST_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * The wallet type for all wallets to be derived
   *
   * @generated from enum value: WALLET_REQUEST_TYPE_ALL = 1;
   */
  ALL = 1,

  /**
   * The wallet type for native Fusion chain cosmos accounts (not ERC-20 QRDO
   * tokens)
   *
   * @generated from enum value: WALLET_REQUEST_TYPE_QRDO = 2;
   */
  QRDO = 2,

  /**
   * The wallet type for mainnet ETH and its ERC-20 tokens (including non-native
   * QRDO)
   *
   * @generated from enum value: WALLET_REQUEST_TYPE_ETH = 3;
   */
  ETH = 3,

  /**
   * The wallet type for Sepolia testnet ETH and its ERC-20 tokens
   *
   * @generated from enum value: WALLET_REQUEST_TYPE_ETH_SEPOLIA = 4;
   */
  ETH_SEPOLIA = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(WalletRequestType)
proto3.util.setEnumType(WalletRequestType, "fusionchain.treasury.WalletRequestType", [
  { no: 0, name: "WALLET_REQUEST_TYPE_UNSPECIFIED" },
  { no: 1, name: "WALLET_REQUEST_TYPE_ALL" },
  { no: 2, name: "WALLET_REQUEST_TYPE_QRDO" },
  { no: 3, name: "WALLET_REQUEST_TYPE_ETH" },
  { no: 4, name: "WALLET_REQUEST_TYPE_ETH_SEPOLIA" },
]);

/**
 * @generated from message fusionchain.treasury.Wallet
 */
export class Wallet extends Message<Wallet> {
  /**
   * @generated from field: uint64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: fusionchain.treasury.WalletType type = 2;
   */
  type = WalletType.UNSPECIFIED;

  /**
   * @generated from field: uint64 key_id = 3;
   */
  keyId = protoInt64.zero;

  constructor(data?: PartialMessage<Wallet>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.treasury.Wallet";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(WalletType) },
    { no: 3, name: "key_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Wallet {
    return new Wallet().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Wallet {
    return new Wallet().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Wallet {
    return new Wallet().fromJsonString(jsonString, options);
  }

  static equals(a: Wallet | PlainMessage<Wallet> | undefined, b: Wallet | PlainMessage<Wallet> | undefined): boolean {
    return proto3.util.equals(Wallet, a, b);
  }
}

