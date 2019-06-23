// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package endpoints

import (
	"github.com/swinslow/peridot-api-testing/internal/testresult"
	"github.com/swinslow/peridot-api-testing/test/utils"
)

func getSubprojectsTests() []testresult.TestFunc {
	return []testresult.TestFunc{
		subprojectsGetViewer,
		subprojectsPostOperator,
	}
}

// ===== GET /projects

func subprojectsGetViewer(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "subprojects",
		ID:      "GET (viewer)",
	}

	url := root + "/subprojects"

	res.Wanted = `{"subprojects":[{"id":1,"project_id":2,"name":"blorple","fullname":"The blorple Subproject"},{"id":2,"project_id":2,"name":"filfre","fullname":"The filfre Subproject"},{"id":3,"project_id":2,"name":"fweep","fullname":"The fweep Subproject"},{"id":4,"project_id":3,"name":"girgol","fullname":"The girgol Subproject"}]}`
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

// ===== POST /subprojects

func subprojectsPostOperator(root string) *testresult.TestResult {
	res := &testresult.TestResult{
		Suite:   "endpoints",
		Element: "subprojects",
		ID:      "POST (operator)",
	}

	url := root + "/subprojects"

	// first, send POST to add a new project
	body := `{"project_id": 3, "name": "plugh", "fullname": "The plugh Subproject"}`
	res.Wanted = `{"success": true, "id": 5}`
	err := utils.Post(res, "1", url, body, 201, "operator")
	if err != nil {
		return res
	}

	if !utils.IsMatch(res) {
		utils.FailMatch(res, "2")
		return res
	}

	// now, confirm that a new project was actually added
	res.Wanted = `{"subprojects":[{"id":1,"project_id":2,"name":"blorple","fullname":"The blorple Subproject"},{"id":2,"project_id":2,"name":"filfre","fullname":"The filfre Subproject"},{"id":3,"project_id":2,"name":"fweep","fullname":"The fweep Subproject"},{"id":4,"project_id":3,"name":"girgol","fullname":"The girgol Subproject"},{"id": 5, "project_id": 3, "name": "plugh", "fullname": "The plugh Subproject"}]}`
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
