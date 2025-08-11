//Add one to some integer represented as a vector
fn plus_one(digits: Vec<i32>) -> Vec<i32> {
    let mut v = digits.clone();
    let mut carry = 1;
    for n in (0..digits.len()).rev() {
        let mut val = digits[n] + carry;
        if val > 9 {
            val = val - 10;
            v[n] = val;
        } else {
            v[n] = val;
            carry = 0;
            break;
        }
    }

    if carry > 0 {
        v.insert(0, 1);
    }

   v
}

fn main() {
    let i = vec![1,2,3];
    let v = plus_one(i);
    println!("Hello, world!");
    println!("{:?}", v);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_plus_one() {
        struct TestCase {
            input: Vec<i32>,
            output: Vec<i32>,
        }

        let test_cases = vec![
            TestCase {
                input: vec![0],
                output: vec![1],
            },
            TestCase {
                input: vec![1,2,3],
                output: vec![1,2,4],
            },
            TestCase {
                input: vec![9],
                output: vec![1, 0],
            },
            TestCase {
                input: vec![1, 0, 0],
                output: vec![1, 0, 1],
            },
            TestCase {
                input: vec![9, 9, 9],
                output: vec![1, 0, 0, 0],
            },
            TestCase {
                input: vec![1, 9, 9, 9],
                output: vec![2, 0, 0, 0],
            },
            TestCase {
                input: vec![ 1, 9, 9, 9, 9, 9, 9, 9, 9],
                output: vec![2, 0, 0, 0, 0, 0, 0, 0, 0],
            },
        ];

        for test_case in test_cases {
            let result = plus_one(test_case.input);
            assert_eq!(result, test_case.output);
        }
    }
}
