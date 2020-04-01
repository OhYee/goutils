
// If{{if ne . "any"}}{{upperFirstChar .}}{{end}} return value according condition
func If{{if ne . "any"}}{{upperFirstChar .}}{{end}}(condition bool, t {{.}}, f {{.}}) {{.}} {
    if condition {
        return t
    }
    return f
}
