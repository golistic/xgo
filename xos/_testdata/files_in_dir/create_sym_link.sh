#!/bin/sh

#
# Copyright (c) 2023, Geert JM Vanderkelen
#

dir=$(dirname "$0")
cd "${dir}" || exit 1
rm symlink.md
ln -s fileA.md symlink.md
