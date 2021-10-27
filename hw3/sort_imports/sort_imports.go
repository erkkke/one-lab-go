package sort_imports

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func SortAndRewriteImports(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	content := make([]string, 0)
	imports := make([]string, 0)
	isImport := false
	for scanner.Scan() {
		content = append(content, scanner.Text())
		if scanner.Text() == "import (" {
			isImport = true
			continue
		}
		if isImport {
			if scanner.Text() == ")" {
				isImport = false
				continue
			}
			imports = append(imports, scanner.Text())
		}
	}

	sort.Strings(imports)
	for i, v := range content {
		if v == "import (" {
			i++
			j := 0
			for content[i] != ")" && j < len(imports){
				content[i] = imports[j]
				i++
				j++
			}
			break
		}
	}

	result := new(strings.Builder)
	for _, str := range content {
		result.WriteString(str)
		result.WriteString("\n")
	}

	if err = cleanFile(path); err != nil{
		return err
	}
	if _, err = f.Seek(0, 0); err != nil {
		return err
	}
	if _, err = f.WriteString(result.String()); err != nil {
		return err
	}

	return nil
}

func cleanFile(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	if err = f.Truncate(0); err != nil {
		return err
	}
	return nil
}