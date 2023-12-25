package logs

import ( 
	"unicode/utf8"
    "strings"
    )

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
        switch {
            case v == '❗':
        		return "recommendation"
            case v == '🔍':
        		return "search"
            case v == '☀':
        		return "weather"
        }
    }
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    log = strings.Replace(log, string(oldRune), string(newRune), -1)
	return log
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
        return true
    }
	return false
}
