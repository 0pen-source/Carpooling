package dao

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossClient *oss.Client
var bucketIDCards *oss.Bucket
var err error

func InitOSSClient() {
	ossClient, err = oss.New(config.OSSConfig.EndPoint, config.OSSConfig.AccessKeyId, config.OSSConfig.AccessKeySecret)
	fmt.Println("授权信息", config.OSSConfig.EndPoint, config.OSSConfig.AccessKeyId, config.OSSConfig.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}

	bucketIDCards, err = ossClient.Bucket(config.OSSConfig.BucketIDCards)
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
