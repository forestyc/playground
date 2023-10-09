#!/bin/bash

tag_cout=$(git tag | wc -l)

# 第一次集成, 直接统计总数
if [ ${tag_cout} -lt 2 ]; then
  cloc .
else
  last_sha=$(git log --pretty=oneline `git tag --sort=taggerdate | sed -n $((tag_cout - 1))p` | sed -n 1p | awk -F ' ' '{print $1}')
  latest_sha=$(git log --pretty=oneline `git tag --sort=taggerdate | sed -n ${tag_cout}p` | sed -n 1p | awk -F ' ' '{print $1}')
  cloc -diff ${last_sha} ${latest_sha}
fi