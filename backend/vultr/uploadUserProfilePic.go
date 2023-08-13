package vultr

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func UploadUserProfilePic(file *os.File) (string, error) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		return "", fmt.Errorf("could not load config: %v", err)
	}

	reg := config.VultrRegion

	profileFile, err := os.Open(file.Name())
	if err != nil {
		return "", fmt.Errorf("could not open profile pic: %v", err)
	}

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.VultrAccessKey, config.VultrSecretKey, ""),
		Endpoint:         aws.String(fmt.Sprintf("https://%s", config.VultrHostname)),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(reg),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return "", fmt.Errorf("could not create new session: %v", err)
	}

	s3Client := s3.New(newSession)

	s3InputObject := s3.PutObjectInput{
		Bucket: aws.String("/user-profile-pictures"),
		Key:    aws.String(uuid.New().String()),
		Body:   profileFile,
		ACL:    aws.String("public-read"),
	}

	_, err = s3Client.PutObject(&s3InputObject)
	if err != nil {
		return "", fmt.Errorf("could not upload user profile pic: %v", err)
	}

	vultrProfileUrl := fmt.Sprintf("https://%s.vultrobjects.com/user-profile-pictures/%s", reg, *s3InputObject.Key)

	return vultrProfileUrl, nil
}
