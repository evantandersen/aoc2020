package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	return lines
}

func toInts(strs []string) []int {
	out := make([]int, len(strs))
	for i,str := range strs {
		val,err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		out[i] = int(val)
	}
	return out
}

func adventDay1A(path string) {
	strs := readLines(path)
	ints := toInts(strs)
	for i := range ints {
		for j := i+1; j < len(ints); j++ {
			if ints[i] + ints[j] == 2020 {
				fmt.Printf("%d\n", ints[i]*ints[j])
			}
		}
	}
}

func adventDay1B(path string) {
	strs := readLines(path)
	ints := toInts(strs)
	for i := range ints {
		for j := i+1; j < len(ints); j++ {
			for k := range ints {
				if k == i || k == j {
					continue
				}
				if ints[i] + ints[j] + ints[k] == 2020 {
					fmt.Printf("%d\n", ints[i]*ints[j]*ints[k])
				}
			}
		}
	}
}

func day2Parse(s string) (min, max int, letter rune, password string) {
	a := strings.SplitN(s, "-", 2)
	min,_ = strconv.Atoi(a[0])
	a = strings.SplitN(a[1], " ", 2)
	max,_ = strconv.Atoi(a[0])
	//fmt.Printf("%s\n", a[1])
	a = strings.SplitN(a[1], ":", 2)
	//fmt.Printf("%s\n", a[0])
	letter = []rune(a[0])[0]
	password = a[1][1:]

	return
}


func adventDay2A(path string) {
	strs := readLines(path)
	numValid := 0
	for _,str := range strs {
		min,max,letter,password := day2Parse(str)
		count := 0
		for _,char := range password {
			if char == letter {
				count++
			}
		}
		//fmt.Printf("count %d\n", count)
		valid := false
		if count >= min && count <= max {
			valid = true
			numValid++
		}
		fmt.Printf("%d, %d, %c, %s, %t\n", min, max, letter, password, valid)

		//fmt.Printf("numValid = %d\n", numValid)

	}
	fmt.Printf("%d valid\n", numValid)
}

func adventDay2B(path string) {
	strs := readLines(path)
	numValid := 0
	for _,str := range strs {
		min,max,letter,password := day2Parse(str)
		//fmt.Printf("count %d\n", count)
		first := rune(password[min-1])
		last := rune(password[max-1])

		if (first == letter) != (last == letter) {
			numValid++
		}

		//fmt.Printf("numValid = %d\n", numValid)

	}
	fmt.Printf("%d valid\n", numValid)
}

func checkSlope(isTree []bool, width, height, slopeX,slopeY int) int {
	xPos := 0
	treesHit := 0
	for yPos := 0; yPos < height; yPos += slopeY {
		if isTree[yPos*width + (xPos % width)] {
			treesHit++
		}
		xPos += slopeX
	}
	return treesHit
}

func adventDay3A(path string) {
	strs := readLines(path)
	width := len(strs[0])

	isTree := make([]bool, len(strs)*width)
	for i,str := range strs {
		for j,c := range str {
			isTree[i*width + j] = c == '#'
		}
	}
	treesHit := checkSlope(isTree, width, len(strs), 3, 1)
	fmt.Printf("hit %d trees\n", treesHit)
}

func adventDay3B(path string) {
	strs := readLines(path)
	width := len(strs[0])

	isTree := make([]bool, len(strs)*width)
	for i,str := range strs {
		for j,c := range str {
			isTree[i*width + j] = c == '#'
		}
	}
	a := checkSlope(isTree, width, len(strs), 1, 1)
	b := checkSlope(isTree, width, len(strs), 3, 1)
	c := checkSlope(isTree, width, len(strs), 5, 1)
	d := checkSlope(isTree, width, len(strs), 7, 1)
	e := checkSlope(isTree, width, len(strs), 1, 2)
	fmt.Printf("hit %d trees\n", a*b*c*d*e)
}


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readPassports(path string) ([]string) {
	file, err := os.Open(path)
	if err != nil {
		 panic(err)
	}
	defer file.Close()

	var passports []string
	scanner := bufio.NewScanner(file)
	passport := ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			passport += line + "\n"
		} else {
			passports = append(passports, passport)
			passport = ""
		}
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	return passports
}

func adventDay4A(path string) {
	passports := readPassports(path)

	fmt.Printf("%d passports\n", len(passports))
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	numValid := 0
	for _,passport := range passports {
		left := make(map[string]bool)
		for _,str := range required {
			left[str] = true
		}
		pairs := strings.Fields(passport)
		for _,pair := range pairs {
			vals := strings.Split(pair, ":")
			//fmt.Printf("prop %s\n", vals[0])
			//fmt.Printf("val %s\n", vals[1])
			delete(left, vals[0])
		}
		if len(left) == 0 {
			numValid++
		}
		//panic(nil)
	}
	fmt.Printf("%d valid\n", numValid)
}

