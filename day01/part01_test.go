package main

import (
	"testing"
)



func TestPart01 (t *testing.T){
  got :=part01("./input01test.txt") 
  want :=142

  if got !=want{
    t.Errorf("got %d, wanted %d",got,want)
  }

}

