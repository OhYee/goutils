// Any{{if ne . "any"}}{{upperFirstChar .}}{{end}} return true if any of values is matched
func Any{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}}, int)bool, input []{{.}}) (output bool){
    output = false
    for idx, data := range input {
        output = output || f(data, idx)
        if output {
            break
        }
    }
    return
}
