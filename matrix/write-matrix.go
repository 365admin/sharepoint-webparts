/*---
title: Write Matrix
description: Write the matrix to a file
output: matrix.json
---
*/

package matrix

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

func WriteMatrix() (filePath string, e error) {
	data, err := AnalysePages()

	if err != nil {
		return filePath, err
	}
	filePath = path.Join(viper.GetString("WORKDIR"), "matrix.json")

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
		return filePath, err
	}
	err = os.WriteFile(filePath, file, 0644)
	return filePath, err
}
