{{define "testcases"}}
    type args struct {
        f     func({{.}}, int) bool
        input []{{.}}
    }
    tests := []struct {
        name       string
        args       args
        wantOutput bool
    }{
{{- if (eq . "string")}}
        {
			name:       "all item is true",
			args:       args{f: func(s {{.}}, idx int) bool { return len(s) > 0 }, input: []{{.}}{"a", "bb", "ccc", "dddd"}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(s {{.}}, idx int) bool { return len(s) > 2 }, input: []{{.}}{"a", "bb", "ccc", "dddd"}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(s {{.}}, idx int) bool { return len(s) > 5 }, input: []{{.}}{"a", "bb", "ccc", "dddd"}},
			wantOutput: false,
		},
{{else if eq . "any"}}
        {
			name:       "all item is true",
			args:       args{f: func(s {{.}}, idx int) bool { return reflect.TypeOf(s).Kind() == reflect.Int }, input: []{{.}}{1, 2}},
			wantOutput: true,
		},
        {
			name:       "some items are true and others are false",
			args:       args{f: func(s {{.}}, idx int) bool { return reflect.TypeOf(s).Kind() == reflect.Int }, input: []{{.}}{1, 2, "A"}},
			wantOutput: true,
		},
        {
            name:       "all item is false",
            args:       args{f: func(s {{.}}, idx int) bool { return reflect.TypeOf(s).Kind() == reflect.Int }, input: []{{.}}{"A", "B"}},
            wantOutput: false,
        },
{{else}}
        {
			name:       "all item is true",
			args:       args{f: func(n {{.}}, idx int) bool { return n > 0 }, input: []{{.}}{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
        {
			name:       "some items are true and others are false",
			args:       args{f: func(n {{.}}, idx int) bool { return n > 3 }, input: []{{.}}{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
        {
            name:       "all item is false",
            args:       args{f: func(n {{.}}, idx int) bool { return n > 5 }, input: []{{.}}{1, 2, 3, 4, 5}},
            wantOutput: false,
        },
{{end}}
    }
{{end}}