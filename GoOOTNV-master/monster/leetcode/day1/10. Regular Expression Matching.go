/**
 * @Difficulty: Hard
 * @Description: Implement regular expression matching with support for '.' and '*'.
 *               实现支持“.”和“*”的正则表达式匹配。
 * @Example: isMatch("aa","a") → false
 *           isMatch("aa","aa") → true
 *           isMatch("aa", "a*") → true
 *           isMatch("aa", ".*") → true
 *           isMatch("ab", ".*") → true
 *           isMatch("aab", "c*a*b") → true
 */

func isMatch(s string, p string) bool {
    lens := len(s)
    lenp := len(p)
    if lens == 0 && lenp ==0{
        return true
    }
    dp := make([][]bool,lenp+1)
    for i:=0;i<=lenp;i++{
        dp[i] = make([]bool,lens+1)
    }
    dp[0][0] =true
    for i:=2;i<=lenp;i+=2{
        if p[i-1] == '*'{
            dp[i][0] = true
        }else{
            break
        }
    }
    for i :=1;i<=lenp;i++{
        if p[i-1] =='*'{
            for j:=1;j<=lens;j++{
                dp[i][j] = dp[i-2][j]||((p[i-2]=='.'||p[i-2]==s[j-1])&&dp[i][j-1])
            }
        }else{
            for j:=1;j<=lens;j++{
                dp[i][j] = dp[i-1][j-1]&&(p[i-1]=='.'||p[i-1]==s[j-1])
            }
        }
    }
    return dp[lenp][lens]
}