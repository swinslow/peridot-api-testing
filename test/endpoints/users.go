// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getUsersTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		usersGet,
	}
}

func usersGet(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users",
		ID:      "GET",
	}

	res.Wanted = `{"users": [{"id": 1, "name": "Admin", "github": "admin", "access": "admin"}]}`
	url := root + "/users"
	err := utils.GetContent(res, "1", url, 200, "admin")
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
