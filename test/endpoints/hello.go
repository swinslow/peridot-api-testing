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

	url := root + "/hello"
	got, err := utils.GetContent(res, "1", url)
	if err != nil {
		return res
	}

	wanted := `{"success": true, "message": "hello"}`
	if !utils.IsMatch(wanted, got) {
		utils.FailMatch(res, "2", wanted, got)
		return res
	}

	utils.Pass(res)
	return res
}
