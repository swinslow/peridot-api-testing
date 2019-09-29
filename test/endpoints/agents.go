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
		agentsGetOneViewer,
		agentsPutOneOperator,
		agentsPutOneIsActiveOnlyOperator,
		agentsPutOneStatusOnlyOperator,
		agentsPutOneAbilitiesOnlyOperator,
		agentsPutOneViewer,
		agentsDeleteOneAdmin,
		agentsDeleteOneOperator,
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

// ===== GET /agents/id

func agentsGetOneViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents/{id}",
		ID:      "GET (viewer)",
	}

	url := root + "/agents/2"

	res.Wanted = `{"agent":{"id":2, "name":"read-magic", "is_active":true, "address":"https://example.com/read-magic", "port":2088, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":true}}`
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

// ===== PUT /agents/id

func agentsPutOneOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents/{id}",
		ID:      "PUT (operator)",
	}

	url := root + "/agents/2"

	// first, send PUT to update an existing agent
	body := `{"is_active":false, "address":"https://example.com/new-address", "port":3077, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false}`
	res.Wanted = ``
	err := utils.Put(res, "1", url, body, 204, "operator")
	if err != nil {
		return res
	}

	if !utils.IsEmpty(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the agent was actually updated
	res.Wanted = `{"agent":{"id":2, "name":"read-magic", "is_active":false, "address":"https://example.com/new-address", "port":3077, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false}}`
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

func agentsPutOneIsActiveOnlyOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents/{id}",
		ID:      "PUT (operator, is_active)",
	}

	url := root + "/agents/2"

	// first, send PUT to update an existing agent
	body := `{"is_active":false}`
	res.Wanted = ``
	err := utils.Put(res, "1", url, body, 204, "operator")
	if err != nil {
		return res
	}

	if !utils.IsEmpty(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the agent was actually updated
	res.Wanted = `{"agent":{"id":2, "name":"read-magic", "is_active":false, "address":"https://example.com/read-magic", "port":2088, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":true}}`
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

func agentsPutOneStatusOnlyOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents/{id}",
		ID:      "PUT (operator, status)",
	}

	url := root + "/agents/2"

	// first, send PUT to update an existing agent
	body := `{"is_active":false, "address":"https://example.com/new-address", "port":3077}`
	res.Wanted = ``
	err := utils.Put(res, "1", url, body, 204, "operator")
	if err != nil {
		return res
	}

	if !utils.IsEmpty(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the agent was actually updated
	res.Wanted = `{"agent":{"id":2, "name":"read-magic", "is_active":false, "address":"https://example.com/new-address", "port":3077, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":true}}`
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

func agentsPutOneAbilitiesOnlyOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents/{id}",
		ID:      "PUT (operator, abilities)",
	}

	url := root + "/agents/2"

	// first, send PUT to update an existing agent
	body := `{"is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false}`
	res.Wanted = ``
	err := utils.Put(res, "1", url, body, 204, "operator")
	if err != nil {
		return res
	}

	if !utils.IsEmpty(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the agent was actually updated
	res.Wanted = `{"agent":{"id":2, "name":"read-magic", "is_active":true, "address":"https://example.com/read-magic", "port":2088, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false}}`
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

func agentsPutOneViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents/{id}",
		ID:      "PUT (viewer)",
	}

	url := root + "/agents/2"

	body := `{"is_active":false, "address":"https://example.com/new-address", "port":3077, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false}`
	res.Wanted = `{"error": "Access denied"}`
	err := utils.Put(res, "1", url, body, 403, "viewer")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the agent was NOT actually updated
	res.Wanted = `{"agent":{"id":2, "name":"read-magic", "is_active":true, "address":"https://example.com/read-magic", "port":2088, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":true}}`
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

// ===== DELETE /agents/id

func agentsDeleteOneAdmin(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents/{id}",
		ID:      "DELETE (admin)",
	}

	url := root + "/agents/2"

	// send a delete request
	res.Wanted = ``
	err := utils.Delete(res, "1", url, ``, 204, "admin")
	if err != nil {
		return res
	}

	if !utils.IsEmpty(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the agent is gone
	allURL := root + "/agents"
	res.Wanted = `{"agents":[
		{"id":1, "name":"do-magic", "is_active":true, "address":"https://example.com/do-magic", "port":2087, "is_codereader":false, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false},
		{"id":3, "name":"disabled", "is_active":false, "address":"localhost", "port":2057, "is_codereader":false, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false},
		{"id":4, "name":"wevs", "is_active":true, "address":"localhost", "port":5010, "is_codereader":true, "is_spdxreader":true, "is_codewriter":true, "is_spdxwriter":false}
	]}`
	err = utils.GetContent(res, "3", allURL, 200, "viewer")
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

func agentsDeleteOneOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "agents/{id}",
		ID:      "DELETE (operator)",
	}

	url := root + "/agents/2"

	// try and fail to delete the agent
	res.Wanted = `{"error": "Access denied"}`
	err := utils.Delete(res, "1", url, ``, 403, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the agent has NOT been deleted
	allURL := root + "/agents"
	res.Wanted = `{"agents":[
		{"id":1, "name":"do-magic", "is_active":true, "address":"https://example.com/do-magic", "port":2087, "is_codereader":false, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false},
		{"id":2, "name":"read-magic", "is_active":true, "address":"https://example.com/read-magic", "port":2088, "is_codereader":true, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":true},
		{"id":3, "name":"disabled", "is_active":false, "address":"localhost", "port":2057, "is_codereader":false, "is_spdxreader":true, "is_codewriter":false, "is_spdxwriter":false},
		{"id":4, "name":"wevs", "is_active":true, "address":"localhost", "port":5010, "is_codereader":true, "is_spdxreader":true, "is_codewriter":true, "is_spdxwriter":false}
	]}`
	err = utils.GetContent(res, "3", allURL, 200, "viewer")
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
