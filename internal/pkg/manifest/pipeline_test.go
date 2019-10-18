// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package manifest

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePipeline(t *testing.T) {
	testCases := map[string]struct {
		beforeEach     func() error
		provider       Provider
		expectedErr    error
		inputStages    []PipelineStage
		expectedStages []PipelineStage
	}{
		"missing provider configuration": {
			beforeEach: func() error {
				// reset the global map before each test case is run
				providerSpecificProperties = make(map[Provider]interface{})
				return nil
			},
			provider: Github,
			expectedErr: &ErrMissingProviderProperties{
				provider: Github,
			},
		},
		"happy case with default stages": {
			beforeEach: func() error {
				// reset the global map before each test case is run
				providerSpecificProperties = make(map[Provider]interface{})

				return Github.ConfiguredWith(&GithubProperties{
					Repository: "aws/amazon-ecs-cli-v2",
				})
			},
			provider: Github,
			expectedStages: []PipelineStage{
				{
					Name: "Test",
				},
				{
					Name: "Prod",
				},
			},
		},
		"happy case with non-default stages": {
			beforeEach: func() error {
				// reset the global map before each test case is run
				providerSpecificProperties = make(map[Provider]interface{})

				return Github.ConfiguredWith(&GithubProperties{
					Repository: "aws/amazon-ecs-cli-v2",
				})
			},
			provider: Github,
			inputStages: []PipelineStage{
				{
					Name: "Chicken",
				},
				{
					Name: "Wings",
				},
			},
			expectedStages: []PipelineStage{
				{
					Name: "Chicken",
				},
				{
					Name: "Wings",
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if tc.beforeEach != nil {
				err := tc.beforeEach()
				require.NoError(t, err, "unexpected error while running beforeEach")
			}
			m, err := CreatePipeline(tc.provider, tc.inputStages...)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				p, ok := m.(*PipelineManifest)
				require.True(t, ok)
				require.Equal(t, tc.expectedStages, p.Environments, "the environments are different from the expected")
			}
		})
	}
}

func TestPipelineManifestMarshal(t *testing.T) {
	wantedContent := `# This YAML file defines the relationship and deployment ordering of your environments.

# The version of the schema used in this template
version: 1

# This section defines the source artifacts.
source:
  # The name of the provider that is used to store the source artifacts.
  provider: github
  # Additional properties that further specifies the exact location
  # the artifacts should be sourced from.
  properties:
    # Repository in the format of <user_name>/<repository_name>.
    repository: aws/amazon-ecs-cli-v2


# The deployment section defines the order the pipeline will deploy
# to your environments
stages:
    - 
      # The name of the environment to deploy to.
      env: Test
    
    - 
      # The name of the environment to deploy to.
      env: Prod
    
`
	// reset the global map before each test case is run
	providerSpecificProperties = make(map[Provider]interface{})
	Github.ConfiguredWith(&GithubProperties{
		Repository: "aws/amazon-ecs-cli-v2",
	})

	m, err := CreatePipeline(Github)
	require.NoError(t, err)

	b, err := m.Marshal()
	require.NoError(t, err)
	require.Equal(t, wantedContent, strings.Replace(string(b), "\r\n", "\n", -1))
}

func TestUnmarshalPipeline(t *testing.T) {
	testCases := map[string]struct {
		inContent        string
		expectedManifest *PipelineManifest
		expectedErr      error
	}{
		"invalid pipeline schema version": {
			inContent: `
version: -1

source:
  provider: github
  properties:
    repository: aws/somethingCool

stages:
    - 
      env: Test
    - 
      env: Prod
`,
			expectedErr: &ErrInvalidPipelineManifestVersion{
				PipelineSchemaMajorVersion(-1),
			},
		},
		"invalid pipeline.yml": {
			inContent:   `corrupted yaml`,
			expectedErr: errors.New("yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `corrupt...` into manifest.PipelineManifest"),
		},
		"valid pipeline.yml": {
			inContent: `
version: 1

source:
  provider: github
  properties:
    repository: aws/somethingCool

stages:
    - 
      env: Chicken
    - 
      env: Wings
`,
			expectedManifest: &PipelineManifest{
				Version: Ver1,
				Source: &Source{
					ProviderName: "github",
					Properties:   map[string]interface{}{"repository": "aws/somethingCool"},
				},
				Environments: []PipelineStage{
					{
						Name: "Chicken",
					},
					{
						Name: "Wings",
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m, err := UnmarshalPipeline([]byte(tc.inContent))

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				actualManifest, ok := m.(*PipelineManifest)
				require.True(t, ok)
				require.Equal(t, actualManifest, tc.expectedManifest)
			}
		})
	}
}
