struct Solution;

impl Solution {
  pub fn next_permutation(nums: &mut Vec<i32>) {
    for i in (0..nums.len() - 1).rev() {
      if nums[i] < nums[i + 1] {
        let mut j = nums.len() - 1;
        while nums[i] >= nums[j] {
          j -= 1;
        }
        nums.swap(i, j);
        nums[i + 1..].reverse();
        return;
      }
    }
    nums.reverse();
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_1() {
    let mut nums = vec![1, 2, 3];
    Solution::next_permutation(&mut nums);
    assert_eq!(nums, vec![1, 3, 2]);
  }

  #[test]
  fn test_2() {
    let mut nums = vec![3, 2, 1];
    Solution::next_permutation(&mut nums);
    assert_eq!(nums, vec![1, 2, 3]);
  }

  #[test]
  fn test_3() {
    let mut nums = vec![1, 1, 5];
    Solution::next_permutation(&mut nums);
    assert_eq!(nums, vec![1, 5, 1]);
  }

  #[test]
  fn test_4() {
    let mut nums = vec![1];
    Solution::next_permutation(&mut nums);
    assert_eq!(nums, vec![1]);
  }
}