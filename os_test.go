package goutil

import "testing"

func TestPathConcat(t *testing.T) {
	type args struct {
		path []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "path separator with /",
			args: args{
				path: []string{"C:/Windows", "System32/drivers/", "etc", "host"},
			},
			want: "C:/Windows/System32/drivers/etc/host",
		},
		{
			name: "path separator with \\",
			args: args{
				path: []string{"C:\\\\Windows", "System32\\\\drivers\\\\", "etc", "host"},
			},
			want: "C:/Windows/System32/drivers/etc/host",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathConcat(tt.args.path...); got != tt.want {
				t.Errorf("PathConcat() = %v, want %v", got, tt.want)
			}
		})
	}
}
