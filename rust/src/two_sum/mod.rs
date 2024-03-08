struct Solution;


impl Solution {
  pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {

    let mut map = std::collections::HashMap::new();

    for (i, num) in nums.iter().enumerate() {
      let complement = target - num;
      if map.contains_key(&complement) {
        return vec![*map.get(&complement).unwrap(), i as i32];
      }
      map.insert(num, i as i32);
    }

    vec![]
      
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_1() {
    let nums = vec![2, 7, 11, 15];
    assert_eq!(Solution::two_sum(nums, 9), vec![0, 1]);
  }

  #[test]
  fn test_2() {
    assert_eq!(Solution::two_sum(vec![3, 2, 4], 6), vec![1, 2]);
  }

  #[test]
  fn test_3() {
    assert_eq!(Solution::two_sum(vec![3, 3], 6), vec![0, 1]);
  }

  #[test]
  fn test_4() {
    assert_eq!(Solution::two_sum(vec![3, 2, 3], 6), vec![0, 2]);
  }
}