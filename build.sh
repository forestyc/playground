#!/bin/bash

set -e

if [ $# -ne 4 ];then
  echo "usage: $0 <branch> <tag> <commit-id> <out_dir>"
  exit 1
fi

branch=$1
tag=$2
commit_id=$3
out_dir=$4

date=$(date +%FT%T%z)
version_pkg="fipdev.dfitc.com.cn/portal/golang/pkg/version"
# common-ldflags
branch_flag="${version_pkg}.gitBranch=${branch}"
tag_flag="${version_pkg}.gitTag=${tag}"
commit_id_flag="${version_pkg}.gitCommit=${commit_id}"
build_date_flag="${version_pkg}.buildDate=${date}"
ldflags="-X ${branch_flag} -X ${tag_flag} -X ${commit_id_flag} -X ${build_date_flag}"

# build
if [ ! -d ~/.golang_version ]; then
  mkdir ~/.golang_version
fi
dir=$(cd $(dirname $0); pwd)
cmds=$(ls ${dir}/cmd)
for cmd in $cmds
do
  echo "build ${cmd}..."
  cmd_dir="${out_dir}/${cmd}"
  mkdir -p ${cmd_dir}
  if [ -f ~/.golang_version/${cmd} ];then
    last_version=$(cat ~/.golang_version/${cmd} 2>/dev/null)
  fi
  cd "${dir}/cmd/${cmd}"
  latest_version=$(cat CHANGELOG.md 2>/dev/null |sed '/^$/d' | head -1 | awk -F ' ' '{print $2}')
  if [ "$last_version" == "$latest_version" ];then
    echo "skip ${cmd}"
    continue
  fi
  ldflags=${ldflags}" -X ${version_pkg}.version=${latest_version}"
  # binary
  echo go build -o ${cmd_dir}/${cmd} -ldflags '"'${ldflags}'"' > ./build_cmd
  source ./build_cmd
  rm ./build_cmd
  # etc
  cp etc ${cmd_dir} -fr
  echo $latest_version > ~/.golang_version/${cmd}
  echo "build ${cmd} complete"
  cd - >/dev/null
done