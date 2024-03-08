#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }

  pub fn append(&mut self, val: i32) {
    let mut current = self;
    while current.next.is_some() {
      current = current.next.as_mut().unwrap()
    }
    current.next = Some(Box::new(ListNode::new(val)));
  }

  pub fn from_vec(vec: Vec<i32>) -> Option<Box<ListNode>> {
    let mut head = Some(Box::new(ListNode::new(0)));
    let mut current = &mut head;
    for val in vec {
      current.as_mut().unwrap().append(val);
      current = &mut current.as_mut().unwrap().next;
    }
    head.unwrap().next
  }

}