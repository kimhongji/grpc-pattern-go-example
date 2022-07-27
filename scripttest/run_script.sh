#!/bin/bash

# script의 pid 및 시작시간 기록
echo $$ > script_pid
date --rfc-3339=ns > start_timestamp

# cmd 실행 및 대기
$1 1>stdout 2>stderr &
CMD_PID=$!
echo $CMD_PID > cmd_pid
wait $CMD_PID

# 종료되면 exit code 및 종료시간 기록
EXIT_CODE=$?
echo $EXIT_CODE > exit_code
date --rfc-3339=ns > end_timestamp