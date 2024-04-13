/**
 * @generated SignedSource<<b5255a02a1ff391fc622531405f4b773>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */

/* eslint-disable */
// @ts-nocheck
import { ConcreteRequest, Query } from "relay-runtime"

export type rootPageQuery$variables = Record<PropertyKey, never>
export type rootPageQuery$data = {
  readonly videos: ReadonlyArray<{
    readonly id: string
    readonly path: string
  }>
}
export type rootPageQuery = {
  response: rootPageQuery$data
  variables: rootPageQuery$variables
}

const node: ConcreteRequest = (function () {
  var v0 = [
    {
      alias: null,
      args: null,
      concreteType: "Video",
      kind: "LinkedField",
      name: "videos",
      plural: true,
      selections: [
        {
          alias: null,
          args: null,
          kind: "ScalarField",
          name: "id",
          storageKey: null,
        },
        {
          alias: null,
          args: null,
          kind: "ScalarField",
          name: "path",
          storageKey: null,
        },
      ],
      storageKey: null,
    },
  ]
  return {
    fragment: {
      argumentDefinitions: [],
      kind: "Fragment",
      metadata: null,
      name: "rootPageQuery",
      selections: v0 /*: any*/,
      type: "Query",
      abstractKey: null,
    },
    kind: "Request",
    operation: {
      argumentDefinitions: [],
      kind: "Operation",
      name: "rootPageQuery",
      selections: v0 /*: any*/,
    },
    params: {
      cacheID: "80fef6e48f7a0a6857b00cdef689605d",
      id: null,
      metadata: {},
      name: "rootPageQuery",
      operationKind: "query",
      text: "query rootPageQuery {\n  videos {\n    id\n    path\n  }\n}\n",
    },
  }
})()

;(node as any).hash = "3691851af2f342fef2bdd73c99468a20"

export default node
