

package main

import (
		"fmt"
        "os"
	)

    func copyFile(){

    }

    func prepStructure(){

        fmt.Printf("Checking folder structure")
		fmt.Printf("Creating file structure...\n")
        os.Mkdir("fredData",os.ModePerm)
        for i:= 1; i< 13; i++ {
            dig := fmt.Sprintf("%02d",i)
            os.Mkdir("fredData/"+dig,os.ModePerm)
        }

    }

    func usage(){
        fmt.Printf("usage : go run nasser_binshabeeb_hw8.go \n")
        fmt.Printf("options -c fredData -f FRED.csv \n")
        fmt.Printf("Both fields are required \n")
    }


	func main(){
	}
