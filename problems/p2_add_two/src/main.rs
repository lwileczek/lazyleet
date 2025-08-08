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
        let mut foo = Vec::from(v);
        foo.reverse();
        let mut l = ListNode::new(foo[0]);
        for i in 1..foo.len() {
            let n = ListNode{
                val: foo[i],
                next: Some(Box::new(l)),
            };
            l = n;
        };

        return l
  }
}

fn add_nodes(l1: &Option<Box<ListNode>>, l2: &Option<Box<ListNode>>) -> i32 {
    let v0 = match l1 {
        Some(v) => v.val,
        None => 0,
    };
    let v1 = match l2 {
        Some(v) => v.val,
        None => 0,
    };

    return v0 + v1;
}

fn add_with_carry(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>, c: i32) -> Option<Box<ListNode>> {
    if l1.is_none() && l2.is_none() {
        if c == 0 {
            return None;
        }
        return Some(Box::new(ListNode::new(c)));
    }

    let mut sum = add_nodes(&l1, &l2) + c;
    let mut carry = 0;
    if sum > 9 {
        carry = 1;
        sum = sum -10;
    }

    let node = Some(Box::new(ListNode{
        val: sum,
        next: add_with_carry(l1.and_then(|v| v.next), l2.and_then(|v| v.next), carry),
    }));
    
    return node;
}

//impl Solution {
    pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        return add_with_carry(l1, l2, 0);
    }
//}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add_ex0() {
        let l1 = Some(Box::new(ListNode::from_vec(&[2, 4, 3])));
        let l2 = Some(Box::new(ListNode::from_vec(&[5, 6, 4])));
        let expected = Some(Box::new(ListNode::from_vec(&[7, 0, 8])));
        assert_eq!(add_two_numbers(l1, l2), expected);
    }
}
