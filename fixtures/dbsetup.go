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
	createFuncs := []func(string) error{
		createUsers,
		createProjects,
		createSubprojects,
		createRepos,
		createRepoBranches,
		createRepoPulls,
		createAgents,
		createJobs,
	}

	for _, f := range createFuncs {
		err := f(root)
		if err != nil {
			return err
		}
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

func createRepos(root string) error {
	url := root + "/repos"

	calls := []struct {
		subprojectID uint32
		name         string
		address      string
	}{
		{2, "filfre-core", "https://example.com/filfre-core.git"},
		{2, "filfre-api", "https://example.com/filfre-api.git"},
		{1, "blorple-c", "https://example.com/blorple-c.git"},
		{4, "girgol", "https://example.com/girgol.git"},
	}

	for _, c := range calls {
		body := fmt.Sprintf(`{"subproject_id": %d, "name": "%s", "address": "%s"}`, c.subprojectID, c.name, c.address)
		err := utils.PostNoRes(url, body, 201, "operator")
		if err != nil {
			return err
		}
	}

	return nil
}

func createRepoBranches(root string) error {
	calls := []struct {
		repoID uint32
		branch string
	}{
		{2, "master"},
		{2, "dev"},
		{4, "master"},
		{2, "dev-2.1"},
		{1, "master"},
		{1, "testing"},
	}

	for _, c := range calls {
		url := fmt.Sprintf("%s/repos/%d/branches", root, c.repoID)
		body := fmt.Sprintf(`{"branch": "%s"}`, c.branch)
		err := utils.PostNoRes(url, body, 201, "operator")
		if err != nil {
			return err
		}
	}

	return nil
}

func createRepoPulls(root string) error {
	calls := []struct {
		repoID uint32
		branch string
		vType  string
		v      string
	}{
		{1, "master", "commit", "22337864e74c9f54b1da4a64aaf7587ffa788039"},
		{2, "dev-2.1", "commit", "7864e74c9f54b1da4a64aaf7587ffa7880392233"},
		{2, "dev", "commit", "e74c9f54b1da4a64aaf7587ffa78803922337864"},
		{2, "dev-2.1", "commit", "9f54b1da4a64aaf7587ffa78803922337864e74c"},
		{1, "testing", "commit", "b1da4a64aaf7587ffa78803922337864e74c9f54"},
	}

	for _, c := range calls {
		url := fmt.Sprintf("%s/repos/%d/branches/%s", root, c.repoID, c.branch)
		body := fmt.Sprintf(`{"%s": "%s"}`, c.vType, c.v)
		err := utils.PostNoRes(url, body, 201, "operator")
		if err != nil {
			return err
		}
	}

	return nil
}

func createAgents(root string) error {
	url := root + "/agents"

	calls := []struct {
		name         string
		isActive     bool
		address      string
		port         int
		isCodeReader bool
		isSpdxReader bool
		isCodeWriter bool
		isSpdxWriter bool
	}{
		{"do-magic", true, "https://example.com/do-magic", 2087, false, true, false, false},
		{"read-magic", true, "https://example.com/read-magic", 2088, true, true, false, true},
		{"disabled", false, "localhost", 2057, false, true, false, false},
		{"wevs", true, "localhost", 5010, true, true, true, false},
	}

	for _, c := range calls {
		body := fmt.Sprintf(`{"name":"%s", "is_active":%t, "address":"%s", "port":%d, "is_codereader":%t, "is_spdxreader":%t, "is_codewriter":%t, "is_spdxwriter":%t}`, c.name, c.isActive, c.address, c.port, c.isCodeReader, c.isSpdxReader, c.isCodeWriter, c.isSpdxWriter)
		err := utils.PostNoRes(url, body, 201, "operator")
		if err != nil {
			return err
		}
	}

	return nil
}

func createJobs(root string) error {
	calls := []struct {
		repopullID  uint32
		agentID     uint32
		priorJobIDs string
		isReady     bool
		config      string
	}{
		{2, 1, "[]", true, `{"kv": {"hi":"steve"}}`},
		{4, 1, "[]", true, `{}`},
		{4, 2, "[2]", true, `{"codereader": {"primary": {"path": "/somewhere"}}}`},
		{4, 4, "[2,3]", false, `{"kv": {"hello":"world"}, "codereader": {"godeps": {"priorjob_id": 3}}, "spdxreader": {"primary": {"path": "/path/wherever"}, "godeps": {"priorjob_id": 3}}}`},
	}

	for _, c := range calls {
		url := fmt.Sprintf("%s/repopulls/%d/jobs", root, c.repopullID)
		body := fmt.Sprintf(`{"agent_id": %d, "priorjob_ids": %s, "is_ready": %t, "config": %s}`, c.agentID, c.priorJobIDs, c.isReady, c.config)
		err := utils.PostNoRes(url, body, 201, "operator")
		if err != nil {
			return err
		}
	}

	return nil
}
