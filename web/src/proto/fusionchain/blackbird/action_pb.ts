// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file fusionchain/blackbird/action.proto (package fusionchain.policy, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * Action wraps a message that needs to be approved by a set of approvers.
 *
 * @generated from message fusionchain.policy.Action
 */
export class Action extends Message<Action> {
  /**
   * @generated from field: uint64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: repeated string approvers = 2;
   */
  approvers: string[] = [];

  /**
   * @generated from field: bool completed = 3;
   */
  completed = false;

  /**
   * Optional policy id that must be satisfied by the approvers.
   * If not specified, it's up to the creator of the action to decide what to
   * apply.
   *
   * @generated from field: uint64 policy_id = 4;
   */
  policyId = protoInt64.zero;

  /**
   * Original message that started the action, it will be executed when the
   * policy is satisfied.
   *
   * @generated from field: google.protobuf.Any msg = 5;
   */
  msg?: Any;

  constructor(data?: PartialMessage<Action>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fusionchain.policy.Action";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "approvers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "completed", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "policy_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 5, name: "msg", kind: "message", T: Any },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Action {
    return new Action().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Action {
    return new Action().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Action {
    return new Action().fromJsonString(jsonString, options);
  }

  static equals(a: Action | PlainMessage<Action> | undefined, b: Action | PlainMessage<Action> | undefined): boolean {
    return proto3.util.equals(Action, a, b);
  }
}

