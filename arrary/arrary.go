package arrary

import (
	"math/rand"
	"reflect"
	"strings"

	"github.com/gookit/goutil/mathutil"
)

// ArrarysShuffle shuffles the array using a random source
func ArrarysShuffle(array []interface{}, source rand.Source) {
	random := rand.New(source)
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

// ArrarysReverse string slice [site user info 0] -> [0 info user site]
func ArrarysReverse(ss []string) {
	ln := len(ss)

	for i := 0; i < ln/2; i++ {
		li := ln - i - 1
		ss[i], ss[li] = ss[li], ss[i]
	}
}

// ArrarysRemove an value form an string slice
func ArrarysRemove(ss []string, s string) []string {
	var ns []string
	for _, v := range ss {
		if v != s {
			ns = append(ns, v)
		}
	}

	return ns
}

// ArrarysTrim trim string slice item.
func ArrarysTrim(ss []string, cutSet ...string) (ns []string) {
	hasCutSet := len(cutSet) > 0 && cutSet[0] != ""

	for _, str := range ss {
		if hasCutSet {
			ns = append(ns, strings.Trim(str, cutSet[0]))
		} else {
			ns = append(ns, strings.TrimSpace(str))
		}
	}
	return
}

// ArrarysGetRandomOne get random element from an array/slice
func ArrarysGetRandomOne(arr interface{}) interface{} {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return arr
	}

	i := mathutil.RandomInt(0, rv.Len())
	r := rv.Index(i).Interface()

	return r
}

// ArrarysHas check the []string contains the given element
func ArrarysHas(ss []string, val string) bool {
	for _, ele := range ss {
		if ele == val {
			return true
		}
	}
	return false
}
