// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file fusionchain/identity/keyring.proto (package fusionchain.identity, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message fusionchain.identity.Keyring
 */
export class Keyring extends Message<Keyring> {
  /**
   * @generated from field: string address = 1;
   */
  address = "";

  /**
   * @generated from field: string creator = 2;
   */
  creator = "";

  /**
   * @generated from field: string description = 3;
   */
  description = "";

  /**
   * @generated from field: repeated string admins = 4;
   */
  admins: string[] = [];

  /**
   * @generated from field: repeated string parties = 5;
   */
  parties: string[] = [];

  /**
   * @generated from field: uint64 admin_policy_id = 6;
   */
  adminPolicyId = protoInt64.zero;

  /**
   * @generated from field: fusionchain.identity.KeyringFees fees = 7;
   */
  fees?: KeyringFees;

  /**
   * @generated from field: bool is_active = 8;
   */
  isActive = false;

  constructor(data?: PartialMessage<Keyring>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.identity.Keyring";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "creator", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "admins", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "parties", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "admin_policy_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 7, name: "fees", kind: "message", T: KeyringFees },
    { no: 8, name: "is_active", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Keyring {
    return new Keyring().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Keyring {
    return new Keyring().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Keyring {
    return new Keyring().fromJsonString(jsonString, options);
  }

  static equals(a: Keyring | PlainMessage<Keyring> | undefined, b: Keyring | PlainMessage<Keyring> | undefined): boolean {
    return proto3.util.equals(Keyring, a, b);
  }
}

/**
 * @generated from message fusionchain.identity.KeyringFees
 */
export class KeyringFees extends Message<KeyringFees> {
  /**
   * @generated from field: uint64 key_req = 1;
   */
  keyReq = protoInt64.zero;

  /**
   * @generated from field: uint64 sig_req = 2;
   */
  sigReq = protoInt64.zero;

  constructor(data?: PartialMessage<KeyringFees>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.identity.KeyringFees";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key_req", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "sig_req", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KeyringFees {
    return new KeyringFees().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KeyringFees {
    return new KeyringFees().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KeyringFees {
    return new KeyringFees().fromJsonString(jsonString, options);
  }

  static equals(a: KeyringFees | PlainMessage<KeyringFees> | undefined, b: KeyringFees | PlainMessage<KeyringFees> | undefined): boolean {
    return proto3.util.equals(KeyringFees, a, b);
  }
}

