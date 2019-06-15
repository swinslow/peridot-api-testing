// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getUsersTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		usersGetAdmin,
		usersPostAdmin,
	}
}

func usersGetAdmin(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users",
		ID:      "GET (admin)",
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

func usersPostAdmin(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users",
		ID:      "POST (admin)",
	}

	// first, send POST to add a new user
	body := `{"name": "Steve Winslow", "github": "swinslow", "access": "operator"}`
	res.Wanted = `{"success": true, "id": 2}`
	url := root + "/users"
	err := utils.Post(res, "1", url, body, 200, "admin")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new user was actually added
	res.Wanted = `{"users": [{"id": 1, "name": "Admin", "github": "admin", "access": "admin"}, {"id": 2, "name": "Steve Winslow", "github": "swinslow", "access": "operator"}]}`
	err = utils.GetContent(res, "1", url, 200, "admin")
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
