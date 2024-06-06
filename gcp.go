package main

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"

	"google.golang.org/api/option"
)

func (app *application) ConnectToGcp() *storage.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()
	options := option.WithCredentialsFile("credentials.json")
	client, err := storage.NewClient(ctx, options)

	if err != nil {
		fmt.Println(err)
	}

	return client

}

func TestBucket(bucketName string, client storage.Client) bool {
	bucket := client.Bucket(bucketName)
	if bucket != nil {
		return true
	}
	return false
}


func GetAllFileInBucket() ([]string,error) {

}

func PushTrackToBucket(*os.File, bucketName) error {


}


func CreateBucket(filename string, client storage.Client, ctx context.Context) error {




}
