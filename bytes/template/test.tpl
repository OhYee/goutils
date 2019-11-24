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
			if gotB := From{{upperFirstChar .type}}(tt.value); !goutils.Equal(gotB, tt.wantB) {
				t.Errorf("From{{upperFirstChar .type}}() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

