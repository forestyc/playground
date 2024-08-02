#!/usr/bin/env bash

if [ $# -eq 3 ]; then
  echo "$0 <token> <date> <task_id>"
  exit 1
fi

token=$1
date=$2
task_id=$3

curl 'http://124.93.32.64:9000/pm/ETHtmlInterfaceService' \
  -H 'Accept: application/json, text/plain, */*' \
  -H 'Accept-Language: zh-CN,zh;q=0.9' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Origin: http://124.93.32.64:9000' \
  -H 'Proxy-Connection: keep-alive' \
  -H 'Referer: http://124.93.32.64:9000/pm/appH5/' \
  -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36' \
  --data-raw 'operation=saveOneTimeSheet&isSubmit=true&client=android&sheetDate=${date}&taskID=${task_id}&type=0&userID=1075&userType=&projectTaskID=208889&departmentTaskID=0&status=0&percent=6&bill=false&address=&saveType=daily&taskNameInput=%E4%BC%9A%E5%91%98-2024%E9%9C%80%E6%B1%82%E7%AE%A1%E7%90%86%E5%8F%8A%E5%88%86%E6%9E%90&hourRemarks=&productIds=266592&normal=7&userName=xnBVobZjQMisF4EqQHThkQ%3D%3D&companyID=%2BteY5l39jcEOHZkkKUnMsQ%3D%3D&token=${token}' \
  --insecure