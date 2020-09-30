package main

/*
  #include "volcasample/syro/korg_syro_comp.c"
  #include "volcasample/syro/korg_syro_comp.h"
  #include "volcasample/syro/korg_syro_func.c"
  #include "volcasample/syro/korg_syro_func.h"
  #include "volcasample/syro/korg_syro_type.h"
  #include "volcasample/syro/korg_syro_volcasample.c"
  #include "volcasample/syro/korg_syro_volcasample.h"
*/
import "C"

import (
	"fmt"
)

func main () {
	fmt.Println(C.VOLCA_SAMPLE_FS)
}
