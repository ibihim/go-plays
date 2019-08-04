#!/usr/bin/env bash

# Copyright 2012 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# git gofmt pre-commit hook
#
# To use, store as .git/hooks/pre-commit inside your repository and make sure
# it has execute permissions.
#
# This script does not handle file names that contain spaces.

# git diff shows difference between commits.
# --cached is a synonym for --staged and it used to show staged files.
# --name-only shows only the names of the files.
# \.go$ is a pattern that selects only results that end with '.go'.
STAGED_GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')
[ -z "$STAGED_GO_FILES" ] && exit 0

for FILE in $STAGED_GO_FILES; do
    gofmt -w $FILE
    goimports -w $FILE

    golint "-set_exit_status" $FILE
    if [[ $? == 1 ]]; then
        printf "golint erred\n"
        exit 1
    fi

    go vet $FILE
    if [[ $? != 0 ]]; then
        printf "go vet failed\n"
    fi
done

exit 0
