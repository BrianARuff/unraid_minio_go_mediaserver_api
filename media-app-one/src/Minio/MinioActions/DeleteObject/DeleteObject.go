package deleteobject

import (
	"context"
	"fmt"
	"log"
	minioClient "main/src/Minio/MinioClient"

	"github.com/minio/minio-go/v7"
)

func DeleteObject(bucketName string, objectName string) {
	minioClient := minioClient.Minioclient()

	err := minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nDeleted Object!")
}
