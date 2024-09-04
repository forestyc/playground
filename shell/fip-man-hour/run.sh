#!/usr/bin/env bash

if [ $# -ne 4 ]; then
  echo "$0 <token> <date> <task_id> <task_name>"
  exit 1
fi

token=$1
date=$2
task_id=$3
task_name=`echo $4 | xxd -plain | sed 's/\(..\)/%\1/g' | tr -d '\n'`

curl "http://124.93.32.64:9000/pm/ETHtmlInterfaceService" \
  -H "Accept: application/json, text/plain, */*" \
  -H "Accept-Language: zh-CN,zh;q=0.9" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -H "Origin: http://124.93.32.64:9000" \
  -H "Proxy-Connection: keep-alive" \
  -H "Referer: http://124.93.32.64:9000/pm/appH5/" \
  -H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36" \
  --data-raw "operation=saveOneTimeSheet&isSubmit=true&client=android&sheetDate=${date}&taskID=${task_id}&type=0&userID=1075&userType=&projectTaskID=208889&departmentTaskID=0&status=0&percent=6&bill=false&address=&saveType=daily&taskNameInput=${task_name}&hourRemarks=&productIds=266592&normal=7&userName=xnBVobZjQMisF4EqQHThkQ%3D%3D&companyID=%2BteY5l39jcEOHZkkKUnMsQ%3D%3D&token=${token}" \
  --insecure