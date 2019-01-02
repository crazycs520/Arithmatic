impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        if s.len() == 0 {
            return ""
        }
        let mut start: usize  =  0;
        let mut end: usize = 0;

        for i in (0..s.len()) {
            let l1 = expand_around_center(&s, i, i);
            let l2 = expand_around_center(&s, i, i+1);
            let ll = l1.max(l2);
            if ll > (end - start) {
                start = i - (ll - 1)/2;
                end = i + ll / 2;
            }
        }
        return String::from(s[start..end+1])
    }
    fn expand_around_center(s: &str, mut l: usize,mut r: usize) -> usize {
        while l >= 0 && r < s.len() && s[l] == s[r] {
            l-=1;
            r+=1;
        }
        return r-l-1
    }
}