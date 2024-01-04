package putfile

import (
	"context"
	"fmt"
	"log"
	minioClient "main/src/Minio/MinioClient"
	"os"

	"github.com/minio/minio-go/v7"
)

// TODO: Make Progress a utility function
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

func FPutFile(bucketname string, objectName string, filePath string) {
	minioClient := minioClient.Minioclient()

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	progress, err := NewProgress(file)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = minioClient.FPutObject(context.Background(), bucketname, objectName, filePath, minio.PutObjectOptions{ContentType: "application/octet-stream", Progress: progress, ContentEncoding: "gzip", PartSize: 1024 * 1024 * 10, NumThreads: 10, ConcurrentStreamParts: true})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nUpload File!")
}
