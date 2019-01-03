fn main() {
    let cases = Case::test_cases();
    for case in cases {
        let re = longest_palindrome(case.s);
        println!("\nin: {}\t\tout: {}\t\tok: {}", case.s, re, re == case.o);
    }
}

fn longest_palindrome(s: &str) -> &str {
    if s.len() == 0 {
        return "";
    }

    let mut start: usize = 0;
    let mut end: usize = 0;
    let bs = s.as_bytes();
    for i in 0..bs.len() {
        let l1 = expand_around_center(bs, i, i);
        let l2 = expand_around_center(bs, i, i + 1);
        let ll = l1.max(l2);
        if ll > (end - start) {
            start = i - (ll - 1) / 2;
            end = i + ll / 2;
        }
    }
    return &s[start..=end];
}
fn expand_around_center(s: &[u8], mut l: usize, mut r: usize) -> usize {
    while l > 0 && r < s.len() && s[l] == s[r] {
        l -= 1;
        r += 1;
    }
    if l == 0 && r < s.len() && s[l] == s[r] {
        return r - l + 1;
    }
    return r - l - 1;
}

fn longest_palindrome2(s: &str) -> &str {
    if s.len() == 0 {
        return "";
    }
    let mut i1 = 0;
    let mut i2 = s.len() - 1;
    let mut start = 0;
    let mut end = 0;
    let bs = s.as_bytes();
    while i1 < i2 {
        if is_palindrome(&bs[i1..=i2]) {
            if (i2 - i1) > (end - start) {
                start = i1;
                end = i2;
            }
        }
        i2 -= 1;
        if i2 == i1 {
            i1 += 1;
            i2 = s.len() - 1;
            if (i2 - i1) < (end - start) {
                break;
            }
        }
    }
    return &s[start..=end];
}

fn is_palindrome(s: &[u8]) -> bool {
    let mut i = 0;
    let mut j = s.len() - 1;
    while i < j {
        if s[i] != s[j] {
            return false;
        }
        i += 1;
        j -= 1;
    }
    return true;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_longest_palindrome() {
        let cases = Case::test_cases();
        for case in cases {
            assert_eq!(case.o, longest_palindrome(case.s));
        }
    }

    #[test]
    fn test_longest_palindrome2() {
        let cases = Case::test_cases();
        for case in cases {
            assert_eq!(case.o, longest_palindrome2(case.s));
        }
    }
}

struct Case<'a> {
    s: &'a str,
    o: &'a str,
}

impl<'a> Case<'a> {
    fn new(ss: [&str; 2]) -> Case {
        if ss.len() < 2 {
            return Case { s: "", o: "" };
        }
        return Case { s: ss[0], o: ss[1] };
    }
    fn new_vec(ss_vec: Vec<[&str; 2]>) -> Vec<Case> {
        let mut cases: Vec<Case> = Vec::with_capacity(ss_vec.len());
        for ss in ss_vec {
            cases.push(Case::new(ss));
        }
        return cases;
    }
    fn test_cases<'b>() -> Vec<Case<'b>> {
        Case::new_vec(vec![
            ["", ""],
            ["a", "a"],
            ["aa", "aa"],
            ["abaaba", "abaaba"],
            ["ababa", "ababa"],
            ["abbd", "bb"],
        ])
    }
}
