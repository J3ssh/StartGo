package main //declare main package

import ( //declare imports
  "fmt" //format package 
)

func main() { //delcare main function 
  for i := 1; i <= 1024; i++ { //create for loop to itterate numbers in a rage
    address := fmt.Sprintf ("scanme.nmap.org:%d", i) //append %d loop output to end of address string
    fmt.Println(address) //print address var
   }
}
