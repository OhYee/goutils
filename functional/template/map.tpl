func Map{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}}, int){{.}}, input []{{.}}) (output []{{.}}){
    output = make([]{{.}}, 0)
    for idx, data := range input {
        output = append(output, f(data, idx))
    }
    return
}
