struct Solution;

impl Solution {
  pub fn unique_paths(m: i32, n: i32) -> i32 {
    let mut dp = vec![vec![0; n as usize]; m as usize];
    for i in 0..m as usize {
      for j in 0..n as usize {
        if i == 0 || j == 0 {
          dp[i][j] = 1;
        } else {
          dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
        }
      }
    }
    dp[m as usize - 1][n as usize - 1]
  }
}


#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_62() {
    assert_eq!(Solution::unique_paths(3, 2), 3);
    assert_eq!(Solution::unique_paths(7, 3), 28);
  }
}