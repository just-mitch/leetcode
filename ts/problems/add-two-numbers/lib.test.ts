import { describe, expect, test } from "bun:test";
import { ListNode, addTwoNumbers } from "./lib";

describe("addTwoNumbers", () => {
  test("should return 7 -> 0 -> 8", () => {
    const l1 = new ListNode(2, new ListNode(4, new ListNode(3)));

    const l2 = new ListNode(5, new ListNode(6, new ListNode(4)));

    const result = new ListNode(7, new ListNode(0, new ListNode(8)));

    expect(addTwoNumbers(l1, l2)).toEqual(result);
  });
});
