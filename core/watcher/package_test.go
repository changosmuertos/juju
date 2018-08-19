// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package watcher_test

import (
	stdtesting "testing"

	gc "gopkg.in/check.v1"

	coretesting "github.com/juju/juju/testing"
)

func TestPackage(t *stdtesting.T) {
	gc.TestingT(t)
}

type ImportTest struct{}

var _ = gc.Suite(&ImportTest{})

func (*ImportTest) TestImports(c *gc.C) {
	found := coretesting.FindJujuCoreImports(c, "github.com/juju/juju/core/watcher")

	// This package brings in nothing else from juju/juju
	c.Assert(found, gc.HasLen, 0)
}
