package initEnvironmentVariables

import "os"

func UnraidCaPrivateKeyMediaAppOne() string {
	unraidCaPrivateKeyMediaAppOne := os.Getenv("UNRAID_CA_PRIVATE_KEY_MEDIA_APP_ONE")

	return unraidCaPrivateKeyMediaAppOne
}

func UnraidCaServerPemMediaAppOne() string {
	unraidCaServerPemMediaAppOne := os.Getenv("UNRAID_CA_SERVER_PEM_MEDIA_APP_ONE")

	return unraidCaServerPemMediaAppOne
}

func UnraidMinioExternalServerAddress() string {
	unraidMinioExternalServerAddress := os.Getenv("UNRAID_MINIO_EXTERNAL_SERVER_ADDRESS")

	return unraidMinioExternalServerAddress
}

func UnraidServerMinioAccessKeyIdMediaAppOne() string {
	unraidServerMinioAccessKeyIdMediaAppOne := os.Getenv("UNRAID_SERVER_MINIO_ACCESS_KEY_ID_MEDIA_APP_ONE")

	return unraidServerMinioAccessKeyIdMediaAppOne
}

func UnraidServerMinioSecretAccessKeyMediaAppOne() string {
	unraidServerMinioSecretAccessKeyMediaAppOne := os.Getenv("UNRAID_SERVER_MINIO_SECRET_ACCESS_KEY_MEDIA_APP_ONE")

	return unraidServerMinioSecretAccessKeyMediaAppOne
}

func UnraidServerMinioTokenMediaAppOne() string {
	unraidServerMinioTokenMediaAppOne := os.Getenv("UNRAID_SERVER_MINIO_TOKEN_MEDIA_APP_ONE")

	return unraidServerMinioTokenMediaAppOne
}

func UnraidServerMinioRegionMediaAppOne() string {
	unraidServerMinioRegionMediaAppOne := os.Getenv("UNRAID_SERVER_MINIO_REGION_MEDIA_APP_ONE")

	return unraidServerMinioRegionMediaAppOne
}
