package auth

import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
)

var (
	ErrorPassword       = errors.New("invalid password")
	ErrorToken          = errors.New("invalid token")
	ErrorOtpExpiry      = errors.New("otp expired")
	ErrorConvertTime    = errors.New("could not convert interface to time.Time")
	ErrorMemberLogin    = errors.New("select any one of the config for member login")
	ErrorUnauthorized   = errors.New("Unauthorized")
	ErrorInactive       = errors.New("user disabled please contact admin")
	ErrorInvalidOTP     = errors.New("invalid OTP")
	ErrorInactiveMember = errors.New("inactive member")
	ErrorTokenExpiry    = errors.New("token expired")
)

type Action string

const (
	Create Action = "Create"

	Read Action = "View"

	Update Action = "Update"

	Delete Action = "Delete"

	CRUD Action = "CRUD"
)

var (
	AWSID     string
	AWSKEY    string
	AWSREGION string
	AWSBUCKET string
)

func GetSelectedType(Db *gorm.DB) (storageType TblStorageType, err error) {

	storageType, err = Authmodel.GetStorageValue(Db)
	if err != nil {

		return TblStorageType{}, err

	}

	return storageType, nil
}

func SetS3value(Db *gorm.DB) error {

	storagetype, err := GetSelectedType(Db)
	if err != nil {
		return err
	}

	if storagetype.Aws != nil {

		AWSID = storagetype.Aws["AccessId"].(string)

		AWSKEY = storagetype.Aws["AccessKey"].(string)

		AWSREGION = storagetype.Aws["Region"].(string)

		AWSBUCKET = storagetype.Aws["BucketName"].(string)

	}

	return nil

}

func CreateS3Session(Db *gorm.DB) (ses *s3.S3, err error) {

	SetS3value(Db)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWSREGION),
		Credentials: credentials.NewStaticCredentials(AWSID, AWSKEY, ""),
	})

	if err != nil {

		log.Println("Error creating session: ", err)

		return nil, err

	}

	svc := s3.New(sess)

	return svc, nil

}

func CreateFolderToS3(foldername string, folderpath string, auth *Auth) (folderPath string, err error) {

	if foldername != "" {

		svc, _ := CreateS3Session(auth.DB)

		// fmt.Println("inside folder create s3", folderpath+foldername)

		put, err := svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(AWSBUCKET),
			Key:    aws.String(folderpath + foldername),
			Body:   bytes.NewReader(nil),
		})

		if err != nil {
			return "", fmt.Errorf("failed to create folder, %v", err)
		}

		fmt.Printf("create folder to, %s\n", put)

		var s3Path = folderPath + foldername + "/"

		return s3Path, nil

	}

	return "", errors.New("foldername is empty can't create")
}

