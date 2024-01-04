package addfolder

import (
	"context"
	"fmt"
	"log"
	minioClient "main/src/Minio/MinioClient"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
)

// TODO: Make Progress and related code into a utility function
type Progress struct {
	file         *os.File
	totalSize    int64
	uploadedSize int64
}

func NewProgress(file *os.File) (*Progress, error) {
	fileInfo, err := file.Stat()

	if err != nil {
		return nil, err
	}

	return &Progress{
		file:      file,
		totalSize: fileInfo.Size(),
	}, nil
}

func (pb *Progress) Read(p []byte) (int, error) {
	n, err := pb.file.Read(p)
	pb.uploadedSize += int64(n)

	percentage := float64(pb.uploadedSize) / float64(pb.totalSize) * 100
	fmt.Printf("\rUploading: %s [%d%%]", pb.file.Name(), int(percentage))

	return n, err
}

func UploadFolder(bucketName string, localFolderPath string, remoteFolderPath string) error {
	minioClient := minioClient.Minioclient()

	ctx := context.Background()

	return filepath.Walk(localFolderPath, func(localPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(localPath)
		if err != nil {
			return err
		}
		defer file.Close()

		progress, err := NewProgress(file)

		if err != nil {
			log.Fatalln(err)
		}

		relativePath, _ := filepath.Rel(localFolderPath, localPath)

		remotePath := filepath.Join(remoteFolderPath, relativePath)
		remotePath = strings.ReplaceAll(remotePath, "\\", "/")

		_, err = minioClient.PutObject(ctx, bucketName, remotePath, file, info.Size(), minio.PutObjectOptions{PartSize: 1024 * 1024 * 10, NumThreads: 10, ConcurrentStreamParts: true, Progress: progress, ContentEncoding: "gzip"})
		if err != nil {
			return err
		}

		fmt.Println("Successfully uploaded folder "+remoteFolderPath+" and its contents", remotePath)

		return nil
	})
}
