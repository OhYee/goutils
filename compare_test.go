package goutil

import "testing"

func TestEqual(t *testing.T) {
	type args struct {
		a Any
		b Any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "bytes true",
			args: args{[]byte{0x01, 0x02}, []byte{0x01, 0x02}},
			want: true,
		},
		{
			name: "bytes false",
			args: args{[]byte{0x01, 0x02}, []byte{0x01}},
			want: false,
		},
		{
			name: "string true",
			args: args{"Abc1", "Abc1"},
			want: true,
		},
		{
			name: "string false",
			args: args{"Abc1", "abc1"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
