fn main() {
    let data = "PAYPALISHIRING".to_string();
    let v = convert(data, 3);
    println!("Converted string (3): {}", v);
    let ex = "PAHNAPLSIIGYIR".to_string();
    println!("Converted string (3): {}", ex);
}

pub fn convert(s: String, num_rows: i32) -> String {
    if num_rows <= 1 {
        return s;
    };

    let mut b: Vec<Vec<char>> = Vec::new();
    for _ in 0..num_rows {
        b.push(Vec::new());
    }

    let m = num_rows as usize - 1;
    let mut pos = 0;
    let mut delta: i32 = 1;
    for c in s.chars() {
        b[pos].push(c);
        pos = ((pos as i32) + delta) as usize;
        if m <= pos || pos <= 0 {
            delta *= -1;
        }
    }

    let zigzag = b.iter().flatten().collect::<String>();

    return zigzag;
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_convert() {
        struct TestCase {
            input: String,
            output: String,
            rows: i32,
        }

        let test_cases = vec![
            TestCase {
                input: "PPAY".to_string(),
                output: "PPAY".to_string(),
                rows: 1,
            },
            TestCase {
                input:  "PAYPALISHIRING".to_string(),
                output: "PAHNAPLSIIGYIR".to_string(),
                rows: 3,
            },
            TestCase {
                input:  "PAYPALISHIRING".to_string(),
                output: "PINALSIGYAHRPI".to_string(),
                rows: 4,
            },
        ];

        for test_case in test_cases {
            let result = convert(test_case.input, test_case.rows);
            assert_eq!(result, test_case.output);
        }
    }
}
