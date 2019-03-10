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

package challenge

import (
	"fmt"
	"io/ioutil"
	"os"

	config "github.com/keremk/challenge/config"
)

// Creates a coding challenge for a given candidate and challenge type.
// The coding challenge is created based on the configuration settings the .challenge.yaml file
func CreateChallenge(candidateName string, discipline string) {
	fmt.Println("Creating coding challenge")
	challengeConfig := config.GetConfigInstance()

	repoName := generateChallengeRepositoryName(candidateName, discipline)
	fmt.Println("Challenge repo name: ", repoName)
	fmt.Println("Organization: ", challengeConfig.Organization)
	fmt.Println("Owner: ", challengeConfig.Owner)
	fmt.Println("TrackingRepo: ", challengeConfig.TrackingRepoName)

	challengeRepoURL, err := createRepository(repoName, challengeConfig.Organization, challengeConfig.Creator.GithubToken)
	if err != nil {
		fmt.Println("Cannot create the challenge repository")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Created: ", challengeRepoURL)

	challenge, err := challengeConfig.FindChallenge(discipline)
	if err != nil {
		fmt.Println("Invalid challenge discipline ", discipline)
		os.Exit(1)
	}

	templateRepoURL := generateTemplateRepositoryName(challengeConfig.GithubAccount, challenge.TemplateRepo)
	fmt.Println(templateRepoURL)

	err = pushStarterProject(templateRepoURL, challengeRepoURL, challengeConfig.Creator.GithubToken)

	if err != nil {
		fmt.Println("Could not push the starter project")
		fmt.Println(err)
		os.Exit(1)
	}

	CreateCandidateTask(candidateName, discipline, 0)
	createTrackingIssue(candidateName, discipline, challengeRepoURL)

	err = addCollaborator(candidateName, challengeConfig.Owner, repoName, challengeConfig.Creator.GithubToken)

	if err != nil {
		fmt.Println("Cannot add the candidate as a collaborator ", candidateName)
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Challenge created successfully.")
}

func CreateCandidateTask(candidateName string, discipline string, level int) {
	fmt.Println("Creating candidate task")
	challengeConfig := config.GetConfigInstance()

	challenge, err := challengeConfig.FindChallenge(discipline)
	if err != nil {
		fmt.Println("Invalid challenge discipline ", discipline)
		os.Exit(1)
	}

	if level >= len(challenge.Tasks) {
		fmt.Println("No task specified for the level ", level)
		os.Exit(1)
	}

	task := challenge.Tasks[level]
	descriptionFilePath := challengeConfig.ChallengeFolder + "issue-templates/" + challenge.Tasks[level].DescriptionFile
	description, err := readDescription(descriptionFilePath)
	if err != nil {
		fmt.Println("Aborting task creation")
		fmt.Println(err)
		os.Exit(1)
	}

	issue := Issue{
		Title:       task.Title,
		Discipline:  discipline,
		Description: description,
	}

	repoName := generateChallengeRepositoryName(candidateName, discipline)
	err = createIssue(issue, challengeConfig.Owner, repoName, challengeConfig.Creator.GithubToken)

	if err != nil {
		fmt.Println("Could not create a candidate task at ", repoName)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Candidate task created at: ", repoName)
}

func createTrackingIssue(candidateName string, discipline string, challengeRepoURL string) {
	challengeConfig := config.GetConfigInstance()

	title := "Coding Challenge for: " + candidateName
	description := "Coding challenge is located at: " + challengeRepoURL

	issue := Issue{
		Title:       title,
		Discipline:  discipline,
		Description: description,
	}
	err := createIssue(issue, challengeConfig.Owner, challengeConfig.TrackingRepoName, challengeConfig.Creator.GithubToken)

	if err != nil {
		fmt.Println("Could not create a tracking issue at ", challengeConfig.TrackingRepoName)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Tracking issue created at: ", challengeConfig.TrackingRepoName)
}

func readDescription(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("There is no file or cannot read in location: ", filePath)
		return "", err
	}

	return string(bytes[:]), nil
}
