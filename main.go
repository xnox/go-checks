package main

import "crypto/md5"
import "fmt"
import "os/exec"

func main() {
	fmt.Printf("Running `echo Hello | openssl md5`\n")
	cmd := exec.Command("sh", "-c", "echo Hello | openssl md5")
	output, _ := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)

	fmt.Printf("Running `python3 -c 'print(hashlib.md5(b\"Hello\\n\").hexdigest())'`\n")
	cmd = exec.Command("python3", "-c", "import hashlib; print(hashlib.md5(b'Hello\\n').hexdigest())")
	output, _ = cmd.CombinedOutput()
	fmt.Printf("%s\n", output)

	fmt.Printf("Calculating `md5.Sum([]byte(\"Hello\\n\"))`\n")
	fmt.Printf("MD5(go)=    %x\n", md5.Sum([]byte("Hello\n")))
}
