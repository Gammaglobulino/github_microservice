package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJSON(t *testing.T) {
	request := CreateRepoRequest{
		Name:         "golang intro",
		Description:  "a golang intro repository",
		Homepage:     "https://github.com",
		Private:      true,
		Has_issue:    true,
		Has_projects: true,
		Has_wiki:     true,
	}
	bytes, err := json.Marshal(request)
	var jsonBK CreateRepoRequest
	err = json.Unmarshal(bytes, &jsonBK)
	assert.Nil(t, err)
	assert.EqualValues(t, jsonBK.Name, request.Name)
}
