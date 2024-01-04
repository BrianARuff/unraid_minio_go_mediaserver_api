package listobjects

import (
	"context"
	"fmt"
	minioClient "main/src/Minio/MinioClient"

	"github.com/minio/minio-go/v7"
)

func ListObjects(bucketName string) {
	minioClient := minioClient.Minioclient()

	oChan := minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{Recursive: true})

	for object := range oChan {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println("Object key: ", object.Key)
	}

	fmt.Println("\nListed All Objects!")
}
