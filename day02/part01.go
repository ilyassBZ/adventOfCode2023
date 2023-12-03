package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)
func part01(file string) (int,error) {
  f,err:=os.Open(file)
  if err!=nil{
    log.Fatal(err)
  }
  defer f.Close()
  scanner:=bufio.NewScanner(f)

  idc:=0
  sum_id_cubes:=0
  var cubes =map[string]int{
    "red":12,
    "green":13,
    "blue":14,
  }
  for scanner.Scan(){
  posible:=true
  idc++

    games:=strings.Split(scanner.Text(), ":")

    part:=strings.Split(games[1], ";")

    for i:=0;i<len(part);i++{
      par:=strings.Split(part[i],",")
      for _,l:=range par{
        c:=strings.Split(l," ")
        n,err:=strconv.Atoi(c[1])
        if err!=nil{
          log.Fatal(err)
        }
      if n > cubes[c[2]]{
        posible=false
      }
      }
    }
  if posible{
    sum_id_cubes+=idc
  }
}

return sum_id_cubes,nil
}

func main() {
  fmt.Println(part01("./input.txt"))
}
