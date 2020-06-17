```
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
```


```
go doc  exec.Command
package exec // import "os/exec"

func Command(name string, arg ...string) *Cmd
```



```
go doc  exec.Cmd   | grep -vE  '//|^$'   
type Cmd struct {
	Path string
	Args []string
	Env []string
	Dir string
	Stdin io.Reader
	Stdout io.Writer
	Stderr io.Writer
	ExtraFiles []*os.File
	SysProcAttr *syscall.SysProcAttr
	Process *os.Process
	ProcessState *os.ProcessState
}
    Cmd represents an external command being prepared or run.
    A Cmd cannot be reused after calling its Run, Output or CombinedOutput
    methods.
func Command(name string, arg ...string) *Cmd
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
func (c *Cmd) CombinedOutput() ([]byte, error)
func (c *Cmd) Output() ([]byte, error)
func (c *Cmd) Run() error
func (c *Cmd) Start() error
func (c *Cmd) StderrPipe() (io.ReadCloser, error)
func (c *Cmd) StdinPipe() (io.WriteCloser, error)
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
func (c *Cmd) String() string
func (c *Cmd) Wait() error
```


```
go doc  exec.Cmd.CombinedOutput | grep -vE '//|^$'
func (c *Cmd) CombinedOutput() ([]byte, error)
    CombinedOutput runs the command and returns its combined standard output and
    standard error.
```














