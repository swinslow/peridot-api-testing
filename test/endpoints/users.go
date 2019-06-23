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
		usersGetOneAdmin,
		usersGetOneOperatorSelf,
		usersGetOneOperatorOther,
		usersPutOneAdmin,
		usersPutOneOperatorSelf,
		usersPutOneOperatorOther,
	}
}

// ===== GET /users

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

// ===== POST /users

func usersPostAdmin(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users",
		ID:      "POST (admin)",
	}

	// first, send POST to add a new user
	body := `{"name": "Steve Winslow", "github": "swinslow", "access": "operator"}`
	res.Wanted = `{"id": 6}`
	url := root + "/users"
	err := utils.Post(res, "1", url, body, 201, "admin")
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
	res.Wanted = `{"error": "Access denied"}`
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

// ===== GET /users/id

func usersGetOneAdmin(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users/{id}",
		ID:      "GET (admin)",
	}

	res.Wanted = `{"user":{"id":2,"name":"Operator User","github":"operator","access":"operator"}}`
	url := root + "/users/2"
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

func usersGetOneOperatorSelf(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users/{id}",
		ID:      "GET (operator-self)",
	}

	res.Wanted = `{"user":{"id":2,"name":"Operator User","github":"operator","access":"operator"}}`
	url := root + "/users/2"
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

func usersGetOneOperatorOther(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users/{id}",
		ID:      "GET (operator-other)",
	}

	res.Wanted = `{"user":{"id":4,"github":"viewer"}}`
	url := root + "/users/4"
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

// ===== PUT /users/id

func usersPutOneAdmin(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users/{id}",
		ID:      "PUT (admin)",
	}

	// first, send PUT to modify an existing user
	body := `{"name": "Steve Winslow", "github": "swinslow", "access": "operator"}`
	res.Wanted = `{"success": true}`
	url := root + "/users/5"
	err := utils.Put(res, "1", url, body, 200, "admin")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the user data was actually updated
	res.Wanted = `{"user":{"id":5,"name":"Steve Winslow","github":"swinslow","access":"operator"}}`
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

func usersPutOneOperatorSelf(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users/{id}",
		ID:      "PUT (operator-self)",
	}

	// first, send PUT to modify own name (NOT github / access)
	body := `{"name": "Steve Winslow"}`
	res.Wanted = `{"success": true}`
	url := root + "/users/2"
	err := utils.Put(res, "1", url, body, 200, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the user data was actually updated
	res.Wanted = `{"user":{"id":2,"name":"Steve Winslow","github":"operator","access":"operator"}}`
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

func usersPutOneOperatorOther(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "users/{id}",
		ID:      "PUT (operator-other)",
	}

	// try and fail to send PUT to modify other's name
	body := `{"name": "OOPS"}`
	res.Wanted = `{"error": "Access denied"}`
	url := root + "/users/3"
	err := utils.Put(res, "1", url, body, 403, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// also try and fail to send PUT to modify other's github
	body = `{"github": "oops"}`
	res.Wanted = `{"error": "Access denied"}`
	err = utils.Put(res, "3", url, body, 403, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "4")
		return res
	}

	// finally, confirm that the other user's data was NOT actually updated
	res.Wanted = `{"user":{"id":3,"github":"commenter"}}`
	err = utils.GetContent(res, "5", url, 200, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "6")
		return res
	}

	utils.Pass(res)
	return res
}