func byr(in string) bool {
	val, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return false
	}
	return val >= 1920 && val <= 2002
}

func iyr(in string) bool {
	val, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return false
	}
	return val >= 2010 && val <= 2020
}

func eyr(in string) bool {
	val, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return false
	}
	return val >= 2020 && val <= 2030
}

func hgt(in string) bool {
	if strings.HasSuffix(in, "in") {
		a := strings.TrimSuffix(in, "in")
		val, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return false
		}
		return val >= 59 && val <= 76
	}
	if strings.HasSuffix(in, "cm") {
		a := strings.TrimSuffix(in, "cm")
		val, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return false
		}
		return val >= 150 && val <= 193
	}
	return false
}
func hcl(in string) bool {
	if in[0] != '#' {
		return false
	}
	if len(in) != 7 {
		return false
	}
	if len(strings.Trim(in[1:], "0123456789abcdef")) != 0 {
		return false
	}
	return true
}

var validECL = map[string]bool {
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func ecl(in string) bool {
	_,ok := validECL[in]
	return ok
}

func pid(in string) bool {
	if len(in) != 9 {
		return false
	}
	_, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func adventDay4B(path string) {
	passports := readPassports(path)

	fmt.Printf("%d passports\n", len(passports))
	required := map[string]func(string)bool{
		"byr":byr,
		"iyr":iyr,
		"eyr":eyr,
		"hgt":hgt,
		"hcl":hcl,
		"ecl":ecl,
		"pid":pid,
	}

	numValid := 0
	for _,passport := range passports {
		left := make(map[string]func(string)bool)
		for key,val := range required {
			left[key] = val
		}
		pairs := strings.Fields(passport)
		for _,pair := range pairs {
			vals := strings.Split(pair, ":")
			//fmt.Printf("prop %s\n", vals[0])
			//fmt.Printf("val %s\n", vals[1])
			f, ok := left[vals[0]]
			if !ok {
				continue
			}
			if f(vals[1]) {
				delete(left, vals[0])
			}
		}
		if len(left) == 0 {
			numValid++
		}
		//panic(nil)
	}
	fmt.Printf("%d valid\n", numValid)
}

func adventDay5A(path string) {
	strs := readLines(path)
	maxSeatID := 0
	for _,seat := range strs {
		rowMax := 127
		rowMin := 0
		for i := 0; i < 7; i++ {
			halfway := ((rowMax-rowMin)/2) + rowMin
			if seat[i] == 'F' {
				rowMax = halfway
			} else {
				rowMin = halfway+1
			}
		}
		colMin := 0
		colMax := 7
		for j := 7; j < 10; j++ {
			halfway := ((colMax-colMin)/2) + colMin

			if seat[j] == 'L' {
				colMax = halfway
			} else {
				colMin = halfway+1
			}
		}
		seatID := 8*rowMax + colMax
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Printf("max seat id: %d\n", maxSeatID)
}

func adventDay5B(path string) {
	strs := readLines(path)
	seatIDs := make([]int, len(strs))
	for k,seat := range strs {
		rowMax := 127
		rowMin := 0
		for i := 0; i < 7; i++ {
			halfway := ((rowMax-rowMin)/2) + rowMin
			if seat[i] == 'F' {
				rowMax = halfway
			} else {
				rowMin = halfway+1
			}
		}
		colMin := 0
		colMax := 7
		for j := 7; j < 10; j++ {
			halfway := ((colMax-colMin)/2) + colMin

			if seat[j] == 'L' {
				colMax = halfway
			} else {
				colMin = halfway+1
			}
		}
		seatID := 8*rowMax + colMax
		seatIDs[k] = seatID
	}
	sort.Ints(seatIDs)

	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1] == seatIDs[i]+2 {
			fmt.Printf("found missing seat: %d\n", seatIDs[i]+1)
		}
	}

}

func adventDay6A(path string) {
	strs := readPassports(path)

	count := 0
	for _,str := range strs {
		questions := make(map[rune]bool)
		stuff := strings.Fields(str)
		for _,str2 := range stuff {
			for _,char := range str2 {
				questions[char] = true
			}
		}
		count += len(questions)
	}
	fmt.Printf("count %d\n", count)
}

func adventDay6B(path string) {
	strs := readPassports(path)

	count := 0
	for _,str := range strs {
		var questions map[rune]bool
		done1 := false
		stuff := strings.Fields(str)
		for _,str2 := range stuff {
			questions2 := make(map[rune]bool)
			for _,char := range str2 {
				questions2[char] = true
			}
			if !done1 {
				questions = questions2
				done1 = true
			} else {
				for key := range questions {
					ok, _ := questions2[key]
					if !ok {
						delete(questions, key)
					}
				}
			}
		}
		count += len(questions)
	}
	fmt.Printf("count %d\n", count)
}

type Bag struct {
	name string
	contains map[string]uint
	goldLink bool
}

func parseBags(path string) map[string]*Bag {
	strs := readLines(path)
	bags := make(map[string]*Bag)
	for _,str := range strs {
		newBag := &Bag{contains: make(map[string]uint)}

		x := strings.Split(str, "contain")
		bagInfo := strings.Fields(x[0])
		newBag.name = bagInfo[0] + " " + bagInfo[1]

		rules := strings.Split(x[1], ",")
		//fmt.Printf("Bag %s has:\n", newBag.name)
		for _,rule := range rules {
			parts := strings.Fields(rule)
			if parts[0] == "no" {
				break
			}
			amount, err := strconv.ParseUint(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}
			ruleName := parts[1] + " " + parts[2]
			newBag.contains[ruleName] = uint(amount)
			if ruleName == "shiny gold" {
				newBag.goldLink = true
			}
			//fmt.Printf("\t%d x %s\n", amount, ruleName)
		}
		bags[newBag.name] = newBag
		//fmt.Printf("\n")
	}
	return bags
}

func adventDay7A(path string) {
	bags := parseBags(path)

	numChanged := 1
	for numChanged > 0 {
		numChanged = 0
		for _,bag := range bags {
			for contains := range bag.contains {
				if bags[contains].goldLink && bag.goldLink == false {
					bag.goldLink = true
					numChanged++
				}
			}
		}
	}
	numValid := 0
	for _,bag := range bags {
		if bag.goldLink {
			numValid++
			//fmt.Printf("%s could contain shiny gold\n", name)
		}
	}
	fmt.Printf("%d could contain shiny gold\n", numValid)

}


func numberOfBagsInsideBag(start *Bag, bags map[string]*Bag) uint {
	sum := uint(0)
	for name,value := range start.contains {
		sum += value * numberOfBagsInsideBag(bags[name], bags)
		sum += value
	}
	return sum
}

func adventDay7B(path string) {
	bags := parseBags(path)

	fmt.Printf("shiny gold contains %d bags\n", numberOfBagsInsideBag(bags["shiny gold"], bags))

}

type instruction struct {
	kind int
	value int
}

func adventDay8A(path string) {
	strs := readLines(path)

	visited := make(map[int]bool)
	pos := 0
	accum := 0
	for {
		_,ok := visited[pos]
		if ok {
			fmt.Printf("loop, value is %d\n", accum)
			break
		}
		a := strings.Fields(strs[pos])
		val,err := strconv.ParseInt(a[1], 10, 64)
		if err != nil {
			panic(err)
		}
		switch a[0]  {
		case "nop":
		case "acc":
			accum += int(val)
		case "jmp":
			pos += int(val)-1
		}
		visited[pos] = true
		pos++
	}



}

func runProgram(ins []string) bool {
	visited := make(map[int]bool)
	pos := 0
	accum := 0
	for {
		if pos >= len(ins) {
			fmt.Printf("value is %d\n", accum)
			return true
		}
		_,ok := visited[pos]
		if ok {
			return false
		}
		a := strings.Fields(ins[pos])
		val,err := strconv.ParseInt(a[1], 10, 64)
		if err != nil {
			panic(err)
		}
		switch a[0]  {
		case "nop":
		case "acc":
			accum += int(val)
		case "jmp":
			pos += int(val)-1
		}
		visited[pos] = true
		pos++
	}
}

func adventDay8B(path string) {
	strs := readLines(path)

	for i,str := range strs {
		orig := str
		a := strings.Fields(str)
		if a[0] == "nop" {
			strs[i] = "jmp" + orig[3:]
			if runProgram(strs) {
				return
			}
			strs[i] = orig
		} else if a[0] == "jmp" {
			strs[i] = "nop" + orig[3:]
			if runProgram(strs) {
				return
			}
			strs[i] = orig
		} else {
			continue
		}
	}

}

func adventDay9A(path string) {
	strs := readLines(path)
	ints := toInts(strs)

	fifo := make([]int, 25)

	pos := 0

	//fill it in
	for i := 0; i < len(fifo); i++ {
		fifo[pos % len(fifo)] = ints[pos]
		pos++
	}

	for pos < len(ints) {
		toCheck := ints[pos]
		foundSol := false
		out:
		for j := 0; j < len(fifo); j++ {
			for k := j+1; k < len(fifo); k++ {
				a := fifo[(pos+j) % len(fifo)]
				b := fifo[(pos+k) % len(fifo)]
 				if a + b == toCheck {
 					foundSol = true
 					break out
				}
			}
		}
		if !foundSol {
			fmt.Printf("%d can't be constructed with the previous %d entries\n", toCheck, len(fifo))
		}
		fifo[pos % len(fifo)] = ints[pos]

		pos++
	}

}

func adventDay9B(path string) {
	strs := readLines(path)
	ints := toInts(strs)

	fifo := make([]int, 25)

	pos := 0

	//fill it in
	for i := 0; i < len(fifo); i++ {
		fifo[pos % len(fifo)] = ints[pos]
		pos++
	}

	invalidVal := 0

	for pos < len(ints) {
		toCheck := ints[pos]
		foundSol := false
	out:
		for j := 0; j < len(fifo); j++ {
			for k := j+1; k < len(fifo); k++ {
				a := fifo[(pos+j) % len(fifo)]
				b := fifo[(pos+k) % len(fifo)]
				if a + b == toCheck {
					foundSol = true
					break out
				}
			}
		}
		if !foundSol {
			fmt.Printf("%d can't be constructed with the previous %d entries\n", toCheck, len(fifo))
			invalidVal = toCheck
		}
		fifo[pos % len(fifo)] = ints[pos]

		pos++
	}

	start := 0
	end := 0
	sum := 0

	for sum != invalidVal {
		if sum < invalidVal {
			sum += ints[end]
			end++
		}
		if sum > invalidVal {
			sum -= ints[start]
			start++
		}
	}
	sort.Ints(ints[start:end])
	fmt.Printf("range %d to %d, sum is %d\n", start, end, ints[start] + ints[end-1])

}

func adventDay10A(path string) {
	strs := readLines(path)
	jolts := toInts(strs)
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	deltas := make(map[int]int)

	currentJolt := 0
	for _,adapter := range jolts {
		delta := adapter - currentJolt
		deltas[delta]++
		currentJolt = adapter
	}
	fmt.Printf("%d\n", deltas[1]*deltas[3])
}

func adventDay10B(path string) {
	strs := readLines(path)
	jolts := toInts(strs)
	jolts = append(jolts, 0)
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	ways := make([]int, len(jolts))

	for i,adapter := range jolts {
		thisWays := 0
		if i >= 3 && adapter - jolts[i-3] <= 3 {
			thisWays += ways[i-3]
		}
		if i >= 2 && adapter - jolts[i-2] <= 3 {
			thisWays += ways[i-2]
		}
		if i >= 1 && adapter - jolts[i-1] <= 3 {
			thisWays += ways[i-1]
		}
		if i == 0 {
			thisWays = 1
		}
		ways[i] = thisWays
	}
	//fmt.Printf("%v\n", jolts)
	//fmt.Printf("%v\n", ways)
	fmt.Printf("%d ways\n", ways[len(ways)-1])
}

var days = []func(path string){
	adventDay1A, adventDay1B,
	adventDay2A, adventDay2B,
	adventDay3A, adventDay3B,
	adventDay4A, adventDay4B,
	adventDay5A, adventDay5B,
	adventDay6A, adventDay6B,
	adventDay7A, adventDay7B,
	adventDay8A, adventDay8B,
	adventDay9A, adventDay9B,
	adventDay10A, adventDay10B,
}

func usage() {
	fmt.Printf("usage:\n\t%s <day number>\n", os.Args[0])
}

func main() {

	flag.Parse()
	if flag.NArg() != 1 {
		usage()
		return
	}

	dayToRun, err := strconv.ParseInt(flag.Args()[0], 10, 64)
	if err != nil {
		usage()
		return
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var inputs []string
	for _, file := range files {
		filename := file.Name()
		res := strings.Split(filename, "_")
		if len(res) > 1 {
			val, err :=  strconv.ParseInt(res[0], 10, 64)
			if err != nil {
				continue
			}
			if val == dayToRun {
				inputs = append(inputs, filename)
			}
		}
	}

	fmt.Printf("Part A\n=====================\n")
	for _,filename := range inputs {
		fmt.Printf("%s:\n", strings.SplitN(filename, "_", 2)[1])
		days[(dayToRun-1)*2](filename)
		fmt.Printf("\n")
	}

	fmt.Printf("\n\nPart B\n=====================\n")
	for _,filename := range inputs {
		fmt.Printf("%s:\n", strings.SplitN(filename, "_", 2)[1])
		days[(dayToRun-1)*2 + 1](filename)
		fmt.Printf("\n")
	}




}

