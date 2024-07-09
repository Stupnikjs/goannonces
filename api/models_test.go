package api

import (
	"net/http/httptest"
	"testing"
)

func TestParseJsonReq(t *testing.T) {

	_ = `{
                "action": "update",
                "object": {
                    "type" : "track",
                    "id": trackid,
                    "field" : "tag",
                    "body": input.value
                }
            }`

	// create mockrequest
	req := httptest.NewRequest("POST", "/", nil)

	// pass json to request

	jsonReq, err := ParseJsonReq(req)

	if err != nil {
		t.Errorf("expected no error but got %s", err.Error())
	}

	if jsonReq.Action != "update" {
		t.Errorf("expected update action but go %s", jsonReq.Action)
	}

}
