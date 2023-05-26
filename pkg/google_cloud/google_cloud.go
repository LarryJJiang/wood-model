package google_cloud

import (
	documentai "cloud.google.com/go/documentai/apiv1"
	"cloud.google.com/go/documentai/apiv1/documentaipb"
	"context"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"regexp"
	"strings"
	"wood/pkg/logging"
	"wood/pkg/setting"
)

type FileInfo struct {
	FileAddress string
	MimeType    string
}

var ProjectId *string
var Location *string
var ProcessorID *string

func init() {
	ProjectId = &setting.GoogleCloudSetting.ProjectId
	Location = &setting.GoogleCloudSetting.Location
	ProcessorID = &setting.GoogleCloudSetting.ProcessorId
}

// 识别文件内容
func IdentifyFileHandle(file *FileInfo) (string, error) {
	filePath := &file.FileAddress
	mimeType := &file.MimeType
	flag.Parse()

	ctx := context.Background()
	endpoint := fmt.Sprintf("%s-documentai.googleapis.com:443", *Location)
	client, err := documentai.NewDocumentProcessorClient(ctx, option.WithEndpoint(endpoint))
	if err != nil {
		fmt.Println(fmt.Errorf("error creating Document AI client: %w", err))
		return "", err
	}
	defer client.Close()

	// Open local file.
	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println(fmt.Errorf("ioutil.ReadFile: %w", err))
		return "", err
	}
	req := &documentaipb.ProcessRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/processors/%s", *ProjectId, *Location, *ProcessorID),
		Source: &documentaipb.ProcessRequest_RawDocument{
			RawDocument: &documentaipb.RawDocument{
				Content:  data,
				MimeType: *mimeType,
			},
		},
	}
	resp, err := client.ProcessDocument(ctx, req)
	if err != nil {
		fmt.Println(fmt.Errorf("processDocument: %w", err))
		logging.Error(fmt.Sprintf("参数解析错误，对应的值不是布尔类型：%v\n", err))
		return "", err
	}

	// Handle the results.
	document := resp.GetDocument()
	result := document.GetText()
	fmt.Printf("Document Text: %v\n", result)
	return result, nil
}

func IdentifyFile(file *FileInfo) (map[string]interface{}, error) {
	result, err := IdentifyFileHandle(file)
	if err != nil {
		return nil, err
	}
	return DocumentFormat(result), nil
}

func DocumentFormat(document string) map[string]interface{} {
	arr := strings.Split(document, "\n")
	length := len(arr)
	var newSlice []string
	for i := 0; i < length; i++ {
		arr[i] = strings.Trim(arr[i], " ")
		arr[i] = strings.Trim(arr[i], ":")
		arr[i] = strings.Trim(arr[i], " ")
		//reg := regexp.MustCompile(`[0-9]+`)
		//match1 := reg.FindAllString(arr[i], -1)
		//fmt.Println(arr[i], "匹配数字：", match1)
		if strings.Contains(arr[i], ":") {
			reg := regexp.MustCompile(`\s*[0-9]+\:[0-9]+\:[0-9]+\s*[am|pm]{1,2}\s*[0-9]{1,2}-[a-zA-Z]{3}\s*`)
			match := reg.FindAllString(arr[i], -1)
			fmt.Println("匹配：", match)
			if len(match) > 0 {
				arr[i] = strings.Replace(arr[i], match[0], "", 1)
				fmt.Println("替换后：", arr[i])
				if strings.Contains(arr[i], ":") {
					subArr := strings.Split(arr[i], ":")
					newSlice = append(newSlice, subArr...)
				} else {
					newSlice = append(newSlice, arr[i])
				}
			} else {
				if strings.Contains(arr[i], ":") {
					subArr := strings.Split(arr[i], ":")
					newSlice = append(newSlice, subArr...)
				} else {
					newSlice = append(newSlice, arr[i])
				}
			}
		} else {
			if arr[i] == "" {
				continue
			}
			if strings.Contains(arr[i], "Carrier") {
				newSlice = append(newSlice, strings.Split(arr[i], " ")...)
			} else {
				newSlice = append(newSlice, arr[i])
			}
		}
	}
	field := make(map[string]interface{}, 16)
	field = map[string]interface{}{
		"Seq Num":     1,
		"Vehicle":     3,
		"Num Plate":   5,
		"Customer":    11,
		"Product":     12,
		"Forest":      13,
		"Grade":       14,
		"Docket":      18,
		"Destination": 19,
		"Carrier":     22,
		"Gross":       23,
		"Tare":        24,
		"Net Wgt":     25,
		"Gross2":      28,
		"Tare2":       30,
		"Net":         32,
	}
	for key, value := range field {
		sliceKey := value.(int)
		newSlice[sliceKey] = strings.Trim(newSlice[sliceKey], " ")
		fmt.Printf("键：%v 值：%v slice值：%v\n", key, value, newSlice[sliceKey])
		field[key] = newSlice[value.(int)]
	}
	return field
}
