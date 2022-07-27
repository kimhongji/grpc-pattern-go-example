package main

import "os/exec"

func main()  {

	cmd := exec.Command("nohup", "sh", "/Users/kakao_ent/git_hongji/grpc-pattern-go-example/scripttest/run_script.sh", "sleep 100", "> /dev/null", "&")
	cmd.Dir = "/Users/kakao_ent/git_hongji/grpc-pattern-go-example/scripttest/testDir"
	err := cmd.Start()
	if err != nil {
		return
	}
}
