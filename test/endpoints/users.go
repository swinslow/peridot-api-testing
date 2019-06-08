// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/obsidian-api-testing/internal/testresult"
	"github.com/swinslow/obsidian-api-testing/test/utils"
)

func runUsersTests(root string) []*testresult.TestResult {
	rs := []*testresult.TestResult{}
	var res *testresult.TestResult

	res = usersGet(root)
	rs = append(rs, res)

	return rs
}

func usersGet(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users",
		ID:      "GET",
	}

	res.Wanted = `{"users": [{"id": 1, "name": "Admin", "email": "test@example.com", "access": admin"}]}`
	url := root + "/users"
	err := utils.GetContent(res, "1", url, 200)
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
