// All{{if ne . "any"}}{{upperFirstChar .}}{{end}} return true if all values are matched
func All{{if ne . "any"}}{{upperFirstChar .}}{{end}}(f func({{.}})bool, input []{{.}}) (output bool){
    output = true
    for _, data := range input {
        output = output && f(data)
        if !output {
            break
        }
    }
    return
}
