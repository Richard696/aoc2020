package main

/**
 * This program is to solve the advent of code day 07
 * This solution is create by Tsoding
 * I am copying this program by following Tsoding stream on Twitch
 * Date: 25 Dec 2020
 * Ori Author: Tsoding
 * Author: Yee Heng
 */

// import library (auto manage by go lint if the vs code had configured)
import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// definition for type start here

// Color is the alias of string
type Color = string

// Child is to hold the value for color and count of the bag
type Child struct {
	color  Color
	Amount int
}

// Rule is to hold the color of the container and value for child class
type Rule struct {
	container Color //
	children  []Child
}

// Rules is a hashmap that contain the key value pair of color(string) and an array of Child class
type Rules = map[Color][]Child

// end of defining type

func lineToRule(line string) Rule {
	var tmp = strings.Split(line, "contain")
	// TrimRight: Returns a slice of a given string with all trailing characters contained in the cutset removed.
	var container = strings.TrimRight(strings.Split(tmp[0], "bags")[0], " ")
	children := []Child{}

	// expect numbers -> ([0-9]); expect random stuff -> (.*); "?" is referred as lazy matching
	r := regexp.MustCompile(" ([0-9]+) (.*) bags?")
	// fmt.Println("Children bags: ")
	if tmp[1] != " no other bags." {
		for _, child := range strings.Split(tmp[1], ",") {
			result := r.FindStringSubmatch(child)
			amount, err := strconv.Atoi(result[1]) // convert string to int, strconf return reult & err hence the result is assgined to amount(variable)
			if err != nil {
				panic(err)
			}
			children = append(children, Child{
				color:  result[2],
				Amount: amount,
			})
			// fmt.Printf("%#v\n", )
		}
		// %#v	a Go-syntax representation of the value
	}

	return Rule{
		container: container,
		children:  children,
	}
}

// Return a map of color of parent bag & a list of object(contain color of the bag & amount)
func rulesFromFile(filePath string) Rules {
	content, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer content.Close()

	rules := map[Color][]Child{}
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()
		rule := lineToRule(line)
		rules[rule.container] = rule.children
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rules
}

func createDotFileForRules(rules Rules, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "digraph Rules {\n")
	for key, children := range rules {
		for _, child := range children {
			fmt.Fprintf(file, "\t%#v -> %#v;\n", key, child.color)
		}
	}
	fmt.Fprintf(file, "}\n")
}

func findContainers(rules Rules, color Color) []Color {
	result := []Color{}
	for container, children := range rules {
		for _, child := range children {
			if child.color == color {
				result = append(result, container)
				break
			}
		}
	}
	return result
}

// solution for part 1; travel the grpah up
// Use binary tree to find how many type of bags can hold shiny gold bag
func countWaysV1(rules Rules, color Color) int {
	visited := map[string]bool{} // create set
	queue := list.New()          // use linked list to hold the value

	queue.PushBack(color) // add color to the queue
	for queue.Len() > 0 {
		next := queue.Front() // return the first value of the list
		// check is the element visited; only check for not visited node
		if _, ok := visited[next.Value.(Color)]; !ok {
			for _, container := range findContainers(rules, next.Value.(Color)) {
				queue.PushBack(container)
			}
		}
		visited[next.Value.(Color)] = true // mark true after visited; Implementation of BSF
		queue.Remove(next)
	}

	return len(visited) - 1
}

// travel the graph down
func countWaysV2(rules Rules, color Color) int {
	result := 0
	if children, ok := rules[color]; ok {
		for _, child := range children {
			result += child.Amount + child.Amount*countWaysV2(rules, child.color)
		}
	}
	return result
}

func solveFile(filePath string) {
	start := "shiny gold"
	rules := rulesFromFile(filePath)
	createDotFileForRules(rules, filePath+".dot")
	fmt.Printf("Part 1: %d\n", countWaysV1(rules, start))
	fmt.Printf("Part 2: %d\n", countWaysV2(rules, start))
}

func main() {
	for _, filePath := range os.Args[1:] {
		solveFile(filePath)
	}
}
