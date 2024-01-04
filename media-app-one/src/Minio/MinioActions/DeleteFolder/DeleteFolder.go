package deletefolder

import (
	"context"
	"fmt"
	minioClient "main/src/Minio/MinioClient"

	"github.com/minio/minio-go/v7"
)

func DeleteFolder(bucketName string, folderName string) error {
	minioClient := minioClient.Minioclient()

	ctx := context.Background()

	if folderName != "" && folderName[len(folderName)-1] != '/' {
		folderName += "/"
	}

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix:    folderName,
		Recursive: true,
	})

	objectsToDelete := make([]minio.ObjectInfo, 0)
	for object := range objectCh {
		if object.Err != nil {
			return object.Err
		}
		objectsToDelete = append(objectsToDelete, object)
	}

	for _, obj := range objectsToDelete {
		err := minioClient.RemoveObject(ctx, bucketName, obj.Key, minio.RemoveObjectOptions{})
		if err != nil {
			return err
		}
	}

	fmt.Println("Deleted Folder " + folderName + " and its contents successfully")

	return nil
}
