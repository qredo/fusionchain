// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file fusionchain/policy/policy.proto (package fusionchain.policy, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message fusionchain.policy.Policy
 */
export class Policy extends Message<Policy> {
  /**
   * @generated from field: uint64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * The actual policy informations. It must be one the supported policy types:
   * - BlackbirdPolicy
   *
   * @generated from field: google.protobuf.Any policy = 3;
   */
  policy?: Any;

  constructor(data?: PartialMessage<Policy>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.policy.Policy";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "policy", kind: "message", T: Any },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Policy {
    return new Policy().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Policy {
    return new Policy().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Policy {
    return new Policy().fromJsonString(jsonString, options);
  }

  static equals(a: Policy | PlainMessage<Policy> | undefined, b: Policy | PlainMessage<Policy> | undefined): boolean {
    return proto3.util.equals(Policy, a, b);
  }
}

/**
 * @generated from message fusionchain.policy.BlackbirdPolicy
 */
export class BlackbirdPolicy extends Message<BlackbirdPolicy> {
  /**
   * @generated from field: bytes data = 1;
   */
  data = new Uint8Array(0);

  /**
   * @generated from field: repeated fusionchain.policy.BlackbirdPolicyParticipant participants = 2;
   */
  participants: BlackbirdPolicyParticipant[] = [];

  constructor(data?: PartialMessage<BlackbirdPolicy>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.policy.BlackbirdPolicy";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "participants", kind: "message", T: BlackbirdPolicyParticipant, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BlackbirdPolicy {
    return new BlackbirdPolicy().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BlackbirdPolicy {
    return new BlackbirdPolicy().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BlackbirdPolicy {
    return new BlackbirdPolicy().fromJsonString(jsonString, options);
  }

  static equals(a: BlackbirdPolicy | PlainMessage<BlackbirdPolicy> | undefined, b: BlackbirdPolicy | PlainMessage<BlackbirdPolicy> | undefined): boolean {
    return proto3.util.equals(BlackbirdPolicy, a, b);
  }
}

/**
 * @generated from message fusionchain.policy.BlackbirdPolicyParticipant
 */
export class BlackbirdPolicyParticipant extends Message<BlackbirdPolicyParticipant> {
  /**
   * @generated from field: string abbreviation = 1;
   */
  abbreviation = "";

  /**
   * @generated from field: string address = 2;
   */
  address = "";

  constructor(data?: PartialMessage<BlackbirdPolicyParticipant>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.policy.BlackbirdPolicyParticipant";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "abbreviation", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BlackbirdPolicyParticipant {
    return new BlackbirdPolicyParticipant().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BlackbirdPolicyParticipant {
    return new BlackbirdPolicyParticipant().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BlackbirdPolicyParticipant {
    return new BlackbirdPolicyParticipant().fromJsonString(jsonString, options);
  }

  static equals(a: BlackbirdPolicyParticipant | PlainMessage<BlackbirdPolicyParticipant> | undefined, b: BlackbirdPolicyParticipant | PlainMessage<BlackbirdPolicyParticipant> | undefined): boolean {
    return proto3.util.equals(BlackbirdPolicyParticipant, a, b);
  }
}

/**
 * @generated from message fusionchain.policy.BlackbirdPolicyPayload
 */
export class BlackbirdPolicyPayload extends Message<BlackbirdPolicyPayload> {
  /**
   * @generated from field: bytes witness = 1;
   */
  witness = new Uint8Array(0);

  constructor(data?: PartialMessage<BlackbirdPolicyPayload>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.policy.BlackbirdPolicyPayload";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "witness", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BlackbirdPolicyPayload {
    return new BlackbirdPolicyPayload().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BlackbirdPolicyPayload {
    return new BlackbirdPolicyPayload().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BlackbirdPolicyPayload {
    return new BlackbirdPolicyPayload().fromJsonString(jsonString, options);
  }

  static equals(a: BlackbirdPolicyPayload | PlainMessage<BlackbirdPolicyPayload> | undefined, b: BlackbirdPolicyPayload | PlainMessage<BlackbirdPolicyPayload> | undefined): boolean {
    return proto3.util.equals(BlackbirdPolicyPayload, a, b);
  }
}

/**
 * @generated from message fusionchain.policy.BlackbirdPolicyMetadata
 */
export class BlackbirdPolicyMetadata extends Message<BlackbirdPolicyMetadata> {
  /**
   * The "decompiled" version of the policy, in a readable format.
   *
   * @generated from field: string pretty = 1;
   */
  pretty = "";

  constructor(data?: PartialMessage<BlackbirdPolicyMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.policy.BlackbirdPolicyMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pretty", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BlackbirdPolicyMetadata {
    return new BlackbirdPolicyMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BlackbirdPolicyMetadata {
    return new BlackbirdPolicyMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BlackbirdPolicyMetadata {
    return new BlackbirdPolicyMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: BlackbirdPolicyMetadata | PlainMessage<BlackbirdPolicyMetadata> | undefined, b: BlackbirdPolicyMetadata | PlainMessage<BlackbirdPolicyMetadata> | undefined): boolean {
    return proto3.util.equals(BlackbirdPolicyMetadata, a, b);
  }
}

