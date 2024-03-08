export function threeSumClosest(nums: number[], target: number): number {
  if (nums.length < 3) {
    return Math.max();
  }

  nums.sort((a, b) => a - b);

  let closestSum = nums[0] + nums[1] + nums[2];

  for (let i = 0; i < nums.length - 2; i += 1) {
    let left = i + 1;
    let right = nums.length - 1;

    while (left < right) {
      const sum = nums[i] + nums[left] + nums[right];

      if (Math.abs(target - sum) < Math.abs(target - closestSum)) {
        closestSum = sum;
      }
      if (sum === target) {
        return sum;
      }

      if (sum < target) {
        left += 1;
      } else {
        right -= 1;
      }
    }
  }

  return closestSum;
}
