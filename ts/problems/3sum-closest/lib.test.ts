import { describe, expect, test } from "bun:test";
import { threeSumClosest } from "./lib";

describe("3sum-closest", () => {
  test("Simple cases", () => {
    expect(threeSumClosest([-1, 2, 1, -4], 1)).toEqual(2);
    expect(threeSumClosest([0, 0, 0], 1)).toEqual(0);
    expect(threeSumClosest([1, 1, 1, 0], -100)).toEqual(2);
  });
});
