#!/usr/bin/env bash
# This script is used to bump the version of the operator.
# It uses semtag to bump the version.
#
# Usage: hack/bump.sh <stage> <scope> <dryrun>
# stage: major, minor, patch, final
# scope:

# tagging before release, will write VERSION file and generate tag
# when writing VERSION file, it will force to add `-dev` suffix, but push tag will not have `-dev`
# if your want release with no `-dev` suffix, you should sed the VERSION file before push tag

# cd root of the repo
pushd "$(dirname "$0")/.." > /dev/null

set -e

bump_stage=$1
bump_scope=$2
bump_dry_run=$3

if [ -z "$bump_stage" ]; then
  echo "bump stage is required"
  exit 1
fi
if [ -z "$bump_scope" ]; then
  echo "bump scope is required"
  exit 1
fi
if [ -z "$bump_dry_run" ]; then
  echo "bump dryrun is required"
  exit 1
fi

next=$(semtag "$bump_stage" -s "$bump_scope" -f -o)
echo "next version: $next"

# dry run
if [ "$bump_dry_run" = "true" ]; then
  echo "dryrun: true"
  exit 0
fi

# bump and tag
echo "dryrun: false"
# VERSION in file always has the -dev suffix
echo "${next}-dev" > VERSION
git add VERSION
git commit -m "chore: Bump version to $next"

# git tag did not contains dev suffix
semtag "$bump_stage" -s "$bump_scope"

popd > /dev/null
