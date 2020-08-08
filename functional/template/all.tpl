// All{{if ne . "any"}}{{upperFirstChar .}}{{end}} return true if all values are matched
func All{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}}, int)bool, input []{{.}}) (output bool){
    output = true
    for idx, data := range input {
        output = output && f(data, idx)
        if !output {
            break
        }
    }
    return
}
