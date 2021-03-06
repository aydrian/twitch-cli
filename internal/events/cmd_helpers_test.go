// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package events

import (
	"testing"

	"github.com/twitchdev/twitch-cli/internal/util"
)

func TestValidTriggers(t *testing.T) {
	a := util.SetupTestEnv(t)

	t1 := ValidTriggers()
	a.NotEmpty(t1)
}

func TestValidTransports(t *testing.T) {
	a := util.SetupTestEnv(t)

	t1 := ValidTransports()
	a.NotEmpty(t1)
}
