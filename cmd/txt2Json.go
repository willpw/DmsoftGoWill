package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	dmsoft "github.com/willpw/DmsoftGoWill"
)

func main() {
	DllJson := make(map[string]string)
	_ = json.Unmarshal(dmsoft.DllJson, &DllJson)
	txtfile, err := os.ReadFile("cmd/读我.txt")
	if err != nil {
		log.Panic(err)
	}

	for _, str := range strings.Split(string(txtfile), "\n") {
		str = strings.TrimSpace(str)
		switch {
		case StrBool(str, "对象名:"):
			s := strings.Split(str, ":")[1]

			if s != "" {
				DllJson["Willpwr"] = s
			}

		case StrBool(str, "DmReg.dll"):

			for _, S := range strings.Split(str, ",") {
				key, s := "", ""
				switch {
				case StrBool(S, "SetDllPathA"):
					key = "SetDllPathA"
				case StrBool(S, "SetDllPathW"):
					key = "SetDllPathW"
				}
				if StrBool(S, "函数名改为:") {
					s = strings.Split(S, "函数名改为:")[1]
					if key != "" && s != "" {
						log.Println(key, s)
						DllJson[key] = s
					}
				}

			}

		case StrBool(str, "->") && !StrBool(str, "删除"):
			key, s := strings.Split(str, "->")[0], strings.Split(str, "->")[1]
			if StrBool(key, "(") {
				key = strings.Split(str, "(")[0]
			}

			if key != "" && s != "" {
				DllJson[key] = s
			}

		default:
			log.Println(str)
		}
	}
	DllJson["DmReg.dll"] = "reg.dll"
	Jsonfile, _ := json.MarshalIndent(DllJson, "", "\t")

	_ = os.WriteFile("Tdll.json", Jsonfile, 0644)

	fmt.Println(DllJson["Willpwr"], DllJson["DmReg.dll"], DllJson["SetDllPathA"], DllJson["SetDllPathW"])
}

func StrBool(Str string, str ...string) bool {
	for _, s := range str {
		if s != "" && strings.Contains(Str, s) {
			return true
		}
	}
	return false
}
