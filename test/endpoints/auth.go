// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func runAuthTests(root string) []*testresult.TestResult {
	rs := []*testresult.TestResult{}
	var res *testresult.TestResult

	res = loginGet(root)
	rs = append(rs, res)

	return rs
}

func loginGet(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "login",
		ID:      "GET",
	}

	url := root + "/auth/login"
	err := utils.GetContentNoFollow(res, "1", url, 307, "none")
	if err != nil {
		return res
	}

	utils.Pass(res)
	return res
}
