// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package plugin

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/go-plugin"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/opentf"
)

func TestUIInput_impl(t *testing.T) {
	var _ opentf.UIInput = new(UIInput)
}

func TestUIInput_input(t *testing.T) {
	client, server := plugin.TestRPCConn(t)
	defer client.Close()

	i := new(opentf.MockUIInput)
	i.InputReturnString = "foo"

	err := server.RegisterName("Plugin", &UIInputServer{
		UIInput: i,
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	input := &UIInput{Client: client}

	opts := &opentf.InputOpts{
		Id: "foo",
	}

	v, err := input.Input(context.Background(), opts)
	if !i.InputCalled {
		t.Fatal("input should be called")
	}
	if !reflect.DeepEqual(i.InputOpts, opts) {
		t.Fatalf("bad: %#v", i.InputOpts)
	}
	if err != nil {
		t.Fatalf("bad: %#v", err)
	}

	if v != "foo" {
		t.Fatalf("bad: %#v", v)
	}
}
