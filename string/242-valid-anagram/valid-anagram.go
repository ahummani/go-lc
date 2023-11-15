func isAnagram(s string, t string) bool {
    charsA := strings.Split(s,"")
    charsB := strings.Split(t,"")
    sort.Strings(charsA)
    sort.Strings(charsB)
    sA := strings.Join(charsA, "")
    sB := strings.Join(charsB, "")
    // fmt.Println(sA, sB)
    if sA == sB{
        return true
    }
    return false
}