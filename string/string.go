package string

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/bsync-tech/util/xrunes"
)

const (
	space = " "
)

// IsEmpty returns true if the string is empty
func IsEmpty(text string) bool {
	return len(text) == 0
}

// IsNotEmpty returns true if the string is not empty
func IsNotEmpty(text string) bool {
	return !IsEmpty(text)
}

// IsBlank returns true if the string is blank (all whitespace)
func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

// IsNotBlank returns true if the string is not blank
func IsNotBlank(text string) bool {
	return !IsBlank(text)
}

// Reverse reverses the input while respecting UTF8 encoding and combined characters
func Reverse(text string) string {
	tr := []rune(text)
	trl := len(tr)
	if trl <= 1 {
		return text
	}

	i, j := 0, 0
	for i < trl && j < trl {
		j = i + 1
		for j < trl && xrunes.IsMark(tr[j]) {
			j++
		}

		if xrunes.IsMark(tr[j-1]) {
			// Reverses Combined Characters
			reverse(tr[i:j], j-i)
		}

		i = j
	}

	// Reverses the entire array
	reverse(tr, trl)

	return string(tr)
}

func reverse(runes []rune, length int) {
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

// 字符串数组去重 不保序
func StringsUniq(l []string) []string {
	m := make(map[string]interface{})
	if len(l) <= 0 {
		return []string{}
	}
	for _, v := range l {
		m[v] = "true"
	}
	var datas []string
	for k := range m {
		if k == "" {
			continue
		}
		datas = append(datas, k)
	}
	return datas
}

const (
	empty = ""
	tab   = "\t"
)

// JSONString returns a string representation
func JSONString(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return empty, err
	}
	return string(b), nil
}

// DecodeString returns the original representation
func DecodeString(data string, value interface{}) error {
	return json.Unmarshal([]byte(data), value)
}

func JSONStringPretty(v interface{}) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// StringsToInts string slice to int slice. alias of the arrutil.StringsToInts()
func StringsToInts(ss []string) (ints []int, err error) {
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}
	return
}

func Interface2string(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int, int64:
		return strconv.FormatInt(v.(int64), 10)
	case uint, uint64:
		return strconv.FormatUint(v.(uint64), 10)
	case float32:
		return strconv.FormatFloat(v.(float64), 'f', 6, 32)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', 6, 64)
	case bool:
		return strconv.FormatBool(v.(bool))
	default:
		return ""
	}
}

func SqlEscape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
			break
		case '\n': /* Must be escaped for logs */
			escape = 'n'
			break
		case '\r':
			escape = 'r'
			break
		case '\\':
			escape = '\\'
			break
		case '\'':
			escape = '\''
			break
		case '"': /* Better safe than sorry */
			escape = '"'
			break
		case '\032': /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}

