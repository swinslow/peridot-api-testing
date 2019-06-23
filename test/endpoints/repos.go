// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getReposTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		reposGetViewer,
		reposPostOperator,
		reposSubGetViewer,
		reposSubPostOperator,
		// reposGetOneViewer,
		// reposPutOneOperator,
		// reposPutOneViewer,
		// reposDeleteOneAdmin,
		// reposDeleteOneOperator,
	}
}

// ===== GET /repos

func reposGetViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repos",
		ID:      "GET (viewer)",
	}

	url := root + "/repos"

	res.Wanted = `{"repos":[{"id":1,"subproject_id":2,"name":"filfre-core","address":"https://example.com/filfre-core.git"},{"id":2,"subproject_id":2,"name":"filfre-api","address":"https://example.com/filfre-api.git"},{"id":3,"subproject_id":1,"name":"blorple-c","address":"https://example.com/blorple-c.git"},{"id":4,"subproject_id":4,"name":"girgol","address":"https://example.com/girgol.git"}]}`
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

// ===== POST /repos

func reposPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "repos",
		ID:      "POST (operator)",
	}

	url := root + "/repos"

	// first, send POST to add a new repo
	body := `{"subproject_id": 2, "name": "filfre-webapp", "address": "https://example.com/filfre-webapp.git"}`
	res.Wanted = `{"id": 5}`
	err := utils.Post(res, "1", url, body, 201, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new repo was actually added
	res.Wanted = `{"repos":[{"id":1,"subproject_id":2,"name":"filfre-core","address":"https://example.com/filfre-core.git"},{"id":2,"subproject_id":2,"name":"filfre-api","address":"https://example.com/filfre-api.git"},{"id":3,"subproject_id":1,"name":"blorple-c","address":"https://example.com/blorple-c.git"},{"id":4,"subproject_id":4,"name":"girgol","address":"https://example.com/girgol.git"},{"id":5,"subproject_id":2,"name":"filfre-webapp","address":"https://example.com/filfre-webapp.git"}]}`
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

// ===== GET /subprojects/id/repos

func reposSubGetViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "subprojects/{id}/repos",
		ID:      "GET (viewer)",
	}

	url := root + "/subprojects/2/repos"

	res.Wanted = `{"repos":[{"id":1,"subproject_id":2,"name":"filfre-core","address":"https://example.com/filfre-core.git"},{"id":2,"subproject_id":2,"name":"filfre-api","address":"https://example.com/filfre-api.git"}]}`
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

// ===== POST /subprojects/id/repos

func reposSubPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "subprojects/{id}/repos",
		ID:      "POST (operator)",
	}

	url := root + "/subprojects/2/repos"

	// first, send POST to add a new repo
	body := `{"name": "filfre-webapp", "address": "https://example.com/filfre-webapp.git"}`
	res.Wanted = `{"id": 5}`
	err := utils.Post(res, "1", url, body, 201, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new repo was actually added
	url = root + "/repos"
	res.Wanted = `{"repos":[{"id":1,"subproject_id":2,"name":"filfre-core","address":"https://example.com/filfre-core.git"},{"id":2,"subproject_id":2,"name":"filfre-api","address":"https://example.com/filfre-api.git"},{"id":3,"subproject_id":1,"name":"blorple-c","address":"https://example.com/blorple-c.git"},{"id":4,"subproject_id":4,"name":"girgol","address":"https://example.com/girgol.git"},{"id":5,"subproject_id":2,"name":"filfre-webapp","address":"https://example.com/filfre-webapp.git"}]}`
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

// // ===== GET /subprojects/id

// func subprojectsGetOneViewer(root string) *testresult.TestResult {
// 	res := &testresult.TestResult{
// 		Suite:   "endpoints",
// 		Element: "subprojects/{id}",
// 		ID:      "GET (viewer)",
// 	}

// 	res.Wanted = `{"subproject":{"id":2,"project_id":2,"name":"filfre","fullname":"The filfre Subproject"}}`
// 	url := root + "/subprojects/2"
// 	err := utils.GetContent(res, "1", url, 200, "viewer")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsMatch(res) {
// 		utils.FailMatch(res, "2")
// 		return res
// 	}

// 	utils.Pass(res)
// 	return res
// }

// // ===== PUT /subprojects/id

// func subprojectsPutOneOperator(root string) *testresult.TestResult {
// 	res := &testresult.TestResult{
// 		Suite:   "endpoints",
// 		Element: "subprojects/{id}",
// 		ID:      "PUT (operator)",
// 	}

// 	url := root + "/subprojects/2"

// 	// first, send PUT to update an existing subproject
// 	body := `{"name": "plugh", "fullname": "The plugh Subproject"}`
// 	res.Wanted = ``
// 	err := utils.Put(res, "1", url, body, 204, "operator")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsEmpty(res) {
// 		utils.FailMatch(res, "2")
// 		return res
// 	}

// 	// now, confirm that the subproject was actually updated
// 	res.Wanted = `{"subproject":{"id":2,"project_id":2,"name":"plugh","fullname":"The plugh Subproject"}}`
// 	err = utils.GetContent(res, "3", url, 200, "operator")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsMatch(res) {
// 		utils.FailMatch(res, "4")
// 		return res
// 	}

// 	utils.Pass(res)
// 	return res
// }

// func subprojectsPutOneViewer(root string) *testresult.TestResult {
// 	res := &testresult.TestResult{
// 		Suite:   "endpoints",
// 		Element: "subprojects/{id}",
// 		ID:      "PUT (viewer)",
// 	}

// 	url := root + "/subprojects/2"

// 	body := `{"name": "plugh", "fullname": "The plugh Subproject"}`
// 	res.Wanted = `{"error": "Access denied"}`
// 	err := utils.Put(res, "1", url, body, 403, "viewer")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsMatch(res) {
// 		utils.FailMatch(res, "2")
// 		return res
// 	}

// 	// now, confirm that the subproject was NOT actually updated
// 	res.Wanted = `{"subproject":{"id":2,"project_id":2,"name":"filfre","fullname":"The filfre Subproject"}}`
// 	err = utils.GetContent(res, "3", url, 200, "operator")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsMatch(res) {
// 		utils.FailMatch(res, "4")
// 		return res
// 	}

// 	utils.Pass(res)
// 	return res
// }

// // ===== DELETE /subprojects/id

// func subprojectsDeleteOneAdmin(root string) *testresult.TestResult {
// 	res := &testresult.TestResult{
// 		Suite:   "endpoints",
// 		Element: "subprojects/{id}",
// 		ID:      "DELETE (admin)",
// 	}

// 	url := root + "/subprojects/2"

// 	// send a delete request
// 	res.Wanted = ``
// 	err := utils.Delete(res, "1", url, ``, 204, "admin")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsEmpty(res) {
// 		utils.FailMatch(res, "2")
// 		return res
// 	}

// 	// now, confirm that the subproject is gone
// 	allURL := root + "/subprojects"
// 	res.Wanted = `{"subprojects":[{"id":1,"project_id":2,"name":"blorple","fullname":"The blorple Subproject"},{"id":3,"project_id":2,"name":"fweep","fullname":"The fweep Subproject"},{"id":4,"project_id":3,"name":"girgol","fullname":"The girgol Subproject"}]}`
// 	err = utils.GetContent(res, "3", allURL, 200, "viewer")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsMatch(res) {
// 		utils.FailMatch(res, "4")
// 		return res
// 	}

// 	utils.Pass(res)
// 	return res
// }

// func subprojectsDeleteOneOperator(root string) *testresult.TestResult {
// 	res := &testresult.TestResult{
// 		Suite:   "endpoints",
// 		Element: "subprojects/{id}",
// 		ID:      "DELETE (operator)",
// 	}

// 	url := root + "/subprojects/2"

// 	// try and fail to delete the subproject
// 	res.Wanted = `{"error": "Access denied"}`
// 	err := utils.Delete(res, "1", url, ``, 403, "operator")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsMatch(res) {
// 		utils.FailMatch(res, "2")
// 		return res
// 	}

// 	// now, confirm that the subproject has NOT been deleted
// 	allURL := root + "/subprojects"
// 	res.Wanted = `{"subprojects":[{"id":1,"project_id":2,"name":"blorple","fullname":"The blorple Subproject"},{"id":2,"project_id":2,"name":"filfre","fullname":"The filfre Subproject"},{"id":3,"project_id":2,"name":"fweep","fullname":"The fweep Subproject"},{"id":4,"project_id":3,"name":"girgol","fullname":"The girgol Subproject"}]}`
// 	err = utils.GetContent(res, "3", allURL, 200, "viewer")
// 	if err != nil {
// 		return res
// 	}

// 	if !utils.IsMatch(res) {
// 		utils.FailMatch(res, "4")
// 		return res
// 	}

// 	utils.Pass(res)
// 	return res
// }
