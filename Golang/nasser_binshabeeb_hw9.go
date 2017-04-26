//this runs the main function
package main

import (
		"fmt" //to print in proper format
        "os" //to create directories
        "os/exec" //to scp
        "time" //for time
        "flag" //getopts
	)

    //This function copies the file from icarus
    //it takes two parameters 
    // c=> customer folder
    // f=> dataFile
    //it also gets the time using the time module
    //It then runs a command to scp using all the parameters
    //If the scp fails, it shows an error message
        func copyFile(c string, f string){

        //get the time
        //then setup the variables
	    ourTime := time.Now().String()
        var stamp  = ourTime[0:10] //use slicing
        var month = c+"/"+stamp[5:7] //concatenation
        var dest = month+"/FRED.csv."+stamp
        var server = "nb19445@icarus.cs.weber.edu"
        var fileAt = ":/home/hvalle/submit/cs3030/files/"+f
        var target = server+fileAt

        //double return and double assigne with double init
        output, err := exec.Command("scp",target,dest).CombinedOutput()
        if err != nil { //on scp failure
              fmt.Println("Failed to copy file")
              fmt.Println("File doesn't exist !")
              os.Exit(1)
          }
          fmt.Printf(string(output))
          fmt.Println("File transfer complete")
          fmt.Println("File Located in [",dest,"]")

    }

    //this method creates the file structure
    //loops and builds the folders as it runs
    func prepStructure(){

        fmt.Printf("Checking folder structure\n")
		fmt.Printf("Creating file structure...\n")

        os.Mkdir("fredData",os.ModePerm)

        //forloop, the only loop
        for i:= 1; i< 13; i++ {
                dig := fmt.Sprintf("%02d",i)
                os.Mkdir("fredData/"+dig,os.ModePerm)
        }
		fmt.Printf("File stucture build complete...\n")

    }

    func usage(){
        fmt.Printf("usage : ./nasser_binshabeeb_hw8 -c custFolder -f dataFile \n")
        fmt.Printf("options -c fredData -f FRED.csv \n")
        fmt.Printf("Both fields are required \n")
    }

    //main function, this will run "because of the first line in the script"
    //it uses getopts to determine if the options were passed in
    //some pointer manipulation does the trick
	func main(){
        //basically getopts
        Cust := flag.String("c","","Customer folder")
        Data := flag.String("f","","Data File")
        Help := flag.Bool("help",false,"Help")
        flag.Parse()

        //on the fly, this could be done better
        if *Cust == "" || *Data == "" || *Help {

            usage()
            os.Exit(1)
        } else{
            prepStructure()
            copyFile(*Cust,*Data)
        }
      }

