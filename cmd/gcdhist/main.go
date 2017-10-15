package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {

	gcdroot := os.Getenv("GCDROOT")
	if gcdroot == "" {
		gcdroot = filepath.Join(os.Getenv("HOME"), ".gcd")
	}

	rootInfo, err := os.Stat(gcdroot)
	if err != nil {
		err := os.Mkdir(gcdroot, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	} else if !rootInfo.IsDir() {
		log.Fatalln("$GCDROOT is not a directory")
	}

	gcdHistFile := filepath.Join(gcdroot, ".gcdhist")

	args := os.Args
	if len(args) == 1 {
		args = append(args, "list")
	}

	switch args[1] {

	case "add":
		if len(args) == 2 {
			args = append(args, os.Getenv("HOME"))
		}

		if args[2] == "-" {
			hist, err := list(gcdHistFile)
			if err != nil {
				log.Fatalln(err)
			}
			args[2] = hist[len(hist)-1]
			return
		}

		f, err := os.OpenFile(gcdHistFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		absDir, err := filepath.Abs(args[2])
		if err != nil {
			log.Fatalln(err)
			return
		}

		dirInfo, err := os.Stat(absDir)
		if err != nil {
			log.Println(err)
			return
		} else if !dirInfo.IsDir() {
			log.Printf("%s is not a directory\n", absDir)
			return
		}

		fmt.Fprintln(f, absDir)
		return

	case "list":
		hist, err := list(gcdHistFile)
		if err != nil {
			log.Fatalln(err)
		}
		for _, h := range hist {
			fmt.Println(h)
		}
		return

	case "latest":
		hist, err := list(gcdHistFile)
		if err != nil {
			log.Fatalln(err)
		}

		printed := make([]string, 0)
		for i := len(hist) - 1; i >= 0; i-- {

			isDuplicate := false
			for j := range printed {
				if hist[i] == printed[j] {
					isDuplicate = true
				}
			}

			if isDuplicate {
				continue
			}

			fmt.Println(hist[i])
			printed = append(printed, hist[i])
		}
		return

	case "frequent":
		hist, err := list(gcdHistFile)
		if err != nil {
			log.Fatalln(err)
		}

		fList := make(frequencyList, 0)

		for _, h := range hist {
			hExists := false
			for i := range fList {
				if h == fList[i].dir {
					fList[i].frequency++
					hExists = true
					break
				}
			}
			if !hExists {
				fList = append(fList, frequency{
					dir:       h,
					frequency: 1,
				})
			}
		}

		sort.Sort(fList)

		for _, f := range fList {
			if f.frequency > 1 {
				fmt.Println(f.dir)
			}
		}
		return

	default:
		fmt.Println("Command invalid")
		return
	}
}

func list(gcdHistFile string) (hist []string, err error) {
	f, err := os.Open(gcdHistFile)
	if err != nil {
		return hist, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		hist = append(hist, str)
	}
	return hist, nil
}

type frequencyList []frequency

type frequency struct {
	dir       string
	frequency int
}

func (l frequencyList) Len() int {
	return len(l)
}

func (l frequencyList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l frequencyList) Less(i, j int) bool {
	return l[i].frequency > l[j].frequency
}
