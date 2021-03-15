package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"errors"
	"strconv"
)

// Run with
//		go run main.go
// Send request with:
//	curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
fmt.Fprint(w,"Hello welcome to home page")
})	
http.HandleFunc("/echo",echoHandler)
http.HandleFunc("/sum",sumHandler)
http.HandleFunc("/multiply",multiplyHandler)
http.HandleFunc("/invert",invertHandler)
http.HandleFunc("/flatten",flattenHandler)
http.ListenAndServe(":8080", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
      records:=fileReader(w ,r)
	 if isValidMatrix(records){
		var response string
		response=formatMatrix(records)
		fmt.Fprint(w,response);
	 } else{
		fmt.Fprint(w,"Error:File should only contain numbers")
	 }
	 
}

func sumHandler(w http.ResponseWriter, r *http.Request){
	records:=fileReader(w ,r)
	if isValidMatrix(records){
		sum:=0
		for i:=0;i<len(records);i++ {
			for j:=0;j<len(records);j++ {
				number,_:=strconv.Atoi(records[i][j])
				sum+= number
			}
		}
		fmt.Fprint(w,sum)
	} else {
		fmt.Fprint(w,"Error:File should only contain numbers")
	}
}
func multiplyHandler(w http.ResponseWriter, r *http.Request){
	records:=fileReader(w ,r)
	if isValidMatrix(records){
		multiplyResult:=1
		for i:=0;i<len(records);i++ {
			for j:=0;j<len(records);j++ {
				number,_:=strconv.Atoi(records[i][j])
				multiplyResult*=number
			}
		}
		fmt.Fprint(w,multiplyResult)
	}else {
		fmt.Fprint(w,"Error:File should only contain numbers")
	}
}
func flattenHandler(w http.ResponseWriter, r *http.Request){
	records:=fileReader(w ,r)
	if isValidMatrix(records){
		var response string 
		for _, row := range records {
			response =fmt.Sprintf("%s%s", response, strings.Join(row, ",")) + ","
		}
		fmt.Fprint(w,strings.TrimSuffix(response,","));
	} else {
		fmt.Fprint(w,"Error:File should only contain numbers")
	}
  

}
func invertHandler(w http.ResponseWriter, r *http.Request){
	records:=fileReader(w ,r)
	 if isValidMatrix(records){
		length:=len(records)
		var response string
		for layer:=0;layer<length;layer++ {
			for j:=layer+1;j<length;j++ {
				temp:=records[layer][j]
				records[layer][j]=records[j][layer]
				records[j][layer]=temp;
			}
		}
		response=formatMatrix(records)
		fmt.Fprint(w,response)
	 }else {
		 fmt.Fprint(w,"Error:File should only contain numbers")
	 }

}
//read file from url
func fileReader(w http.ResponseWriter, r *http.Request) ([][]string){
	     //copy file from url
		file, _, err:= r.FormFile("file")
		if err != nil {
			fmt.Fprint(w,errors.New("Error:please enter file"))
			return nil
		}
		defer file.Close()
		records, _ := csv.NewReader(file).ReadAll()
		if len(records)==0 || len(records)!=len(records[0]){
			fmt.Fprint(w,"Error:File should be non empty and have equal number of rows and columns")
			return nil

		}
	 return records
	
}
 
//format records by removing braces
func formatMatrix(records[][] string)string{
	var response string
	for _, row := range records {
		response= fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

//checks all records are number
func isValidMatrix(records[][] string)bool{
	for i:=0;i<len(records);i++ {
		for j:=0;j<len(records);j++ {
			_,err:=strconv.Atoi(records[i][j])
            if err!=nil{
				return false
			} 
		}
	}
	return true
}


