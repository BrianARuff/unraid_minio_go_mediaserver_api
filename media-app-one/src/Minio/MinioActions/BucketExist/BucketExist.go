package bucketexist

import (
	"context"
	"fmt"
	"log"
	minioClient "main/src/Minio/MinioClient"
)

func BucketExist(bucketName string) {
	minioClient := minioClient.Minioclient()

	hasBucket, foundErr := minioClient.BucketExists(context.Background(), bucketName)

	if foundErr != nil {
		log.Fatalln("Error finding bucket: ", foundErr)
		return
	}

	fmt.Println("Found by name "+bucketName+": ", hasBucket)
}
