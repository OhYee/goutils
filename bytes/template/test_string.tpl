func TestFromString(t *testing.T) {
	tests := []struct {
		name  string
		value string
		wantB []byte
	}{
        {{- range $case := .testcases}}
		{
			name:"test {{$.type}} {{$case.value}}",
			value: "{{$case.value}}",
			wantB:[]byte{ {{$case.b}} },
		},
        {{end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := FromString(tt.value); !goutils.Equal(gotB, tt.wantB) {
				t.Errorf("FromString() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

