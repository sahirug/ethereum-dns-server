package util

import (
	"log"
	"regexp"
)

func ConvertByteArrToString(data [][32]byte) []string {
	var arr []string;
	for i := 0; i < len(data); i++ {
		addr := string(data[i][:])
		addr = clean(addr)
		arr = append(arr, addr)
		//arr = append(arr, "5.5.5.5")
	}
	//log.Fatalf("ded")
	return arr
}

func clean(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9.:]+")
	if err != nil {
		log.Fatal(err)
	}
	pString := reg.ReplaceAllString(s, "")
	return pString
}

