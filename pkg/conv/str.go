package conv

import (
	"strconv"
)

func StrRaw2Interpreted(str string) (string, error) {
	return strconv.Unquote(`"` + str + `"`)
}
