{{define "testcases"}}
    type args struct {
        f     func({{.}}, int) {{.}}
        input []{{.}}
    }
    tests := []struct {
        name       string
        args       args
        wantOutput []{{.}}
    }{
{{- if (eq . "string")}}
		{
			name:       "add a 1 at the end of the string",
			args:       args{f: func(s {{.}}, idx int) {{.}} { return s + "1" }, input: []{{.}}{"a", "bb", "ccc", "dddd"}},
			wantOutput: []{{.}}{"a1", "bb1", "ccc1", "dddd1"},
		},
{{else if eq . "any"}}
        {
			name:       "change all int value + 1 or output aaa",
			args:       args{
                f: func(s {{.}}, idx int) {{.}} {
                    if reflect.TypeOf(s).Kind() == reflect.Int {
                        return s.(int) + 1
                    }
                    return "aaa"
                }, 
                input: []{{.}}{1, 2, "c", "d"},
            },
			wantOutput: []{{.}}{2, 3, "aaa", "aaa"},
		},
{{else}}
        {
			name:       "all value + 1",
			args:       args{f: func(n {{.}}, idx int) {{.}} { return n + 1 }, input: []{{.}}{1, 2, 3, 4, 5}},
			wantOutput: []{{.}}{2, 3, 4, 5, 6},
		},
        {
			name:       "all value - 1",
			args:       args{f: func(n {{.}}, idx int) {{.}} { return n - 1 }, input: []{{.}}{1, 2, 3, 4, 5}},
			wantOutput: []{{.}}{0, 1, 2, 3, 4},
		},
{{end}}
    }
{{end}}