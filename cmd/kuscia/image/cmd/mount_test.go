// Copyright 2024 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/secretflow/kuscia/cmd/kuscia/utils"
	storetesting "github.com/secretflow/kuscia/pkg/agent/local/store/testing"
)

func TestMountCommand(t *testing.T) {
	rootDir := t.TempDir()

	args := []string{"123", "app:0.1", filepath.Join(rootDir, "rootfs")}

	err := utils.MountImage(storetesting.NewFakeStore(), filepath.Join(rootDir, "storage"), args[0], args[1], args[2])
	assert.NoError(t, err)
}
