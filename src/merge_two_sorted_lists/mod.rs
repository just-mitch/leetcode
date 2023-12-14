// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }

  fn append(&mut self, val: i32) {
    let mut current = self;
    while current.next.is_some() {
      current = current.next.as_mut().unwrap()
    }
    current.next = Some(Box::new(ListNode::new(val)));
  }

  fn from_vec(vec: Vec<i32>) -> Option<Box<ListNode>> {
    let mut head = Some(Box::new(ListNode::new(0)));
    let mut current = &mut head;
    for val in vec {
      current.as_mut().unwrap().append(val);
      current = &mut current.as_mut().unwrap().next;
    }
    head.unwrap().next
  }

}

pub struct Solution {}

impl Solution {
  pub fn merge_two_lists(list1: Option<Box<ListNode>>, list2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut list1 = list1;
    let mut list2 = list2;
    let mut head = Some(Box::new(ListNode::new(0)));
    let mut current = &mut head;
    while list1.is_some() && list2.is_some() {
      let val1 = list1.as_ref().unwrap().val;
      let val2 = list2.as_ref().unwrap().val;
      if val1 < val2 {
        let next = list1.as_mut().unwrap().next.take();
        current.as_mut().unwrap().next = list1.take();
        list1 = next;
      } else {
        let next = list2.as_mut().unwrap().next.take();
        current.as_mut().unwrap().next = list2.take();
        list2 = next;
      }
      current = &mut current.as_mut().unwrap().next;
    }
    if list1.is_some() {
      current.as_mut().unwrap().next = list1.take();
    }
    if list2.is_some() {
      current.as_mut().unwrap().next = list2.take();
    }
    head.unwrap().next

      
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_1() {
    let list1 = ListNode::from_vec(vec![1, 2, 4]);

    let list2 = ListNode::from_vec(vec![1, 3, 4]);

    let expected = ListNode::from_vec(vec![1, 1, 2, 3, 4, 4]);

    assert_eq!(Solution::merge_two_lists(list1, list2), expected);
  }
}