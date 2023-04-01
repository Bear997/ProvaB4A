package utility

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func firebaseInit(ctx context.Context) *storage.BucketHandle {
	config := &firebase.Config{
		StorageBucket: "tearcard-85619.appspot.com",
	}

	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		fmt.Println(err)
	}
	client, err := app.Storage(ctx)
	if err != nil {
		fmt.Println(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		fmt.Println(err)
	}
	return bucket
}

func UploadImageToFirebaseStorage(image *multipart.FileHeader) error {
	ctx := context.Background()
	bucket := firebaseInit(ctx)
	imagePath := "tmp/" + image.Filename

	obj := bucket.Object(imagePath)

	imageR, errR := image.Open()
	if errR != nil {
		fmt.Println("sto in errore dopo image.opne")
		fmt.Println(errR)
	}

	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, imageR); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(obj.ObjectName())
	defer imageR.Close()
	defer writer.Close()

	return nil
}
