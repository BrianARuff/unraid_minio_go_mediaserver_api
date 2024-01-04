package getfile

import (
	"context"
	"fmt"
	minioClient "main/src/Minio/MinioClient"

	"github.com/minio/minio-go/v7"
)

func GetFile(bucketName string, objectName string, filePath string) {
	minioClient := minioClient.Minioclient()

	fgeterr := minioClient.FGetObject(context.Background(), bucketName, objectName, filePath, minio.GetObjectOptions{})

	if fgeterr != nil {
		fmt.Println(fgeterr)
		return
	}

	fmt.Println("\nGot File!")
}
