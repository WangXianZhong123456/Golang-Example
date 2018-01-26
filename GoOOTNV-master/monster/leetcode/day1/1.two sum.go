/**
 * @Difficulty: Easy
 * @Description: Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 *               给定一个整数数组, 将两个数字相加得到特定的目标, 返回这两个数字的索引。
 * @Example: Given nums = [2, 7, 11, 15], target = 9,
 *           Because nums[0] + nums[1] = 2 + 7 = 9,
 *           return [0, 1].
 */

func twoSum(nums []int, target int) []int {
    // m 负责保存map[整数]整数的序列号
    m := make(map[int]int, 2)

    // 通过 for 循环，获取b的序列号
    for i, b := range nums {
        // 通过查询map，获取a = target - b的序列号
        if j, ok := m[target-b]; ok {
            // ok 为 true
            // 说明在i之前，存在nums[j] == a
            return []int{j, i}
            // 注意，顺序是j，i，因为j<i
        }

        // 把i的值，存入map
        m[nums[i]] = i
    }

    return nil
}
