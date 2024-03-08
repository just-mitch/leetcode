import { describe, expect, test } from "bun:test";
import { letterCombinations } from "./lib";

describe("letterCombinations", () => {
  test("Simple cases", () => {
    expect(letterCombinations("23")).toEqual([
      "ad",
      "ae",
      "af",
      "bd",
      "be",
      "bf",
      "cd",
      "ce",
      "cf",
    ]);
  });
});
