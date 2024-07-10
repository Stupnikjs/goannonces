package gstore

import (
	"testing"
)

var TestDeleteBucketName = "somedeletebuckettestname"

func TestListObjectsBucket(t *testing.T) {

	data := []byte("this is test files content")

	err := LoadToBucket(TestBucketName, "test.txt", data)

	// Call get bucket method
	if err != nil {
		t.Errorf(" error in getter object %s", err)
	}

	l, err := ListObjectsBucket(TestBucketName)

	if err != nil {

		t.Errorf("expected no error calling listobjectbucket but got %s", err.Error())

	}

	if l[0] != "test.txt" {
		t.Errorf("expected test.txt as file but got %s", l[0])
	}

}
