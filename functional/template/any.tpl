// Any{{if ne . "any"}}{{upperFirstChar .}}{{end}} return true if any of values is matched
func Any{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}})bool, input []{{.}}) (output bool){
    output = false
    for _, data := range input {
        output = output || f(data)
        if output {
            break
        }
    }
    return
}
