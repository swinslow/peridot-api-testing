// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getAgentsTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		agentsGetViewer,
		agentsPostOperator,
	}
}

// ===== GET /agents

func agentsGetViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents",
		ID:      "GET (viewer)",
	}

	url := root + "/agents"

	res.Wanted = `{"agents":[
		{"id":1, "name":"do-magic", "is_active":true, "address":"https://example.com/do-magic", "port":2087, "is_codereader":false, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false},
		{"id":2, "name":"read-magic", "is_active":true, "address":"https://example.com/read-magic", "port":2088, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":true},
		{"id":3, "name":"disabled", "is_active":false, "address":"localhost", "port":2057, "is_codereader":false, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false},
		{"id":4, "name":"wevs", "is_active":true, "address":"localhost", "port":5010, "is_codereader":true, "is_spdxreader":true, "is_codewriter":true, "is_spdxwriter":false}
	]}`
	err := utils.GetContent(res, "1", url, 200, "viewer")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	utils.Pass(res)
	return res
}

// ===== POST /agents

func agentsPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents",
		ID:      "POST (operator)",
	}

	url := root + "/agents"

	// first, send POST to add a new agent
	body := `{"name":"idsearcher", "is_active":true, "address":"localhost", "port":9014, "is_codereader":true, "is_spdxreader":false, "is_codewriter":false, "is_spdxwriter":true}`
	res.Wanted = `{"id": 5}`
	err := utils.Post(res, "1", url, body, 201, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new agent was actually added
	res.Wanted = `{"agents":[
		{"id":1, "name":"do-magic", "is_active":true, "address":"https://example.com/do-magic", "port":2087, "is_codereader":false, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false},
		{"id":2, "name":"read-magic", "is_active":true, "address":"https://example.com/read-magic", "port":2088, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":true},
		{"id":3, "name":"disabled", "is_active":false, "address":"localhost", "port":2057, "is_codereader":false, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false},
		{"id":4, "name":"wevs", "is_active":true, "address":"localhost", "port":5010, "is_codereader":true, "is_spdxreader":true, "is_codewriter":true, "is_spdxwriter":false},
		{"id":5, "name":"idsearcher", "is_active":true, "address":"localhost", "port":9014, "is_codereader":true, "is_spdxreader":false, "is_codewriter":false, "is_spdxwriter":true}
	]}`
	err = utils.GetContent(res, "3", url, 200, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "4")
		return res
	}

	utils.Pass(res)
	return res
}
