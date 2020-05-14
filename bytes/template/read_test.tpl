func TestRead{{upperFirstChar .type}}(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    {{.type}}
        wantErr bool
	}{
        {{- range $case := .testcases}}
		{
			name:"test {{$.type}} {{$case.value}}",
			b: []byte{ {{$case.b}} },
			want: {{$case.value}},
            wantErr: false,
		},
        {{end}}
        {
			name:"test empty bytes",
			b: []byte{},
			want: 0,
            wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            got, err := Read{{upperFirstChar .type}}(bytes.NewBuffer(tt.b))
            if (err != nil) != tt.wantErr {
                t.Errorf("Read{{upperFirstChar .type}} want error %v, but got %v\n", tt.wantErr, err)
            }
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("Read{{upperFirstChar .type}}() = %v, want %v", got, tt.want)
			}
		})
	}
}
