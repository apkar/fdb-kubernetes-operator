/*
 * start_restore.go
 *
 * This source file is part of the FoundationDB open source project
 *
 * Copyright 2020 Apple Inc. and the FoundationDB project authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controllers

import (
	ctx "context"
	"fmt"
	"strings"
	"time"

	fdbtypes "github.com/FoundationDB/fdb-kubernetes-operator/api/v1beta1"
)

// StartRestore provides a reconciliation step for starting a new restore.
type StartRestore struct {
}

// Reconcile runs the reconciler's work.
func (s StartRestore) Reconcile(r *FoundationDBRestoreReconciler, context ctx.Context, restore *fdbtypes.FoundationDBRestore) (bool, error) {
	adminClient, err := r.AdminClientForRestore(context, restore)
	if err != nil {
		return false, err
	}

	status, err := adminClient.GetRestoreStatus()
	fmt.Printf("JPB got restore status -%s-, -%s-\n", status, strings.TrimSpace(status))
	if err != nil {
		return false, err
	}

	if len(strings.TrimSpace(status)) == 0 {
		fmt.Printf("JPB starting restore for url %s\n", restore.Spec.BackupURL)
		err = adminClient.StartRestore(restore.Spec.BackupURL)
		if err != nil {
			return false, err
		}

		restore.Status.Running = true
		err = r.Status().Update(context, restore)
		if err != nil {
			return false, err
		}
	}
	fmt.Printf("JPB done with StartRestore\n")

	return true, nil
}

// RequeueAfter returns the delay before we should run the reconciliation
// again.
func (s StartRestore) RequeueAfter() time.Duration {
	return 0
}
