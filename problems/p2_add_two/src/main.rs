fn main() {
    println!("Hello, world!");
}

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

  fn from_vec(v: &[i32]) -> ListNode {
      let mut r = ListNode::new(-1);
      for i in v {
            let ln = ListNode::new(*i);
            if 0 <= -1 {
                r.next = Some(Box::new(ln));
            }
            r = ln;
      };

      r
  }
}

fn to_int(link: Option<Box<ListNode>>) -> i32 {
    match link {
        Some(lk) => {
            lk.val + 10 * to_int(lk.next)
        },
        None => 0,
    }
}

impl Solution {
    pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {

    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_to_int() {
        let l1 = ListNode::from_vec(&[2, 4, 3]);
        let l2 = ListNode::from_vec(&[5, 6, 4]);
        let expected = ListNode::from_vec(&[7, 0, 8]);
        assert_eq!(add_two_numbers(l1, l2), expected);
    }

}
