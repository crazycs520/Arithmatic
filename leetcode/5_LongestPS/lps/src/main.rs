fn main() {
    let s = longest_palindrome("abaaba");
    println!("{}", s);
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
    return &s[start..end + 1];
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_longest_palindrome() {
        let cases = test_cases();
        for case in cases {
            assert_eq!(case.o, longest_palindrome(case.s));
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
}

fn test_cases() -> Vec<Case> {
    Case::new_vec(vec![
        [],
        ["a", "a"],
        ["aa", "aa"],
        ["abaaba", "abaaba"],
        ["ababa", "ababa"],
        ["abbd", "bb"],
    ])
}
