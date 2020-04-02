
// If{{template "functionName" .type}} return value according condition
func If{{template "functionName" .type}}(condition bool, t {{.type}}, f {{.type}}) {{.type}} {
    if condition {
        return t
    }
    return f
}
