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

	"github.com/google/go-github/github"
)

func createRepository(repoName string, organization string, token string) (string, error) {
	context := context.Background()

	private := true
	repositoryInput := github.Repository{
		Name:    &repoName,
		Private: &private,
	}

	tokenClient := getTokenClient(token)
	client := github.NewClient(tokenClient)
	repository, _, err := client.Repositories.Create(context, organization, &repositoryInput)
	if err != nil {
		return "", err
	}

	return *repository.CloneURL, nil
}

func createIssue(issue Issue, accountName string, repoName string, token string) error {
	context := context.Background()

	issueRequest := github.IssueRequest{
		Title:  &issue.Title,
		Body:   &issue.Description,
		Labels: &[]string{issue.Discipline},
	}

	tokenClient := getTokenClient(token)
	client := github.NewClient(tokenClient)

	_, _, err := client.Issues.Create(context, accountName, repoName, &issueRequest)

	return err
}

func addCollaborator(githubName string, accountName string, repoName string, token string) error {
	context := context.Background()

	tokenClient := getTokenClient(token)
	client := github.NewClient(tokenClient)

	options := github.RepositoryAddCollaboratorOptions{
		Permission: "push",
	}

	_, err := client.Repositories.AddCollaborator(context, accountName, repoName, githubName, &options)

	return err
}
