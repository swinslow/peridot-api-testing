// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getUsersTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		usersGetAdmin,
		usersGetOperator,
		usersPostAdmin,
		usersPostOperator,
	}
}

func usersGetAdmin(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users",
		ID:      "GET (admin)",
	}

	res.Wanted = `{"users":[{"id":1,"name":"Admin","github":"admin","access":"admin"},{"id":2,"name":"Operator User","github":"operator","access":"operator"},{"id":3,"name":"Commenter User","github":"commenter","access":"commenter"},{"id":4,"name":"Viewer User","github":"viewer","access":"viewer"},{"id":5,"name":"Disabled User","github":"disabled","access":"disabled"}]}`
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

func usersGetOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users",
		ID:      "GET (operator)",
	}

	res.Wanted = `{"users":[{"id":1,"github":"admin"},{"id":2,"github":"operator"},{"id":3,"github":"commenter"},{"id":4,"github":"viewer"},{"id":5,"github":"disabled"}]}`
	url := root + "/users"
	err := utils.GetContent(res, "1", url, 200, "operator")
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
	res.Wanted = `{"success": true, "id": 6}`
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
	res.Wanted = `{"users":[{"id":1,"name":"Admin","github":"admin","access":"admin"},{"id":2,"name":"Operator User","github":"operator","access":"operator"},{"id":3,"name":"Commenter User","github":"commenter","access":"commenter"},{"id":4,"name":"Viewer User","github":"viewer","access":"viewer"},{"id":5,"name":"Disabled User","github":"disabled","access":"disabled"}, {"id": 6, "name": "Steve Winslow", "github": "swinslow", "access": "operator"}]}`
	err = utils.GetContent(res, "3", url, 200, "admin")
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

func usersPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users",
		ID:      "POST (operator)",
	}

	// first, send POST to add a new user
	body := `{"name": "Steve Winslow", "github": "swinslow", "access": "operator"}`
	res.Wanted = `{"success": false, "error": "Access denied"}`
	url := root + "/users"
	err := utils.Post(res, "1", url, body, 403, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// and confirm that a new user was NOT actually added
	res.Wanted = `{"users":[{"id":1,"github":"admin"},{"id":2,"github":"operator"},{"id":3,"github":"commenter"},{"id":4,"github":"viewer"},{"id":5,"github":"disabled"}]}`
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
