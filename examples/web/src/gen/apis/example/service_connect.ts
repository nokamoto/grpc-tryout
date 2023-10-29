// @generated by protoc-gen-connect-es v1.1.3 with parameter "target=ts"
// @generated from file apis/example/service.proto (package example, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetShelfRequest } from "./service_pb.js";
import { Shelf } from "./resource_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service example.Library
 */
export const Library = {
  typeName: "example.Library",
  methods: {
    /**
     * @generated from rpc example.Library.GetShelf
     */
    getShelf: {
      name: "GetShelf",
      I: GetShelfRequest,
      O: Shelf,
      kind: MethodKind.Unary,
    },
  }
} as const;

