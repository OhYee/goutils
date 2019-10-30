func Filter{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}})bool, input []{{.}}) (output []{{.}}){
    output = make([]{{.}}, 0)
    for _, data := range input {
        if f(data) {
            output = append(output, data)
        }
    }
    return
}
