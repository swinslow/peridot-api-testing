// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/swinslow/obsidian-api-testing/internal/testresult"
	"github.com/yudai/gojsondiff"
)

// Pass fills in the success fields.
func Pass(res *testresult.TestResult) {
	res.Success = true
}

// FailTest fills in the failure fields for a test that failed
// for some reason other than because the JSON strings did not
// match.
func FailTest(res *testresult.TestResult, step string, msg error) {
	res.Success = false
	res.FailStep = step
	res.FailError = msg
}

// FailMatch fills in the failure fields for a test that failed
// because the desired JSON string did not match the JSON string
// that was received.
func FailMatch(res *testresult.TestResult, step string) {
	res.Success = false
	res.FailStep = step
}

// IsMatch compares a wanted string and a got byte slice containing
// JSON data, and returns a bool indicating whether they contained
// equivalent content. It will also return "false" if there is e.g.
// an error with the JSON unmarshalling, etc.
func IsMatch(res *testresult.TestResult) bool {
	differ := gojsondiff.New()
	d, err := differ.Compare([]byte(res.Wanted), res.Got)
	// fmt.Printf("*** WANTED:  %#v\n", wanted)
	// fmt.Printf("*** GOT:     %#v\n", got)
	// fmt.Printf("*** ERR:     %#v\n", err)
	// fmt.Printf("*** DIFF:    %#v\n", d)
	if err != nil {
		return false
	}

	return !d.Modified()
}

// GetContent makes an HTTP GET call to the indicated URL.
// It checks whether the expected HTTP status code is returned;
// a different code is treated as a failure.
// On success, it reads the response body into a got byte slice
// and handles closing the body. On failure, it fills in the
// failure code in the TestResult and returns an error.
func GetContent(res *testresult.TestResult, step string, url string, code int) error {
	// make GET call
	resp, err := http.Get(url)
	if err != nil {
		FailTest(res, step, err)
		return err
	}

	// parse content body
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		FailTest(res, step, err)
		return err
	}

	// record in testresult
	res.Got = b

	// check expected status code
	if resp.StatusCode != code {
		err = fmt.Errorf("expected HTTP status code %d, got %d", code, resp.StatusCode)
		FailTest(res, step, err)
		return err
	}

	return nil
}
