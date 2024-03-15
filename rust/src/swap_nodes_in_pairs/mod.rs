use crate::data_structures::list_node::ListNode;
struct Solution;
impl Solution {
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head;
        if head.is_none() {
            return None;
        } else if head.as_ref().unwrap().next.is_none() {
            return head;
        }

        let mut next_pair_head = head.as_mut().unwrap().next.take();

        head.as_mut().unwrap().next =
            Solution::swap_pairs(next_pair_head.as_mut().unwrap().next.take());

        next_pair_head.as_mut().unwrap().next = head;

        next_pair_head
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_1() {
        let head = ListNode::from_vec(vec![1, 2, 3, 4]);

        let expected = ListNode::from_vec(vec![2, 1, 4, 3]);

        assert_eq!(Solution::swap_pairs(head), expected);
    }
}
