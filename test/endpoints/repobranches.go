// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getRepoBranchesTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		repoBranchesSubGetViewer,
		repoBranchesSubPostOperator,
	}
}

// ===== GET /repos/id/branches

func repoBranchesSubGetViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repos/{id}/branches",
		ID:      "GET (viewer)",
	}

	url := root + "/repos/2/branches"

	// should be returned in alphabetical order
	res.Wanted = `{"branches":["dev","dev-2.1","master"]}`
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

// ===== POST /repos/id/branches

func repoBranchesSubPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repos/{id}/branches",
		ID:      "POST (operator)",
	}

	url := root + "/repos/2/branches"

	// first, send POST to add a new branch to the existing repo
	body := `{"branch": "issue-47"}`
	res.Wanted = `{"branch": "issue-47"}`
	err := utils.Post(res, "1", url, body, 201, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new repo was actually added
	// should be returned in alphabetical order
	res.Wanted = `{"branches":["dev","dev-2.1","issue-47","master"]}`
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
