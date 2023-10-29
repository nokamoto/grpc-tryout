// @generated by protoc-gen-es v1.4.1 with parameter "target=ts"
// @generated from file apis/tryout/tryout.proto (package tryout, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message tryout.Method
 */
export class Method extends Message<Method> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: repeated string fields = 2;
   */
  fields: string[] = [];

  constructor(data?: PartialMessage<Method>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tryout.Method";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "fields", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Method {
    return new Method().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Method {
    return new Method().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Method {
    return new Method().fromJsonString(jsonString, options);
  }

  static equals(a: Method | PlainMessage<Method> | undefined, b: Method | PlainMessage<Method> | undefined): boolean {
    return proto3.util.equals(Method, a, b);
  }
}

/**
 * @generated from message tryout.Service
 */
export class Service extends Message<Service> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: repeated tryout.Method methods = 2;
   */
  methods: Method[] = [];

  constructor(data?: PartialMessage<Service>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tryout.Service";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "methods", kind: "message", T: Method, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Service {
    return new Service().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Service {
    return new Service().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Service {
    return new Service().fromJsonString(jsonString, options);
  }

  static equals(a: Service | PlainMessage<Service> | undefined, b: Service | PlainMessage<Service> | undefined): boolean {
    return proto3.util.equals(Service, a, b);
  }
}

/**
 * @generated from message tryout.Proto
 */
export class Proto extends Message<Proto> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: repeated tryout.Service services = 2;
   */
  services: Service[] = [];

  constructor(data?: PartialMessage<Proto>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tryout.Proto";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "services", kind: "message", T: Service, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Proto {
    return new Proto().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Proto {
    return new Proto().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Proto {
    return new Proto().fromJsonString(jsonString, options);
  }

  static equals(a: Proto | PlainMessage<Proto> | undefined, b: Proto | PlainMessage<Proto> | undefined): boolean {
    return proto3.util.equals(Proto, a, b);
  }
}

