package assignment

import (
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y
	overflow := false
	if sum < x || sum < y {
		overflow = true
	}
	return sum, overflow
}

func CeilNumber(f float64) float64 {
	beforePoint := int(f)
	afterPoint := f - float64(beforePoint)
	switch {
	case afterPoint == 0:
		afterPoint = 0
	case afterPoint <= 0.25:
		afterPoint = 0.25
	case afterPoint <= 0.50:
		afterPoint = 0.50
	case afterPoint <= 0.75:
		afterPoint = 0.75
	default:
		afterPoint = 1
	}
	//math.Ceil(f)
	return float64(beforePoint) + afterPoint
}

func AlphabetSoup(s string) string {
	//tersine eklemeye bakıcam
	output := strings.Split(s, "")
	sort.Slice(output, func(i, j int) bool {
		return output[j] > output[i]
	})
	return strings.Join(output, "")
}

func StringMask(s string, n uint) string {
	output := strings.Split(s, "")
	if s == "" {
		return "*"
	}
	if uint(len(s)) > n {
		short_output := output[n:]
		for i := 0; i < len(short_output); i++ {
			short_output[i] = "*"
		}
	} else {
		for i := 0; i < len(output); i++ {
			output[i] = "*"
		}
	}
	return strings.Join(output, "")
}

func WordSplit(arr [2]string) string {
	words := strings.Split(arr[1], ",")
	findWords := []string{}
	sentences := arr[0][:]
	isFind := false
	for len(sentences) != 0 {
		isFind = false
		for _, v := range words {
			if find := strings.Index(sentences, v); find == 0 {
				findWords = append(findWords, v)
				sentences = sentences[find+len(v):]
				isFind = true
			}
		}
		if !isFind {
			isFind = false
			return "not possible"
		}
	}
	if len(findWords) == 0 {
		return "not possible"
	}
	return strings.Join(findWords, ",")
}

func VariadicSet(i ...interface{}) []interface{} {
	count := make(map[interface{}]interface{})
	keys := []interface{}{}
	for _, v := range i {
		_, ok := count[v]
		if !ok {
			keys = append(keys, v)
			count[v] = 1
		}
	}
	return keys
}
