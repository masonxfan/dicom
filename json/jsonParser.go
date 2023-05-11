package main

import (
	"encoding/json"
	"fmt"
)

var sample = []byte(`{
		   "0020000D": {
			"vr": "UI",
			"Value": [ "1.2.392.200036.9116.2.2.2.1762893313.1029997326.945873" ]
		  }
		}
		{	
		  "0020000D" : {
			"vr": "UI",
			"Value": [ "1.2.392.200036.9116.2.2.2.2162893313.1029997326.945876" ]
		  }
		}`)

type Dicom struct {
	StudyId string `json:"0020000D"`
}

func main() {
	var dicom Dicom
	json.Unmarshal(sample, &dicom)
	fmt.Printf("Description: %s", dicom.StudyId)
}
