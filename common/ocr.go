package common

import (
	"encoding/base64"
	"fmt"
	"os"

	tCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

// https://github.com/esimov/caire
func GetBase64FromImage(imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	var size int64 = fileInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	_, err = file.Read(buf)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}

func getOcr() string {

	credential := tCommon.NewCredential(
		"AKIDnQnukLgkY5LJ1ScV8VbQyFyRdZEcjbCV",
		"4rNcZFfpjMinb8OO7ankQpxpQknZ5KIz",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(credential, "ap-guangzhou", cpf)

	request := ocr.NewGeneralBasicOCRRequest()
	res, err := GetBase64FromImage(`D:\src\bookbook\minimap\Screenshot_2022-06-05-09-55-29-689_com.tencent.mm.jpg`)
	if err != nil {
		fmt.Println(err)
	}
	request.ImageBase64 = tCommon.StringPtr(res)
	request.IsWords = tCommon.BoolPtr(true)

	response, err := client.GeneralBasicOCR(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return ""
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
	return response.ToJsonString()
}
