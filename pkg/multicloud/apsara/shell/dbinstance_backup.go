// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shell

import (
	"yunion.io/x/cloudmux/pkg/multicloud/apsara"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type DBInstanceBackupJobListOptions struct {
		INSTANCE string
		JobId    string
	}
	shellutils.R(&DBInstanceBackupJobListOptions{}, "dbinstance-backup-job-list", "Get dbinstance backup jobs", func(cli *apsara.SRegion, args *DBInstanceBackupJobListOptions) error {
		jobs, err := cli.GetDBInstanceBackupJobs(args.INSTANCE, args.JobId)
		if err != nil {
			return err
		}
		printObject(jobs)
		return nil
	})

	type DBInstanceBackupOptions struct {
		INSTANCE string
		BACKUP   string
	}

	shellutils.R(&DBInstanceBackupOptions{}, "dbinstance-backup-delete", "Delete dbinstance backup", func(cli *apsara.SRegion, args *DBInstanceBackupOptions) error {
		return cli.DeleteDBInstanceBackup(args.INSTANCE, args.BACKUP)
	})

	shellutils.R(&DBInstanceBackupOptions{}, "dbinstance-backup-job-list", "Get dbinstance backup jobs", func(cli *apsara.SRegion, args *DBInstanceBackupOptions) error {
		return cli.DeleteDBInstanceBackup(args.INSTANCE, args.BACKUP)
	})

	type DBInstanceIdExtraOptions struct {
		ID     string `help:"ID of instances to show"`
		Limit  int    `help:"page size"`
		Offset int    `help:"page offset"`
	}

	shellutils.R(&DBInstanceIdExtraOptions{}, "dbinstance-backup-list", "List dbintance backups", func(cli *apsara.SRegion, args *DBInstanceIdExtraOptions) error {
		backups, _, err := cli.GetDBInstanceBackups(args.ID, "", args.Offset, args.Limit)
		if err != nil {
			return err
		}
		printList(backups, 0, 0, 0, []string{})
		return nil
	})

}
