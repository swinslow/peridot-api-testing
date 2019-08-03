// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getRepoPullsTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		repoPullsSubGetViewer,
		repoPullsSubWithCommitPostOperator,
		repoPullsGetOneViewer,
		repoPullsDeleteOneAdmin,
		repoPullsDeleteOneOperator,
	}
}

// ===== GET /repos/id/branches/branch

func repoPullsSubGetViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repos/{id}/branches/{branch}",
		ID:      "GET (viewer)",
	}

	url := root + "/repos/2/branches/dev-2.1"

	res.Wanted = `{"pulls":[
		{"id":2,"repo_id":2,"branch":"dev-2.1","started_at":"0001-01-01T00:00:00Z","finished_at":"0001-01-01T00:00:00Z","status":"startup","health":"ok","commit":"7864e74c9f54b1da4a64aaf7587ffa7880392233","spdx_id":""},
		{"id":4,"repo_id":2,"branch":"dev-2.1","started_at":"0001-01-01T00:00:00Z","finished_at":"0001-01-01T00:00:00Z","status":"startup","health":"ok","commit":"9f54b1da4a64aaf7587ffa78803922337864e74c","spdx_id":""}
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

// ===== POST /repos/id/branches/branch

func repoPullsSubWithCommitPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repos/{id}/branches/{branch}",
		ID:      "POST (operator)",
	}

	url := root + "/repos/2/branches/dev-2.1"

	// first, send POST to set up a repo pull with the requested commit
	// NOTE this is a made-up commit + branch + repo so cannot actually get pulled
	body := `{"commit": "803922337864e74c9f54b1da4a64aaf7587ffa78"}`
	res.Wanted = `{"id":6}`
	err := utils.Post(res, "1", url, body, 201, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new repo pull was actually added
	// NOTE output and tag are omitempty so will not be included here
	res.Wanted = `{"repopull":{"id":6,"repo_id":2,"branch":"dev-2.1","started_at":"0001-01-01T00:00:00Z","finished_at":"0001-01-01T00:00:00Z","status":"startup","health":"ok","commit":"803922337864e74c9f54b1da4a64aaf7587ffa78","spdx_id":""}}`
	repoPullURL := root + "/repopulls/6"
	err = utils.GetContent(res, "3", repoPullURL, 200, "operator")
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

// ===== GET /repopulls/id

func repoPullsGetOneViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repopulls/{id}",
		ID:      "GET (viewer)",
	}

	url := root + "/repopulls/5"

	res.Wanted = `{"repopull":{"id":5,"repo_id":1,"branch":"testing","started_at":"0001-01-01T00:00:00Z","finished_at":"0001-01-01T00:00:00Z","status":"startup","health":"ok","commit":"b1da4a64aaf7587ffa78803922337864e74c9f54","spdx_id":""}}`
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

// ===== DELETE /repopulls/id

func repoPullsDeleteOneAdmin(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repopulls/{id}",
		ID:      "DELETE (admin)",
	}

	url := root + "/repopulls/4"

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

	// now, confirm that the repopull is gone
	allURL := root + "/repos/2/branches/dev-2.1"

	res.Wanted = `{"pulls":[
		{"id":2,"repo_id":2,"branch":"dev-2.1","started_at":"0001-01-01T00:00:00Z","finished_at":"0001-01-01T00:00:00Z","status":"startup","health":"ok","commit":"7864e74c9f54b1da4a64aaf7587ffa7880392233","spdx_id":""}
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

func repoPullsDeleteOneOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repopulls/{id}",
		ID:      "DELETE (operator)",
	}

	url := root + "/repopulls/4"

	// try and fail to delete the repopull
	res.Wanted = `{"error": "Access denied"}`
	err := utils.Delete(res, "1", url, ``, 403, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the repopull has NOT been deleted
	allURL := root + "/repos/2/branches/dev-2.1"

	res.Wanted = `{"pulls":[
		{"id":2,"repo_id":2,"branch":"dev-2.1","started_at":"0001-01-01T00:00:00Z","finished_at":"0001-01-01T00:00:00Z","status":"startup","health":"ok","commit":"7864e74c9f54b1da4a64aaf7587ffa7880392233","spdx_id":""},
		{"id":4,"repo_id":2,"branch":"dev-2.1","started_at":"0001-01-01T00:00:00Z","finished_at":"0001-01-01T00:00:00Z","status":"startup","health":"ok","commit":"9f54b1da4a64aaf7587ffa78803922337864e74c","spdx_id":""}
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
