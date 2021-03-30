/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 *
 * https://leetcode.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (23.76%)
 * Total Accepted:    513.3K
 * Total Submissions: 2.1M
 * Testcase Example:  '"the sky is blue"'
 *
 * Given an input string s, reverse the order of the words.
 *
 * A word is defined as a sequence of non-space characters. The words in s will
 * be separated by at least one space.
 *
 * Return a string of the words in reverse order concatenated by a single
 * space.
 *
 * Note that s may contain leading or trailing spaces or multiple spaces
 * between two words. The returned string should only have a single space
 * separating the words. Do not include any extra spaces.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "the sky is blue"
 * Output: "blue is sky the"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "  hello world  "
 * Output: "world hello"
 * Explanation: Your reversed string should not contain leading or trailing
 * spaces.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "a good   example"
 * Output: "example good a"
 * Explanation: You need to reduce multiple spaces between two words to a
 * single space in the reversed string.
 *
 *
 * Example 4:
 *
 *
 * Input: s = "  Bob    Loves  Alice   "
 * Output: "Alice Loves Bob"
 *
 *
 * Example 5:
 *
 *
 * Input: s = "Alice does not even like bob"
 * Output: "bob like even not does Alice"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s contains English letters (upper-case and lower-case), digits, and spaces '
 * '.
 * There is at least one word in s.
 *
 *
 *
 * Follow up: Could you solve it in-place with O(1) extra space?
 *
 */
func reverseWords(s string) string {
	s = removeSpace(s)
	s = reverse(s, 0, len(s)-1)
	l := len(s)
	if l == 0 {
		return s
	}

	start := 0
	end := 0

	inword := false

	for i := 0; i < l; i++ {
		if (inword) && s[i-1] != ' ' && s[i] == ' ' {
			end = i - 1
			inword = false
			s = reverse(s, start, end)
		}

		if (!inword) || (s[i] != ' ' && s[i-1] == ' ') {
			inword = true
			start = i
		}

		if (inword) && i == l-1 && s[i] != ' ' {
			end = i
			inword = false
			s = reverse(s, start, end)
		}
	}
	return s
}

func removeSpace(ss string) string {
	s := []byte(ss)
	l := len(s)
	if l == 0 {
		return string(s)
	}

	left := 0
	right := 0
	for ; right < l && s[right] == ' '; right++ {
	}
	for ; right < l; right++ {
		if right-1 > 0 && s[right-1] == ' ' && s[right] == ' ' {
			continue
		}
		s[left] = s[right]
		left++
	}

	if left-1 > 0 && s[left-1] == ' ' {
		s = s[:left-1]
	} else {
		s = s[:left]
	}

	return string(s)

}

func reverse(ss string, left, right int) string {
	s := []byte(ss)
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}
