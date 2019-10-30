{{$args := .}}
{{define "functionName"}}{{upperFirstChar .function}}{{if ne .type "any"}}{{upperFirstChar .type}}{{end}}{{end}}
{{range $type := .types}}{{$arg := (makeMap "function" $args.function "type" $type)}}
func Test{{template "functionName" $arg}}(t *testing.T) {
    {{- template "testcases" $type -}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := {{template "functionName" $arg}}(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("{{template "functionName" $arg}}() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
{{end}}