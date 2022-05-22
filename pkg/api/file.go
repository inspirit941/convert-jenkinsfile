package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inspirit941/convert-jenkinsfile/pkg/grammar"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
	"os"
	"path/filepath"
)

// ConvertFile @Summary jenkinsFile to github-action.yaml
// @Tags api
// @Description jenkinsFile to github-action.yaml
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "jenkinsFile"
// @Router /upload [POST]
// @Success 200 {object} gin.H{message=string,result=string} "StatusOK"
// @Failure 400 {object} gin.H{error=string} "StatusBadRequest"
func ConvertFile(c *gin.Context) {
	// File Upload
	file, _ := c.FormFile("file")
	uploadPath := filepath.Join(os.TempDir(), file.Filename)
	defer os.Remove(uploadPath)

	c.SaveUploadedFile(file, uploadPath)
	model, err := grammar.ParseJenkinsfileInDirectory(uploadPath)
	// jenkinsfile 포맷이 아닌 경우
	if err != nil {
		// todo: 에러메시지 구체화
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	asYaml, convertIssues, err := model.ToYaml()
	// 변환에 실패한 경우
	if err != nil {
		fmt.Println("Error converting to Yaml: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var convertIssuesMsg string
	if convertIssues {
		convertIssuesMsg = fmt.Sprintf("ATTENTION: Some contents of the Jenkinsfile could not be converted. Please review the github-action.yml for more information.")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": convertIssuesMsg,
		"result":  asYaml,
	})
}
