// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package fixtures

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/swinslow/peridot-api-testing/test/utils"
)

// ResetDB asks the database to re-initialize itself to a
// initial clean state. Only the initial github admin user
// will be set.
func ResetDB(root string) error {
	resetCommand := `{"command": "resetDB"}`
	client := &http.Client{}
	req, err := http.NewRequest("POST", root+"/admin/db", strings.NewReader(resetCommand))
	if err != nil {
		return fmt.Errorf("got error from resetDB http request creator: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	utils.AddAuthHeader(nil, "", req, "admin")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	bodystr := string(b)

	if bodystr != `{"success": true}` {
		return fmt.Errorf("got error from resetDB command: %s", bodystr)
	}

	return nil
}

// SetupFixture makes calls to the peridot API to create
// objects in its database, so that it is in a useful
// state for functional tests.
func SetupFixture(root string) error {
	err := createUsers(root)
	if err != nil {
		return err
	}

	err = createProjects(root)
	if err != nil {
		return err
	}

	err = createSubprojects(root)
	if err != nil {
		return err
	}

	return nil
}

func createUsers(root string) error {
	url := root + "/users"

	// ID 1, name "Admin", github "admin", access "admin" is
	// created by default on creation and each reset

	// add operator
	body := `{"name": "Operator User", "github": "operator", "access": "operator"}`
	err := utils.PostNoRes(url, body, 201, "admin")
	if err != nil {
		return err
	}

	// add commenter
	body = `{"name": "Commenter User", "github": "commenter", "access": "commenter"}`
	err = utils.PostNoRes(url, body, 201, "admin")
	if err != nil {
		return err
	}

	// add viewer
	body = `{"name": "Viewer User", "github": "viewer", "access": "viewer"}`
	err = utils.PostNoRes(url, body, 201, "admin")
	if err != nil {
		return err
	}

	// add disabled
	body = `{"name": "Disabled User", "github": "disabled", "access": "disabled"}`
	err = utils.PostNoRes(url, body, 201, "admin")
	if err != nil {
		return err
	}

	return nil
}

func createProjects(root string) error {
	url := root + "/projects"

	body := `{"name": "xyzzy", "fullname": "The xyzzy Project"}`
	err := utils.PostNoRes(url, body, 201, "operator")
	if err != nil {
		return err
	}

	body = `{"name": "frotz", "fullname": "The frotz Project"}`
	err = utils.PostNoRes(url, body, 201, "operator")
	if err != nil {
		return err
	}

	body = `{"name": "gnusto", "fullname": "The gnusto Project"}`
	err = utils.PostNoRes(url, body, 201, "operator")
	if err != nil {
		return err
	}

	return nil
}

func createSubprojects(root string) error {
	url := root + "/subprojects"

	body := `{"project_id": 2, "name": "blorple", "fullname": "The blorple Subproject"}`
	err := utils.PostNoRes(url, body, 201, "operator")
	if err != nil {
		return err
	}

	body = `{"project_id": 2, "name": "filfre", "fullname": "The filfre Subproject"}`
	err = utils.PostNoRes(url, body, 201, "operator")
	if err != nil {
		return err
	}

	body = `{"project_id": 2, "name": "fweep", "fullname": "The fweep Subproject"}`
	err = utils.PostNoRes(url, body, 201, "operator")
	if err != nil {
		return err
	}

	body = `{"project_id": 3, "name": "girgol", "fullname": "The girgol Subproject"}`
	err = utils.PostNoRes(url, body, 201, "operator")
	if err != nil {
		return err
	}

	return nil
}
