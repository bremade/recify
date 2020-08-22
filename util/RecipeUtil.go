package util

func ContainsString(slice []string, val string) bool {
    for _, tmp := range slice {
        if tmp == val {
            return true
        }
    }
    return false
}
