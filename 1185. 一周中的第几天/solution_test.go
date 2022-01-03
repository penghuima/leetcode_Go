package leetcode

import "testing"

func TestDayOfTheWeek(t *testing.T){
	//周日
	day,month,year:=18,7,1999
	t.Logf("那天是星期几：%s",dayOfTheWeek(day,month,year))
}