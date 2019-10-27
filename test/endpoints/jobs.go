// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getJobsTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		jobsSubGetOperator,
		jobsSubPostOperator,
	}
}

// ===== GET /repopulls/id/jobs

func jobsSubGetOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repopulls/{id}/jobs",
		ID:      "GET (viewer)",
	}

	url := root + "/repopulls/4/jobs"

	res.Wanted = `{"jobs":[
		{"id":2, "repopull_id":4, "agent_id":1, "started_at":"0001-01-01T00:00:00Z", "finished_at":"0001-01-01T00:00:00Z", "status":"startup", "health":"ok", "is_ready":true, "config":{}},
		{"id":3, "repopull_id":4, "agent_id":2, "priorjob_ids": [2], "started_at":"0001-01-01T00:00:00Z", "finished_at":"0001-01-01T00:00:00Z", "status":"startup", "health":"ok", "is_ready":true, "config":{"codereader": {"primary": {"path": "/somewhere"}}}},
		{"id":4, "repopull_id":4, "agent_id":4, "priorjob_ids": [2,3], "started_at":"0001-01-01T00:00:00Z", "finished_at":"0001-01-01T00:00:00Z", "status":"startup", "health":"ok", "is_ready":false, "config":{"kv": {"hello":"world"}, "codereader": {"godeps": {"priorjob_id": 3}}, "spdxreader": {"primary": {"path": "/path/wherever"}, "godeps": {"priorjob_id": 3}}}}
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

// ===== POST /repopulls/id/jobs

func jobsSubPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repopulls/{id}/jobs",
		ID:      "POST (operator)",
	}

	url := root + "/repopulls/3/jobs"

	// first, send POST to add a new job
	body := `{"agent_id":1, "is_ready":false, "priorjob_ids":[],
		"config":{"kv": {"hi": "there", "hello": "world"}}
	}`
	res.Wanted = `{"id": 5}`
	err := utils.Post(res, "1", url, body, 201, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new job was actually added
	// this should be the only one for repopull 3 so we can reuse the same url
	// priorjob_ids and some config vals should be absent
	res.Wanted = `{"jobs":[
		{"id":5, "repopull_id":3, "agent_id":1, "started_at":"0001-01-01T00:00:00Z", "finished_at":"0001-01-01T00:00:00Z", "status":"startup", "health":"ok", "is_ready":false, "config":{"kv": {"hi": "there", "hello": "world"}}}
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
