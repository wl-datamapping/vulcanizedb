/*
 * Copyright 2018 Dgraph Labs, Inc. and Contributors
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

package cmd

import (
	"gx/ipfs/QmU4emVTYFKnoJ5yK3pPEN9joyEx6U7y892PDx26ZtNxQd/badger"
	"gx/ipfs/QmXj63M2w2Pq7mnBpcrs7Va8prmfhvfMUNqVhJ9TgjiMbT/cobra"
)

var flattenCmd = &cobra.Command{
	Use:   "flatten",
	Short: "Flatten the LSM tree.",
	Long: `
This command would compact all the LSM tables into one level.
`,
	RunE: flatten,
}

var numWorkers int

func init() {
	RootCmd.AddCommand(flattenCmd)
	flattenCmd.Flags().IntVarP(&numWorkers, "num-workers", "w", 1,
		"Number of concurrent compactors to run. More compactors would use more"+
			" server resources to potentially achieve faster compactions.")
}

func flatten(cmd *cobra.Command, args []string) error {
	opts := badger.DefaultOptions
	opts.Dir = sstDir
	opts.ValueDir = vlogDir
	opts.Truncate = truncate
	opts.NumCompactors = 0

	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Flatten(numWorkers)
}
