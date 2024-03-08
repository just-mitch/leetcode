import { describe, expect, test } from "bun:test";
import { invertTree } from "./lib";

describe("invert binary tree", () => {
  test("Simple cases", () => {
    expect(invertTree(null)).toEqual(null);
    expect(invertTree({ val: 1, left: null, right: null })).toEqual({
      val: 1,
      left: null,
      right: null,
    });
    const tree = {
      val: 4,
      left: {
        val: 2,
        left: { val: 1, left: null, right: null },
        right: { val: 3, left: null, right: null },
      },
      right: {
        val: 7,
        left: { val: 6, left: null, right: null },
        right: { val: 9, left: null, right: null },
      },
    };

    const invertedTree = {
      val: 4,
      left: {
        val: 7,
        left: { val: 9, left: null, right: null },
        right: { val: 6, left: null, right: null },
      },
      right: {
        val: 2,
        left: { val: 3, left: null, right: null },
        right: { val: 1, left: null, right: null },
      },
    };
    expect(invertTree(tree)).toEqual(invertedTree);
  });
});
