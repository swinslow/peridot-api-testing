// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/obsidian-api-testing/internal/testresult"
	"github.com/swinslow/obsidian-api-testing/test/utils"
)

func runHelloTests(root string) []*testresult.TestResult {
	rs := []*testresult.TestResult{}
	var res *testresult.TestResult

	res = helloGet(root)
	rs = append(rs, res)

	return rs
}

func helloGet(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "hello",
		ID:      "GET",
	}

	res.Wanted = `{"success": true, "message": "hello"}`
	url := root + "/hello"
	err := utils.GetContent(res, "1", url, 200, "none")
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
