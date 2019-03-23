// Copyright © 2019 Kerem Karatal
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lib

import (
	"context"
	"fmt"
	"net/http"

	homedir "github.com/mitchellh/go-homedir"
	"golang.org/x/oauth2"
)

func GenerateChannelFolderPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return home + "/.coding-challenges", err
}

func getTokenClient(token string) *http.Client {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	return tokenClient
}

func generateChallengeRepositoryName(candidateName string, discipline string) string {
	return "test_" + discipline + "_" + candidateName
}

func generateTemplateRepositoryURL(owner string, organization string, templateRepoName string) string {
	formatString := "https://github.com/%v/%v.git"
	accountName := ownerOrOrganization(owner, organization)
	return fmt.Sprintf(formatString, accountName, templateRepoName)
}

func ownerOrOrganization(owner string, organization string) string {
	if organization != "" {
		return organization
	} else {
		return owner
	}
}

func generateTaskDescriptionFilePath(relativePath string) (string, error) {
	folderPath, err := GenerateChannelFolderPath()
	if err != nil {
		return "", err
	}

	return folderPath + "/issue-templates/" + relativePath, err
}
