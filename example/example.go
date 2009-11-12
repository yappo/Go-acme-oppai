package main

import "fmt"
import "acme/oppai"

func main() {
     var o oppai.Oppai;
     fmt.Printf(
		 o.Massage().
			 Massage().
			 Massage().
			 Oppai("").
			 Oppai("").
			 Oppai("").
			 Massage().
			 Done()
	 );
}
