// Copyright 2023 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shell

import (
	"yunion.io/x/cloudmux/pkg/multicloud/ksyun"
	"yunion.io/x/pkg/util/shellutils"
)

func init() {
	type PolicyListOptions struct {
		Scope string `choices:"System|Custom"`
	}
	shellutils.R(&PolicyListOptions{}, "cloud-policy-list", "List policies", func(cli *ksyun.SRegion, args *PolicyListOptions) error {
		ret, err := cli.GetClient().ListPolicies(args.Scope)
		if err != nil {
			return err
		}
		printList(ret)
		return nil
	})

}
