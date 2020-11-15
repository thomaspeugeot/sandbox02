package walk

import (
	"bytes"
	"strings"
)

func replace2(template, match1, replace1, match2, replace2 string) string {

	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)

	return res
}

func replace3(template, match1, replace1, match2, replace2, match3, replace3 string) string {

	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)
	res = strings.ReplaceAll(res, match3, replace3)

	return res
}

func replace4(template, match1, replace1, match2, replace2, match3, replace3, match4, replace4 string) string {

	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)
	res = strings.ReplaceAll(res, match3, replace3)
	res = strings.ReplaceAll(res, match4, replace4)

	return res
}

func makeFirstLowerCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
