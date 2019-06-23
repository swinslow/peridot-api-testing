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

	if resp.StatusCode != 204 {
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return fmt.Errorf("expected 204, got %d from resetDB command: %s; with error reading response body: %v", resp.StatusCode, string(b), err)
		}
		return fmt.Errorf("expected 204, got %d from resetDB command: %s", resp.StatusCode, string(b))
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

	calls := []struct {
		name   string
		github string
		access string
	}{
		{"Operator User", "operator", "operator"},
		{"Commenter User", "commenter", "commenter"},
		{"Viewer User", "viewer", "viewer"},
		{"Disabled User", "disabled", "disabled"},
	}

	for _, c := range calls {
		body := fmt.Sprintf(`{"name": "%s", "github": "%s", "access": "%s"}`, c.name, c.github, c.access)
		err := utils.PostNoRes(url, body, 201, "admin")
		if err != nil {
			return err
		}
	}

	return nil
}

func createProjects(root string) error {
	url := root + "/projects"

	calls := []struct {
		name     string
		fullname string
	}{
		{"xyzzy", "The xyzzy Project"},
		{"frotz", "The frotz Project"},
		{"gnusto", "The gnusto Project"},
	}

	for _, c := range calls {
		body := fmt.Sprintf(`{"name": "%s", "fullname": "%s"}`, c.name, c.fullname)
		err := utils.PostNoRes(url, body, 201, "operator")
		if err != nil {
			return err
		}
	}

	return nil
}

func createSubprojects(root string) error {
	url := root + "/subprojects"

	calls := []struct {
		projectID uint32
		name      string
		fullname  string
	}{
		{2, "blorple", "The blorple Subproject"},
		{2, "filfre", "The filfre Subproject"},
		{2, "fweep", "The fweep Subproject"},
		{3, "girgol", "The girgol Subproject"},
	}

	for _, c := range calls {
		body := fmt.Sprintf(`{"project_id": %d, "name": "%s", "fullname": "%s"}`, c.projectID, c.name, c.fullname)
		err := utils.PostNoRes(url, body, 201, "operator")
		if err != nil {
			return err
		}
	}

	return nil
}
