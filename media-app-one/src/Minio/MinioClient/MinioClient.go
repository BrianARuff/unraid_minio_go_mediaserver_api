package minioClient

import (
	"log"
	"net/http/httptrace"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Minioclient() *minio.Client {
	// TODO: SET UP ENVIRONMENT VARIABLES
	// godotenv.Read(".env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	minioServerndpoint := os.Getenv("UNRAID_MINIO_EXTERNAL_SERVER_ADDRESS")
	minioBucketAccessKeyID := os.Getenv("UNRAID_SERVER_MINIO_ACCESS_KEY_ID_MEDIA_APP_ONE")
	minioBucketSecretAccessKey := os.Getenv("UNRAID_SERVER_MINIO_SECRET_ACCESS_KEY_MEDIA_APP_ONE")
	minioBucketToken := os.Getenv("UNRAID_SERVER_MINIO_TOKEN_MEDIA_APP_ONE")
	minioBucketRegion := os.Getenv("UNRAID_SERVER_MINIO_REGION_MEDIA_APP_ONE")
	// // minioServerndpoint := "192.168.1.195:9768"
	// minioBucketAccessKeyID := environmentVariables.UnraidServerMinioAccessKeyIdMediaAppOne()
	// // minioBucketAccessKeyID := "VhlqAPCIpFKHAM08TZOE"
	// minioBucketSecretAccessKey := environmentVariables.UnraidServerMinioSecretAccessKeyMediaAppOne()
	// // minioBucketSecretAccessKey := "9dWtP44Iua54t9fxt7xBfnUpbSMfVS4kYAkrbyBh"
	// // minioBucketSecret := ""
	// minioBucketSecret := envriomentVariables.UnraidCaPrivateKeyMediaAppOne()
	// minioBucketAccessToken := environmentVariables.UnraidServerMinioTokenMediaAppOne()
	// // minioBucketRegion := "us-east-1"

	minioClient, err := minio.New(minioServerndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioBucketAccessKeyID, minioBucketSecretAccessKey, minioBucketToken),
		Region: minioBucketRegion,
		// Secure: true,
		Trace: &httptrace.ClientTrace{
			ConnectStart: func(network, addr string) {
				log.Printf("Connecting to %s\n", addr)
			},
			ConnectDone: func(network, addr string, err error) {
				log.Printf("Connected to %s\n", addr)
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}
