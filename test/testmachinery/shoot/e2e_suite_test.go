// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package shoot_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestShoot(t *testing.T) {
	// TODO (plkokanov): add testmachinery tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shoot Suite")
}
