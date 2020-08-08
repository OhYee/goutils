{{define "testcases"}}
    type args struct {
        f     func({{.}}, {{.}}, int) {{.}}
        input []{{.}}
    }
    tests := []struct {
        name       string
        args       args
        wantOutput {{.}}
    }{
{{- if (eq . "string")}}
		{
			name:       "concat string",
			args:       args{f: func(a {{.}}, b {{.}}, idx int) {{.}} { return a + b }, input: []{{.}}{"a", "bb", "ccc", "dddd"}},
			wantOutput: "abbcccdddd",
		},
{{else if eq . "any"}}
        {
			name:       "output as string",
			args:       args{f: func(a {{.}}, b {{.}}, idx int) {{.}} { return fmt.Sprint(a) + fmt.Sprint(b) }, input: []{{.}}{"a", "bb", 3, 4}},
			wantOutput: "<nil>abb34",
		},
{{else}}
        {
			name:       "sum",
			args:       args{f: func(a {{.}}, b {{.}}, idx int) {{.}} { return a + b }, input: []{{.}}{1, 2, 3, 4}},
			wantOutput: 10,
		},
{{end}}
    }
{{end}}