// Filter{{if ne . "any"}}{{upperFirstChar .}}{{end}} return the values which are matched
func Filter{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}}, int)bool, input []{{.}}) (output []{{.}}){
    output = make([]{{.}}, 0)
    for idx, data := range input {
        if f(data, idx) {
            output = append(output, data)
        }
    }
    return
}
