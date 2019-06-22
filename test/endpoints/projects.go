// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getProjectsTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		projectsGetViewer,
		projectsPostOperator,
		projectsPostViewer,
		projectsGetOneViewer,
		projectsPutOneOperator,
		projectsPutOneViewer,
	}
}

// ===== GET /projects

func projectsGetViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "projects",
		ID:      "GET (viewer)",
	}

	url := root + "/projects"

	res.Wanted = `{"projects":[{"id":1,"name":"xyzzy","fullname":"The xyzzy Project"},{"id":2,"name":"frotz","fullname":"The frotz Project"},{"id":3,"name":"gnusto","fullname":"The gnusto Project"}]}`
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

// ===== POST /projects

func projectsPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "projects",
		ID:      "POST (operator)",
	}

	url := root + "/projects"

	// first, send POST to add a new project
	body := `{"name": "plugh", "fullname": "The plugh Project"}`
	res.Wanted = `{"success": true, "id": 4}`
	err := utils.Post(res, "1", url, body, 200, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new project was actually added
	res.Wanted = `{"projects":[{"id":1,"name":"xyzzy","fullname":"The xyzzy Project"},{"id":2,"name":"frotz","fullname":"The frotz Project"},{"id":3,"name":"gnusto","fullname":"The gnusto Project"},{"id":4,"name":"plugh","fullname":"The plugh Project"}]}`
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

func projectsPostViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "projects",
		ID:      "POST (viewer)",
	}

	url := root + "/projects"

	// first, try and fail to add a new project
	body := `{"name": "plugh", "fullname": "The plugh Project"}`
	res.Wanted = `{"success": false, "error": "Access denied"}`
	err := utils.Post(res, "1", url, body, 403, "viewer")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new project was NOT actually added
	res.Wanted = `{"projects":[{"id":1,"name":"xyzzy","fullname":"The xyzzy Project"},{"id":2,"name":"frotz","fullname":"The frotz Project"},{"id":3,"name":"gnusto","fullname":"The gnusto Project"}]}`
	err = utils.GetContent(res, "3", url, 200, "viewer")
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

// ===== GET /projects/id

func projectsGetOneViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "projects/{id}",
		ID:      "GET (viewer)",
	}

	res.Wanted = `{"success": true, "project":{"id":2,"name":"frotz","fullname":"The frotz Project"}}`
	url := root + "/projects/2"
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

// ===== PUT /projects/id

func projectsPutOneOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "projects/{id}",
		ID:      "PUT (operator)",
	}

	url := root + "/projects/2"

	// first, send PUT to update an existing project
	body := `{"name": "plugh", "fullname": "The plugh Project"}`
	res.Wanted = `{"success": true}`
	err := utils.Put(res, "1", url, body, 200, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the project was actually updated
	res.Wanted = `{"success": true, "project":{"id":2,"name":"plugh","fullname":"The plugh Project"}}`
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

func projectsPutOneViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "projects/{id}",
		ID:      "PUT (viewer)",
	}

	url := root + "/projects/2"

	body := `{"name": "plugh", "fullname": "The plugh Project"}`
	res.Wanted = `{"success": false, "error": "Access denied"}`
	err := utils.Put(res, "1", url, body, 403, "viewer")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that the project was NOT actually updated
	res.Wanted = `{"success": true, "project":{"id":2,"name":"frotz","fullname":"The frotz Project"}}`
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