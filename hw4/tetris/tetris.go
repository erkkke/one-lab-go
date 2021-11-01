package tetris

import (
	"fmt"
	"reflect"
	"sort"
)

type fieldInfo struct {
	Name string
	typeName string
	size int
}

func Tetris(myStruct interface{}) {
	if reflect.ValueOf(myStruct).Kind().String() != "struct"  {
		fmt.Println("Type should be the struct")
		return
	}

	fields := make([]fieldInfo, 0)
	for i := 0; i < reflect.TypeOf(myStruct).NumField(); i++ {
		field := reflect.TypeOf(myStruct).FieldByIndex([]int{i})
		tmp := fieldInfo{field.Name, field.Type.Name(), int(field.Type.Size())}
		fields = append(fields, tmp)
	}

	sort.Slice(fields, func(i, j int) bool { return fields[i].size > fields[j].size })

	generateAndPrintBestThreeCases(fields)
}

func generateAndPrintBestThreeCases(fields []fieldInfo) {
	bestThreeCases := make([][]fieldInfo, 3)
	for i := 0; i < 3; i++ {
		tmp := make([]fieldInfo, len(fields))
		copy(tmp, fields)
		bestThreeCases[i] = tmp
		fields[i], fields[i+1] = fields[i+1], fields[i]
	}

	for i, res := range bestThreeCases {
		fmt.Printf("Case #%v:\n", i + 1)
		for _, v := range res {
			fmt.Printf("	%v %v    //%v byte\n", v.Name, v.typeName, v.size)
		}
		fmt.Println()
	}

}


