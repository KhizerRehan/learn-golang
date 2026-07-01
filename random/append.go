package main

import "fmt"

type tags []string

func addTags(tags tags, tag string) tags {

	// check tag is empty or ""
	if tag == "" {
		fmt.Println("Tag is empty, cannot add to tags")
		return tags
	}
	tags = append(tags, tag)
	return tags
}

func main() {
	tags := tags{}

	// Case-1

	tags = addTags(tags, "tag1")

	fmt.Println("Add Tag-1: ", tags)

	// Case-2
	tags = addTags(tags, "")
	fmt.Println("Add Tag-2: ", tags)

}
