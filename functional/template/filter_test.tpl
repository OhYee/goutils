{{define "testcases"}}
    type args struct {
        f     func({{.}}) bool
        input []{{.}}
    }
    tests := []struct {
        name       string
        args       args
        wantOutput []{{.}}
    }{
{{- if (eq . "string")}}
        {
			name:       "test all true",
			args:       args{f: func(s string) bool { return true }, input: []string{"a", "b", "C", "1"}},
			wantOutput: []string{"a", "b", "C", "1"},
		},
		{
			name:       "test len > 2",
			args:       args{f: func(s string) bool { return len(s) > 2 }, input: []string{"a", "b", "C", "1"}},
			wantOutput: []string{},
		},
		{
			name:       "test len > 1",
			args:       args{f: func(s string) bool { return len(s) > 1 }, input: []string{"ab", "bc", "Cd", "1"}},
			wantOutput: []string{"ab", "bc", "Cd"},
		},
		{
			name:       "test first rune is a",
			args:       args{f: func(s string) bool { return s[0] == 'a' }, input: []string{"ab", "bc", "Cd", "1"}},
			wantOutput: []string{"ab"},
		},
{{else if eq . "any"}}
        {
			name:       "is int",
			args:       args{f: func(s any) bool { return reflect.TypeOf(s).Kind() == reflect.Int }, input: []any{1, "A", 2, "c"}},
			wantOutput: []any{1, 2},
		},
        {
			name:       "is string",
			args:       args{f: func(s any) bool { return reflect.TypeOf(s).Kind() == reflect.String }, input: []any{1, "A", 2, "c"}},
			wantOutput: []any{"A", "c"},
		},
{{else}}
        {
            name:       "more than 3",
            args:       args{f: func(n {{.}}) bool { return n > 3 }, input: []{{.}}{1, 2, 3, 4, 5}},
            wantOutput: []{{.}}{4, 5},
        },
        {
            name:       "less than 2",
            args:       args{f: func(n {{.}}) bool { return n < 2 }, input: []{{.}}{1, 2, 3, 4, 5}},
            wantOutput: []{{.}}{1},
        },{{if and (ne . "float32") (ne . "float64") }}
        {
            name:       "even number",
            args:       args{f: func(n {{.}}) bool { return n%2 == 0 }, input: []{{.}}{1, 2, 3, 4, 5}},
            wantOutput: []{{.}}{2, 4},
        },{{end}}
{{end}}
    }
{{end}}