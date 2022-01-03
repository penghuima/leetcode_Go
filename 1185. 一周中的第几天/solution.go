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
	//输入当前年份是闰年还是平年
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
