package dao

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossClient *oss.Client
var bucketIDCards *oss.Bucket
var err error

func InitOSSClient() {
	ossClient, err = oss.New(Config.OSSConfig.EndPoint, Config.OSSConfig.AccessKeyId, Config.OSSConfig.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}

	bucketIDCards, err = ossClient.Bucket(Config.OSSConfig.BucketIDCards)
	if err != nil {
		fmt.Println(err)
	}

}

func PutObject(bucket *oss.Bucket, osPath, filePath string) {

	err = bucket.PutObjectFromFile(osPath, filePath)
	if err != nil {
		fmt.Println(err)
	}
}
