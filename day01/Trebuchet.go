package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func part01( file string) int{
    f,err := os.Open(file)
  if err !=nil{
    log.Fatal(err)
  }

  defer f.Close()

  scanner :=bufio.NewScanner(f)

  i:=0

  var sum int
  for scanner.Scan(){

    w:=make([]string,i+1)
    n:=make([]string,i+1)
    w[i]=scanner.Text()
    for _,c := range w[i]{
      if unicode.IsDigit(c){
        n[i]+=string(c)
      }
    }
    if len(n[i])>2{
      s:=n[i]
      n[i]=string(s[0])+string(s[len(n[i])-1])

    }else if len(n[i])==1{
      n[i]=n[i]+n[i]
    }

    int,err:=strconv.Atoi(n[i])

    if err !=nil {
      continue
    }

   sum+=int
    i++
  }

  return sum

}

func part02 (file string) (string, error) {

  f,err :=os.Open(file)

  if err !=nil{
    log.Fatal(err)
  }
  defer f.Close()
  
  i:=0
    w:=make([]string,i+1)
  scanner := bufio.NewScanner(f)
  for scanner.Scan(){
    if len(scanner.Text())!=0{

w=append(w, scanner.Text())
    }
  }
rx := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)

	total := 0

	for _,line := range w {
		nums := []string{}


		for j := 0; j <= len(line); j++ {
			n := rx.FindString(line[j:])
			if n == "" {
				break
			}
			nums = append(nums, n)
		}

		for i, n := range nums {
			switch n {
			case "one":
				nums[i] = "1"
			case "two":
				nums[i] = "2"
			case "three":
				nums[i] = "3"
			case "four":
				nums[i] = "4"
			case "five":
				nums[i] = "5"
			case "six":
				nums[i] = "6"
			case "seven":
				nums[i] = "7"
			case "eight":
				nums[i] = "8"
			case "nine":
				nums[i] = "9"
			default:
			}
		}
   if len(nums)!=0{

		left := nums[0]
		right := nums[len(nums)-1]
		calib, err := strconv.Atoi(fmt.Sprintf("%s%s", left, right))
		if err != nil {
			return "", err
		}

		total += calib
    } 


	}

	return strconv.Itoa(total), nil
}


 func main(){
  fmt.Println(part01("./input01.txt"))
  fmt.Println(part02("./input02Test.txt"))

} 
