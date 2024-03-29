package keys

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type KeyLoader interface {
	Load(path string) (*rsa.PrivateKey, error)
}

type EncryptionKey struct {
	PrivateKey   *rsa.PrivateKey
	KeyName      string
	KeyID        string
	Revoked      bool
	InvalidAfter *time.Time
}

type KeySet []EncryptionKey

// RemoteLoader load from block storage
type RemoteLoader struct{}

func NewRemoteLoader() KeyLoader {
	return &RemoteLoader{}
}

func (r *RemoteLoader) Load(path string) (*rsa.PrivateKey, error) {
	key := viper.GetString("spaces_key")
	secret := viper.GetString("spaces_secret")
	s3Config := &aws.Config{
		Credentials:                    credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:                       aws.String(viper.GetString("spaces_endpoint")),
		Region:                         aws.String(viper.GetString("spaces_region")),
		DisableRestProtocolURICleaning: aws.Bool(true),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create new S3 client session")
		return nil, err
	}

	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(viper.GetString("keys_bucket")),
		Key:    aws.String(path),
	}

	log.Info().
		Str("key_path", *getObjectInput.Key).
		Str("bucket", *getObjectInput.Bucket).
		Msg("Loading key")

	result := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(newSession)
	_, err = downloader.Download(result, getObjectInput)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to read private key object")
		return nil, err
	}

	rsaKey, err := decodeRSAKey(result.Bytes())
	return rsaKey, err
}

// FSLoader Filesystem key loader
type FSLoader struct{}

func NewFSLoader() KeyLoader {
	return &FSLoader{}
}

func (s *FSLoader) Load(path string) (*rsa.PrivateKey, error) {
	if strings.HasPrefix(path, "file://") {
		path = strings.TrimPrefix(path, "file://")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Key file not found %s", path))
	}

	rawKey, err := os.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Unable to read key: %s", path))
		return nil, err
	}

	return decodeRSAKey(rawKey)
}

func GetLoader(loaderName string) KeyLoader {
	if loaderName == "block" {
		return NewRemoteLoader()
	}

	return NewFSLoader()
}

func decodeRSAKey(rawKey []byte) (*rsa.PrivateKey, error) {
	privateKeyPem, _ := pem.Decode(rawKey)
	return x509.ParsePKCS1PrivateKey(privateKeyPem.Bytes)
}
