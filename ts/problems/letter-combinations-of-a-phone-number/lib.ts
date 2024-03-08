function digitToChars(digit: string): string[] {
  switch (digit) {
    case "2":
      return ["a", "b", "c"];
    case "3":
      return ["d", "e", "f"];
    case "4":
      return ["g", "h", "i"];
    case "5":
      return ["j", "k", "l"];
    case "6":
      return ["m", "n", "o"];
    case "7":
      return ["p", "q", "r", "s"];
    case "8":
      return ["t", "u", "v"];
    case "9":
      return ["w", "x", "y", "z"];
    default:
      throw new Error("digit is not in [2,9]");
  }
}

export function letterCombinations(digits: string): string[] {
  const arr = digits.split("");
  if (arr.length === 0) {
    return [];
  }

  const chars = arr.map(digitToChars);
  const head = chars.shift();

  return chars.reduce((acc, curr) => {
    const newAcc = [];
    for (let i = 0; i < acc.length; i++) {
      for (let j = 0; j < curr.length; j++) {
        newAcc.push(acc[i] + curr[j]);
      }
    }
    return newAcc;
  }, head || []);
}
