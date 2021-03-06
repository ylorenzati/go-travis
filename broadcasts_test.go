// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBroadcastsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/broadcasts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"active": "true", "include": "broadcast.recipient"})
		fmt.Fprint(w, `{"broadcasts": [{"id":125,"message":"We just switched the default image for","active":true,"created_at":"2014-11-19T14:39:51Z"}]}`)
	})

	broadcasts, _, err := client.Broadcasts.List(context.Background(), &BroadcastsOption{Active: true, Include: []string{"broadcast.recipient"}})

	if err != nil {
		t.Errorf("Broadcasts.List returned error: %v", err)
	}

	want := []*Broadcast{{Id: Uint(125), Message: String("We just switched the default image for"), Active: Bool(true), CreatedAt: String("2014-11-19T14:39:51Z")}}
	if !reflect.DeepEqual(broadcasts, want) {
		t.Errorf("Broadcasts.List returned %+v, want %+v", broadcasts, want)
	}
}
