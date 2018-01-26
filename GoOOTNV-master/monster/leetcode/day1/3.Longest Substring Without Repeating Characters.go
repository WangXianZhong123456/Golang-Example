/**
 * @Difficulty: Medium
 * @Description: Given a string, find the length of the longest substring without repeating characters.
 *               给定一个字符串，找到最长的子字符串的长度，不包含重复字符。
 * @Example: Given "abcabcbb", the answer is "abc", which the length is 3.
 *           Given "bbbbb", the answer is "b", with the length of 1.
 *           Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

func lengthOfLongestSubstring(s string) int {
    /**
    * location[s[i]] == j 表示：
    * s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
    * location[s[i]] == -1 表示： s[i] 在s中第一次出现
    */
    location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
    for i := range location {
        location[i] = -1 // 先设置所有的字符都没有见过
    }

    maxLen, left := 0, 0

    for i := 0; i < len(s); i++ {
        /**
         * 说明s[i]已经在s[left:i+1]中重复了
         * 并且s[i]上次出现的位置在location[s[i]]
         */
        if location[s[i]] >= left {
            left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
        } else if i+1-left > maxLen {
            // fmt.Println(s[left:i+1])
            maxLen = i + 1 - left
        }
        location[s[i]] = i
    }

    return maxLen
}