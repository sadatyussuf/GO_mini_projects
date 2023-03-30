package logs
import (
    "unicode/utf8"
)



// Application identifies the application emitting the given log.
func Application(log string) string {
    char := map[rune]string{
            '❗': "recommendation",
            '🔍' : "search",
            '☀'  : "weather",
     }

	for _,ele := range log{
        val,exist := char[ele]
        if exist{
            return val
        } 
    }
return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    runes := []rune(log)
		for i,ele := range runes{
        if ele == oldRune{
            runes[i] = newRune
        }
    }
 return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    lenght_str :=  utf8.RuneCountInString(log)
	if lenght_str <= limit{
        return true
    }
   return false
}
