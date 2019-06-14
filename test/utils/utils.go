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

// addAuthHeader adds the appropriate auth token header to the
// request object, before it is sent. Including "none" as the
// username means that no token will be sent. The token values
// included here are JWT values for the signing key "keyForTesting"
// and it is intentional that they are used here -- but of course
// they should not be used in production in any way!
func addAuthHeader(res *testresult.TestResult, step string, req *http.Request, ghUsername string) {
	switch ghUsername {
	case "none":
		return
	case "admin":
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJnaXRodWIiOiJhZG1pbiJ9.3KnAxp1Tn7O8vHQXBReUy81qG7qfRPsxRXW8Wr68xfc")
	case "operator":
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJnaXRodWIiOiJvcGVyYXRvciJ9.v8xJrGfBKDj9OYF2G58NeV1sGfKNahr-OHzqCXetwUU")
	case "commenter":
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJnaXRodWIiOiJjb21tZW50ZXIifQ.PQDdHhSmjDs9sceGi54cT71ef2IVxiO_Yw0-_YDJ-i8")
	case "viewer":
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJnaXRodWIiOiJ2aWV3ZXIifQ.YQUkHNTbsfA3ldtfxhTkoFI8eHVhfbFLF5vkmOrFJZg")
	case "disabled":
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJnaXRodWIiOiJkaXNhYmxlZCJ9.mqdsZIPPEb1RmmdI1zO0elHFieHbzmleYdg06qRfVbQ")
	default:
		FailTest(res, step, fmt.Errorf("invalid username %s", ghUsername))
	}
}

// GetContent makes an HTTP GET call to the indicated URL.
// It checks whether the expected HTTP status code is returned;
// a different code is treated as a failure.
// On success, it reads the response body into a got byte slice
// and handles closing the body. On failure, it fills in the
// failure code in the TestResult and returns an error.
func GetContent(res *testresult.TestResult, step string, url string, code int, ghUsername string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		FailTest(res, step, err)
		return err
	}
	addAuthHeader(res, step, req, ghUsername)
	resp, err := client.Do(req)

	return helperGetContent(res, resp, step, code)
}

// GetContentNoFollow makes an HTTP GET call to the indicated
// URL, and will NOT follow redirects. It otherwise acts
// identically to GetContent.
func GetContentNoFollow(res *testresult.TestResult, step string, url string, code int, ghUsername string) error {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		FailTest(res, step, err)
		return err
	}
	addAuthHeader(res, step, req, ghUsername)
	resp, err := client.Do(req)
	if err != nil {
		FailTest(res, step, err)
		return err
	}

	return helperGetContent(res, resp, step, code)
}

// helperGetContent does the rest of the GetContent or
// GetContentNoFollow activities, after the decision is
// made on whether to follow any redirects.
func helperGetContent(res *testresult.TestResult, resp *http.Response, step string, code int) error {
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
