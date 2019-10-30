func Map{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}}){{.}}, input []{{.}}) (output []{{.}}){
    output = make([]{{.}}, 0)
    for _, data := range input {
        output = append(output, f(data))
    }
    return
}
