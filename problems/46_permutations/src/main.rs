// run a single example
fn main() {
    let nums = vec![1, 2, 3];
    let p = permute(nums.clone());
    println!("permute {:?} to:\n[", nums);
    for perm in p.iter() {
        println!("\t{:?}", perm);
    }
    println!("]");
}

// permute create all of the permutations of an vector of numbers
// Trying a mini-impprovmeent for cases of length 2 or fewer
// otherwise use dfs to build all permutations
// This would work with DFS only
fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let p = match nums.len() {
        0 => Vec::new(),
        1 => vec![nums],
        2 => {
            let mut v: Vec<Vec<i32>> = Vec::with_capacity(2);
            let mut n = nums.clone();
            v.push(nums);
            n.reverse();
            v.push(n);
            v
        }
        _ => dfs(nums),
    };

    p
}

// depth first search (dfs)
// recursively dive into the vector
// building each permutation
fn dfs(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut vecs: Vec<Vec<i32>> = Vec::new();

    match nums.len() {
        0 | 1 => vecs.push(nums),
        _ => {
            for i in 0..nums.len() {
                let mut arr = nums.clone(); // all of the unvisted nodes
                let v = arr.remove(i); // pull out current node
                let children = dfs(arr); // permutate all unvisited nodes
                // For each permutation returned, add our value to the front, and
                // extend our answer
                for c in children.iter() {
                    let mut p: Vec<i32> = Vec::with_capacity(nums.len());
                    p.push(v); // add the current value first each time
                    // unwrap the permutation of the children nodes
                    for z in c.iter() {
                        p.push(*z);
                    }

                    //save answer
                    vecs.push(p);
                }
            }
        }
    }

    vecs
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_permute() {
        struct TestCase {
            nums: Vec<i32>,
            expected: Vec<Vec<i32>>,
        }

        let test_cases = vec![
            TestCase {
                nums: vec![1],
                expected: vec![vec![1]],
            },
            TestCase {
                nums: vec![0, 1],
                expected: vec![vec![0, 1], vec![1, 0]],
            },
            TestCase {
                nums: vec![1, 2, 3],
                expected: vec![
                    vec![1, 2, 3],
                    vec![1, 3, 2],
                    vec![2, 1, 3],
                    vec![2, 3, 1],
                    vec![3, 1, 2],
                    vec![3, 2, 1],
                ],
            },
            TestCase {
                nums: vec![1, 2, 3, 4],
                expected: vec![
                    vec![1, 2, 3, 4],
                    vec![1, 2, 4, 3],
                    vec![1, 3, 2, 4],
                    vec![1, 3, 4, 2],
                    vec![1, 4, 2, 3],
                    vec![1, 4, 3, 2],
                    vec![2, 1, 3, 4],
                    vec![2, 1, 4, 3],
                    vec![2, 3, 1, 4],
                    vec![2, 3, 4, 1],
                    vec![2, 4, 1, 3],
                    vec![2, 4, 3, 1],
                    vec![3, 1, 2, 4],
                    vec![3, 1, 4, 2],
                    vec![3, 2, 1, 4],
                    vec![3, 2, 4, 1],
                    vec![3, 4, 1, 2],
                    vec![3, 4, 2, 1],
                    vec![4, 1, 2, 3],
                    vec![4, 1, 3, 2],
                    vec![4, 2, 1, 3],
                    vec![4, 2, 3, 1],
                    vec![4, 3, 1, 2],
                    vec![4, 3, 2, 1],
                ],
            },
        ];

        for test_case in test_cases {
            let mut actual = permute(test_case.nums);
            actual.sort();

            assert_eq!(actual, test_case.expected);
        }
    }
}
