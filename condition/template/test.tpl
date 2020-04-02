func TestIf{{template "functionName" .type}}(t *testing.T) {
    {{- if eq .type "function"}}var res bool{{end}}
    tests := []struct {
        name       string
        condition  bool
        t          {{.type}}
        f          {{.type}}
        wantOutput {{.type}}
    }{
        {
			name:       "{{.type}} true",
			condition:  true,
            t:          {{.t}},
            f:          {{.f}},
            wantOutput: {{.t}},
		},
		{
			name:       "{{.type}} false",
			condition:  false,
            t:          {{.t}},
            f:          {{.f}},
            wantOutput: {{.f}},
		},
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            {{- if ne .type "function"}}
                if gotOutput := If{{template "functionName" .type}}(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
                    t.Errorf("If{{template "functionName" .type}}() = %v, want %v", gotOutput, tt.wantOutput)
                }
            {{- else}}
                res = false
                If{{template "functionName" .type}}(tt.condition, tt.t, tt.f)()
                if res != tt.condition {
                    t.Errorf("If{{template "functionName" .type}}() = %v, want %v", res, tt.condition)
                }
            {{end}}
		})
	}
}
