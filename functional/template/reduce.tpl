// Reduce{{if ne . "any"}}{{upperFirstChar .}}{{end}} accumulate all values
func Reduce{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}}, {{.}}){{.}}, input []{{.}}) (output {{.}}){
    for _, data := range input {
        output = f(output, data)
    }
    return
}
