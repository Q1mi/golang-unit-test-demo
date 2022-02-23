package main

import (
	"testing"
	"time"
)

func Test_judgeRateByTime(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want int
	}{
		{
			name: "工作时间",
			arg:  time.Date(2022, 2, 18, 11, 22, 33, 0, time.UTC),
			want: 10,
		},
		{
			name: "晚上",
			arg:  time.Date(2022, 2, 18, 22, 22, 33, 0, time.UTC),
			want: 1,
		},
		{
			name: "凌晨",
			arg:  time.Date(2022, 2, 18, 2, 22, 33, 0, time.UTC),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeRateByTime(tt.arg); got != tt.want {
				t.Errorf("judgeRateByTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
