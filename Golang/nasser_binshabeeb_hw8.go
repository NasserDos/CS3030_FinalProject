

package main

import (
		"fmt" //to print in proper format
        "os" //to create directories
        "os/exec" //to scp
        "time" //for time
        "flag" //getopts
	)

    func copyFile(c string, f string){

	    ourTime := time.Now().String()
        var stamp  = ourTime[0:10]
        var month = c+"/"+stamp[5:7]
        var dest = month+"/FRED.csv."+stamp
        var server = "nb19445@icarus.cs.weber.edu"
        var fileAt = ":/home/hvalle/submit/cs3030/files/"+f
        var target = server+fileAt

        output, err := exec.Command("scp",target,dest).CombinedOutput()
        if err != nil {
              fmt.Println("Failed to copy file")
              fmt.Println("File doesn't exist !")
              os.Exit(1)
          }
          fmt.Printf(string(output))
          fmt.Println("File transfer complete")
          fmt.Println("File Located in [",dest,"]")

    }

    func prepStructure(){

        fmt.Printf("Checking folder structure\n")
		fmt.Printf("Creating file structure...\n")

        os.Mkdir("fredData",os.ModePerm)

        for i:= 1; i< 13; i++ {
                dig := fmt.Sprintf("%02d",i)
                os.Mkdir("fredData/"+dig,os.ModePerm)
        }
		fmt.Printf("File stucture build complete...\n")

    }

    func usage(){
        fmt.Printf("usage : go run nasser_binshabeeb_hw8.go \n")
        fmt.Printf("options -c fredData -f FRED.csv \n")
        fmt.Printf("Both fields are required \n")
    }


	func main(){
        //basically getopts
        Cust := flag.String("c","","Customer folder")
        Data := flag.String("f","","Data File")
        Help := flag.Bool("help",false,"Help")
        flag.Parse()


        if *Cust == "" || *Data == "" || *Help {

            usage()
            os.Exit(1)
        } else{
            prepStructure()
            copyFile(*Cust,*Data)
        }
      }

