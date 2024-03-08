pub struct Solution {}

impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut max_reach = 0;
        for (i, &num) in nums.iter().enumerate() {
            if i > max_reach {
                return false;
            }
            max_reach = max_reach.max(i + num as usize);
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_1() {
        let nums = vec![2, 3, 1, 1, 4];

        assert_eq!(Solution::can_jump(nums), true);
    }

    #[test]
    fn test_2() {
        let nums = vec![3, 2, 1, 0, 4];

        assert_eq!(Solution::can_jump(nums), false);
    }
}