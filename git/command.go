// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Command contains the name, arguments and environment variables of a command.
type Command struct {
	name string
	args []string
	envs []string
}

// String returns the string representation of the command.
func (c *Command) String() string {
	if len(c.args) == 0 {
		return c.name
	}
	return fmt.Sprintf("%s %s", c.name, strings.Join(c.args, " "))
}

// NewCommand creates and returns a new Command with given arguments for "git".
func NewCommand(args ...string) *Command {
	return &Command{
		name: "git",
		args: args,
	}
}

// AddArgs appends given arguments to the command.
func (c *Command) AddArgs(args ...string) *Command {
	c.args = append(c.args, args...)
	return c
}

// AddEnvs appends given environment variables to the command.
func (c *Command) AddEnvs(envs ...string) *Command {
	c.envs = append(c.envs, envs...)
	return c
}

// DefaultTimeout is the default timeout duration for all commands.
const DefaultTimeout = time.Minute

// A limitDualWriter writes to W but limits the amount of data written to just N bytes.
// On the other hand, it passes everything to w.
type limitDualWriter struct {
	W        io.Writer // underlying writer
	N        int64     // max bytes remaining
	prompted bool

	w io.Writer
}

func (w *limitDualWriter) Write(p []byte) (int, error) {
	if w.N > 0 {
		limit := int64(len(p))
		if limit > w.N {
			limit = w.N
		}
		n, _ := w.W.Write(p[:limit])
		w.N -= int64(n)
	}

	if !w.prompted && w.N <= 0 {
		w.prompted = true
		_, _ = w.W.Write([]byte("... (more omitted)"))
	}

	return w.w.Write(p)
}

// RunInDirPipelineWithTimeout executes the command in given directory and timeout duration.
// It pipes stdout and stderr to supplied io.Writer. DefaultTimeout will be used if the timeout
// duration is less than time.Nanosecond (i.e. less than or equal to 0).
// It returns an ErrExecTimeout if the execution was timed out.
func (c *Command) RunInDirPipelineWithTimeout(timeout time.Duration, stdout, stderr io.Writer, dir string) (err error) {
	if timeout < time.Nanosecond {
		timeout = DefaultTimeout
	}

	buf := new(bytes.Buffer)
	w := stdout
	if logOutput != nil {
		buf.Grow(512)
		w = &limitDualWriter{
			W: buf,
			N: int64(buf.Cap()),
			w: stdout,
		}
	}

	defer func() {
		if len(dir) == 0 {
			log("[timeout: %v] %s\n%s", timeout, c, buf.Bytes())
		} else {
			log("[timeout: %v] %s: %s\n%s", timeout, dir, c, buf.Bytes())
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer func() {
		cancel()
		if err == context.DeadlineExceeded {
			err = ErrExecTimeout
		}
	}()

	cmd := exec.CommandContext(ctx, c.name, c.args...)
	if len(c.envs) > 0 {
		cmd.Env = append(os.Environ(), c.envs...)
	}
	cmd.Dir = dir
	cmd.Stdout = w
	cmd.Stderr = stderr
	// fmt.Printf("\r[执行命令]%s\r",cmd)
	if err = cmd.Start(); err != nil {
		return err
	}

	result := make(chan error)
	go func() {
		result <- cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		<-result
		if cmd.Process != nil && cmd.ProcessState != nil && !cmd.ProcessState.Exited() {
			if err := cmd.Process.Kill(); err != nil {
				return fmt.Errorf("kill process: %v", err)
			}
		}

		return ErrExecTimeout
	case err = <-result:
		return err
	}
}

// RunInDirPipeline executes the command in given directory and default timeout duration.
// It pipes stdout and stderr to supplied io.Writer.
func (c *Command) RunInDirPipeline(stdout, stderr io.Writer, dir string) error {
	return c.RunInDirPipelineWithTimeout(DefaultTimeout, stdout, stderr, dir)
}

// RunInDirWithTimeout executes the command in given directory and timeout duration.
// It returns stdout in []byte and error (combined with stderr).
func (c *Command) RunInDirWithTimeout(timeout time.Duration, dir string) ([]byte, error) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	if err := c.RunInDirPipelineWithTimeout(timeout, stdout, stderr, dir); err != nil {
		return nil, concatenateError(err, stderr.String())
	}
	return stdout.Bytes(), nil
}

// RunInDir executes the command in given directory and default timeout duration.
// It returns stdout and error (combined with stderr).
func (c *Command) RunInDir(dir string) ([]byte, error) {
	return c.RunInDirWithTimeout(DefaultTimeout, dir)
}

// RunWithTimeout executes the command in working directory and given timeout duration.
// It returns stdout in string and error (combined with stderr).
func (c *Command) RunWithTimeout(timeout time.Duration) ([]byte, error) {
	stdout, err := c.RunInDirWithTimeout(timeout, "")
	if err != nil {
		return nil, err
	}
	return stdout, nil
}

// Run executes the command in working directory and default timeout duration.
// It returns stdout in string and error (combined with stderr).
func (c *Command) Run() ([]byte, error) {
	return c.RunWithTimeout(DefaultTimeout)
}
