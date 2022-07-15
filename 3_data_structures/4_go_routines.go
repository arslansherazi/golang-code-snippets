package main

import (
	"fmt"
	"sync"
)

func generateUrl() string {
	return "https://s3.test.amazon.com"
}

func uploadImage(wg *sync.WaitGroup) {
	defer wg.Done() // decrease count of wait group by 1
	fmt.Println("Image Uploaded to S3")
}

func saveUserToDB(imageUrl string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("user saved to Db\n" + imageUrl)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2) // pass 2 as we have 2 go routines

	imageUrl := generateUrl()
	uploadImage(wg)
	saveUserToDB(imageUrl, wg)

	wg.Wait() // wait until all go routines got executed
}
