// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getJobsTests() []testresult.TestFunc {
	return []testresult.TestFunc{}
}

// ===== POST /repopulls/id/jobs

func jobsPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repopulls/{id}/jobs",
		ID:      "POST (operator)",
	}

	url := root + "/repopulls/3/jobs"

	// first, send POST to add a new job
	body := `{"agent_id":1, "is_ready":false, "priorjob_ids":[],
	"config":{"kv": {"hi": "there", "hello": "world"}}}`
	res.Wanted = `{"id": 1}`
	err := utils.Post(res, "1", url, body, 201, "operator")
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
