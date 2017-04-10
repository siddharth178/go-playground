package main

import (
	"bytes"
	"testing"
)

var (
	usersData = `name,age
F1 L1,30
F2 L2,20
F3 L3,30
F4 L4,20
F5 L5,30
F6 L6,20
F7 L7,30
F8 L8,20
F9 L9,70`
)

func Test_countRecords(t *testing.T) {
	cnt, err := countRecords(bytes.NewBufferString(usersData), &UserCounter{})
	if err != nil {
		t.Error(err)
	}
	if cnt != 9 {
		t.Errorf("expected: %d, actual: %d", 9, cnt)
	}
}

func Benchmark_countRecords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := countRecords(bytes.NewBufferString(usersData), &UserCounter{})
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_countRecordsTheOldWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := countRecordsTheOldWay(bytes.NewBufferString(usersData))
		if err != nil {
			b.Error(err)
		}
	}
}
