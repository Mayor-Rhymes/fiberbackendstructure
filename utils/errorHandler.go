package errorHandler


import (

	"log"
)

func HandleError(e error){


	if e != nil{

		log.Fatal(e)
	} 
	


}