// Reduce{{if ne . "any"}}{{upperFirstChar .}}{{end}} accumulate all values
func Reduce{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}}, {{.}}, int){{.}}, input []{{.}}) (output {{.}}){
    for idx, data := range input {
        output = f(output, data, idx)
    }
    return
}
