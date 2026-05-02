func longestCommonPrefix(strs []string) string {
    first := strs[0]
    for i := 0; i < len(first); i++ {
        for _,str := range strs{
            if i >= len(str){
                return first[:i]
            }
            if first[i] != str[i]{
                return first[:i]
            }
        }
    }
    return first
}
