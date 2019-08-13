package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type treeItem struct {
	key   string
	value string

	next *treeItem

	firstChild *treeItem
	lastChild  *treeItem
}

func getBaseItemFromKey(base *treeItem, key string) *treeItem {
	for item := base; item != nil; item = item.next {
		if item.key == key {
			return item
		}
	}

	return nil
}

func addChild(base *treeItem, child *treeItem) {
	if base.firstChild == nil {
		base.firstChild = child
		base.lastChild = child
	} else {
		base.lastChild.next = child
		base.lastChild = base.lastChild.next
	}
}

func getCommandBeetweenItems(next *treeItem) string {
	if next == nil {
		return ""
	}
	return ","
}

func printChildJSON(base *treeItem, level int) {
	levelSpace := ""
	for i := 0; i < level; i++ {
		levelSpace = fmt.Sprintf("%s  ", levelSpace)
	}
	for item := base.firstChild; item != nil; item = item.next {
		commaBeeteenItems := getCommandBeetweenItems(item.next)

		fmt.Printf("%s\"%s\": ", levelSpace, item.key)
		if item.value != "" {
			fmt.Printf("\"%s\"%s\n", item.value, commaBeeteenItems)
		}
		if item.firstChild != nil {
			fmt.Printf("{\n")
			printChildJSON(item, level+1)
			fmt.Printf("%s}%s\n", levelSpace, commaBeeteenItems)
		}
	}
}

func printChildYAML(base *treeItem, level int) {
	levelSpace := ""
	for i := 0; i < level; i++ {
		levelSpace = fmt.Sprintf("%s  ", levelSpace)
	}
	for item := base.firstChild; item != nil; item = item.next {
		fmt.Printf("%s%s: ", levelSpace, item.key)
		if item.value != "" {
			fmt.Printf("\"%s\"\n", item.value)
		}
		if item.firstChild != nil {
			fmt.Printf("\n")
			printChildYAML(item, level+1)
		}
	}
}

func customUsage() {
	flag.CommandLine.SetOutput(os.Stdout)
	fmt.Println("envfile 0.0.1.")
	fmt.Println("Usage: envfile [OPTION]")
	flag.PrintDefaults()
}

func main() {

	prefix := flag.String("prefix", "", "The prefix of the environment variables.")
	removePrefix := flag.Bool("r", false, "Is the prefix removed from the variable name ?")
	keySeparator := flag.String("separator", "__", "The key separator.")
	formatter := flag.String("formatter", "json", "The output formatter")
	verbose := flag.Bool("v", false, "Verbose.")
	version := flag.Bool("version", false, "Print the version.")
	help := flag.Bool("help", false, "Print the help.")

	flag.Usage = customUsage
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *version {
		fmt.Println("envfile 0.0.1")
		os.Exit(0)
	}

	if *verbose {
		fmt.Println("prefix is '", *prefix, "'")
		fmt.Println("removePrefix prefix is '", *removePrefix, "'")
		fmt.Println("keySeparator is '", *keySeparator, "'")
	}

	root := treeItem{}

	for _, element := range os.Environ() {
		canAdd := true
		variable := strings.SplitN(element, "=", 2)

		envKey := variable[0]
		envValue := variable[1]

		if len(*prefix) > 0 {
			if !strings.HasPrefix(envKey, *prefix) {
				canAdd = false
			} else {
				if *removePrefix {
					envKey = envKey[len(*prefix):]
				}
			}
		}

		if canAdd {
			baseItem := &root
			envKeys := strings.Split(envKey, *keySeparator)
			totalEnvKey := len(envKeys) - 1

			for i, keyItem := range envKeys {

				if i < totalEnvKey {
					newBaseItem := getBaseItemFromKey(baseItem.firstChild, keyItem)
					if newBaseItem == nil {
						newBaseItem = &treeItem{key: keyItem}
						addChild(baseItem, newBaseItem)
					}

					baseItem = newBaseItem
				} else {
					newLeaf := &treeItem{key: keyItem, value: envValue}
					addChild(baseItem, newLeaf)
				}
			}
		}

	}

	switch *formatter {
	case "json":
		{
			fmt.Printf("{\n")
			printChildJSON(&root, 1)
			fmt.Printf("}\n")
			break
		}
	case "yaml":
		{
			printChildYAML(&root, 0)
			break
		}
	default:
		fmt.Println("Unknown formatter specifed... json/yaml")
		break
	}
}
