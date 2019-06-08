package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/yudai/gojsondiff"
	"github.com/swinslow/obsidian-api-testing/internal/testresult"
)

// Pass fills in the success fields.
func Pass(res *testresult.TestResult) {
	res.Success = true
}

// FailTest fills in the failure fields for a test that failed
// for some reason other than because the JSON strings did not
// match.
func FailTest(res *testresult.TestResult, step float32, msg error) {
	res.Success = false
	res.FailStep = step
	res.FailError = msg
}

// FailMatch fills in the failure fields for a test that failed
// because the desired JSON string did not match the JSON string
// that was received.
func FailMatch(res *testresult.TestResult, step float32, wanted string, got []byte) {
	res.Success = false
	res.FailStep = step
	res.FailWanted = wanted
	res.FailGot = string(got)
}

// IsMatch compares a wanted string and a got byte slice containing
// JSON data, and returns a bool indicating whether they contained
// equivalent content. It will also return "false" if there is e.g.
// an error with the JSON unmarshalling, etc.
func IsMatch(wanted string, got []byte) bool {
	differ := gojsondiff.New()
	d, err := differ.Compare([]byte(wanted), got)
	if err != nil {
		return false
	}

	return d.Modified()
}

// GetContent makes an HTTP GET call to the indicated URL.
// On success, it reads the response body into a got byte slice
// and handles closing the body. On failure, it fills in the
// failure code in the TestResult and returns an error.
func GetContent(res *testresult.TestResult, step float32, url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		FailTest(res, step, err)
		return []byte{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		FailTest(res, step, err)
		return []byte{}, err
	}

	return b, nil
}