#### [1185. 一周中的第几天](https://leetcode-cn.com/problems/day-of-the-week/)

> 难度简单

给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。

输入为三个整数：`day`、`month` 和 `year`，分别表示日、月、年。

您返回的结果必须是这几个值中的一个 `{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}`。

 **示例 1：**

```
输入：day = 31, month = 8, year = 2019
输出："Saturday"
```

**示例 2：**

```
输入：day = 18, month = 7, year = 1999
输出："Sunday"
```

**示例 3：**

```
输入：day = 15, month = 8, year = 1993
输出："Sunday"
```

**提示：**

- 给出的日期一定是在 `1971` 到 `2100` 年之间的有效日期。

> 这个题目不应该告诉我们1970年12月31日是星期几吗？周四

#### 解决思路

题目保证日期是在 1971 到 2100 之间，我们可以计算给定日期距离 1970 的最后一天（星期四）间隔了多少天，从而得知给定日期是周几。

具体的，可以先通过循环处理计算年份在 `[1971, year - 1]` 时间段，经过了多少天（注意平年为 365，闰年为 366）；然后再处理当前年 year的月份在 `[1, month - 1]` 时间段 ，经过了多少天（注意当天年是否为闰年，特殊处理 2月份），最后计算当前月 month经过了多少天，即再增加 day 天。

得到距离 1970 的最后一天（星期四）的天数后进行取模，即可映射回答案

```go
//判断是不是闰年
//普通闰年：公历年份是4的倍数，且不是100的倍数的，为闰年（如2004年、2020年等就是闰年）。
//世纪闰年：公历年份是整百数的，必须是400的倍数才是闰年（如1900年不是闰年，2000年是闰年）。
year%400 == 0 || year%4 == 0 && year%100 != 0
```

#### 代码

```go
package leetcode
var week = []string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day int, month int, year int) string {
	days := 0
	//输入年份之前的年份贡献的天数
	for i := 1971; i < year; i++ {
		if i%400 == 0 || i%4 == 0 && i%100 != 0 {
			days += 366
		} else {
			days += 365
		}
	}
	//输入当前月份之前贡献的天数
	for i := 1; i < month; i++ {
		days += monthDays[i-1]
		if i == 2 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
			days++
		}
	}
	//当前月份的天数贡献
	days += day
	return week[days%7]
}
```

