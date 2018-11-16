package main

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {

	type args struct {
		email string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{email: "nobby.phala@gmail.com"}, true},
		{"Test2", args{email: "ja.ne.doe@tokopedia.com"}, true},
		{"Test3", args{email: "doe@doe.com"}, true},
		{"Test4", args{email: "12A3213__123@numbers.com"}, true},
		{"Test5", args{email: "tokopedia.com"}, false},
		{"Test6", args{email: "john.smith@tokopedia"}, false},
		{"Test7", args{email: "jack.sparrow"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValiddateEmail(tt.args.email); got != tt.want {
				t.Errorf("Testing gagal %v:%v", tt.args.email, got)
			}
		})
	}
}
