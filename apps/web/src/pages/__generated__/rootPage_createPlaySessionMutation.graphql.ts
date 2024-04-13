/**
 * @generated SignedSource<<ff04b5a7a3dedd87b5e507cf809c48ea>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */

/* eslint-disable */
// @ts-nocheck
import { ConcreteRequest, Mutation } from "relay-runtime"

export type rootPage_createPlaySessionMutation$variables = {
  videoID: string
}
export type rootPage_createPlaySessionMutation$data = {
  readonly createPlaySession: {
    readonly id: string
  }
}
export type rootPage_createPlaySessionMutation = {
  response: rootPage_createPlaySessionMutation$data
  variables: rootPage_createPlaySessionMutation$variables
}

const node: ConcreteRequest = (function () {
  var v0 = [
      {
        defaultValue: null,
        kind: "LocalArgument",
        name: "videoID",
      },
    ],
    v1 = [
      {
        alias: null,
        args: [
          {
            fields: [
              {
                kind: "Variable",
                name: "videoID",
                variableName: "videoID",
              },
            ],
            kind: "ObjectValue",
            name: "input",
          },
        ],
        concreteType: "PlaySession",
        kind: "LinkedField",
        name: "createPlaySession",
        plural: false,
        selections: [
          {
            alias: null,
            args: null,
            kind: "ScalarField",
            name: "id",
            storageKey: null,
          },
        ],
        storageKey: null,
      },
    ]
  return {
    fragment: {
      argumentDefinitions: v0 /*: any*/,
      kind: "Fragment",
      metadata: null,
      name: "rootPage_createPlaySessionMutation",
      selections: v1 /*: any*/,
      type: "Mutation",
      abstractKey: null,
    },
    kind: "Request",
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: "Operation",
      name: "rootPage_createPlaySessionMutation",
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: "f0fc7e038032ef3a2c2b6b1893c1971c",
      id: null,
      metadata: {},
      name: "rootPage_createPlaySessionMutation",
      operationKind: "mutation",
      text: "mutation rootPage_createPlaySessionMutation(\n  $videoID: ID!\n) {\n  createPlaySession(input: {videoID: $videoID}) {\n    id\n  }\n}\n",
    },
  }
})()

;(node as any).hash = "15649eb0c2e4f24934815a772755f963"

export default node
