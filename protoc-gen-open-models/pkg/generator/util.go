package generator

import "strings"

func multilineComment(comments string) []string {
	if comments == "" {
		return []string{}
	}
	// remove last \n
	if strings.HasSuffix(comments, "\n") {
		comments = comments[:len(comments)-1]
	}

	return strings.Split(comments, "\n")
}

func multilineCommentString(comments string) string {
	items := multilineComment(comments)
	sArr := []string{}
	for _, item := range items {
		sArr = append(sArr, "// "+item)
	}
	return strings.Join(sArr, "\n")
}
