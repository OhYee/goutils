func TestFrom{{upperFirstChar .type}}(t *testing.T) {
	tests := []struct {
		name  string
		value {{.type}}
		wantB []byte
	}{
        {{- range $case := .testcases}}
		{
			name:"test {{$.type}} {{$case.value}}",
			value: {{$case.value}},
			wantB:[]byte{ {{$case.b}} },
		},
        {{end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := From{{upperFirstChar .type}}(tt.value); !compare.Equal(gotB, tt.wantB) {
				t.Errorf("From{{upperFirstChar .type}}() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestTo{{upperFirstChar .type}}(t *testing.T) {
	tests := []struct {
		name  string
		b     []byte
		want  {{.type}}
	}{
        {{- range $case := .testcases}}
		{
			name:"test {{$.type}} {{$case.value}}",
			b: []byte{ {{$case.b}} },
			want: {{$case.value}},
		},
        {{end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := To{{upperFirstChar .type}}(tt.b); !compare.Equal(got, tt.want) {
				t.Errorf("To{{upperFirstChar .type}}() = %v, want %v", got, tt.want)
			}
		})
	}
}
