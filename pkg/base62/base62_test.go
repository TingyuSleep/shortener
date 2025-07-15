package base62

import "testing"

func TestTo62String(t *testing.T) {
	type args struct {
		seq uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case:0", args: args{seq: 0}, want: "0"},
		{name: "case:1", args: args{seq: 1}, want: "1"},
		{name: "case:62", args: args{seq: 62}, want: "10"},
		{name: "case:6347", args: args{seq: 6347}, want: "1En"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := To62String(tt.args.seq); got != tt.want {
				t.Errorf("To62String() = %v, want %v", got, tt.want)
			}
		})
	}
}
