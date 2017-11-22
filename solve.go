package main
import "unicode"

func RemoveEven(arr []int) []int {
    result := make([]int, 0)
    for _, element := range arr 
 	{
        if element % 2 == 1 
	{
            result = append(result, element)
        }
    }
    return result
}
func PowerGenerator(a int) func() int {
    x := 1
    return func() int 
	{
        x *= a
        return x
    }
}
func DifferentWordsCount(x string) int {
    word := ""
    set := make(map[string]bool)
    answer := 0
    for _, c := range (x + " ") {
        if unicode.IsLetter(c) 
	    {
            word += string(unicode.ToLower(c))
        } else if word != "" 
	    {
            if !set[word] {
                answer += 1
            }
            set[word] = true
            word = ""
        }
    }
    return answer
} 