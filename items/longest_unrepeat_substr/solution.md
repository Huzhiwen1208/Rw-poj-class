## 最长连续不重复子序列
### Description

给定一个字符串，请你找出其中不含有重复字符的最长连续子串的长度。


- Input
输入包含两行，第一行为一个正整数n(n∈(0,10000000])，第二行为一个长度为 n 的字符串。

- Output
输出单独的一行，即给定的最长不重复子串的长度。


Sample Input 1 

```
8
abcabcbb
```
Sample Output 1

```
3
```
Sample Input 2 

```
5
bbbbb
```
Sample Output 2
```
1
```
Sample Input 3 
```
6
pwwkew
```
Sample Output 3
```
3
```

对于此题，显而易见地使用滑动窗口解决，我们保持一个窗口，内部的字符均不重复，如果向右扩展导致窗口内有重复字符则进行向右滑动（即丢弃最左边元素）直到没有重复字符。我们定义窗口的范围为：[currentLeft, currentRight]。持续维护窗口内字符的唯一性，并更新最大值的窗口大小。详见代码`longest_unrepeat_substr.go`