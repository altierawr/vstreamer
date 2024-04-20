package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
)

type Transcoder struct {
	Filepath         string
	OutputDir        string
	stateChangeMutex sync.Mutex
	processWaitGroup sync.WaitGroup
	cmd              *exec.Cmd
}

func (t *Transcoder) Start() error {
	t.processWaitGroup.Add(1)
	t.stateChangeMutex.Lock()
	defer t.stateChangeMutex.Unlock()

	// Give the ffmpeg process its own group id so we can terminate it separate from the app's process
	t.cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	fmt.Printf("starting transcode session: ffmpeg %s\n", strings.Join(t.cmd.Args, " "))
	err := t.cmd.Start()
	if err != nil {
		return err
	}

	go func() {
		t.cmd.Wait()
		t.processWaitGroup.Done()
	}()

	return nil
}

func (s *Transcoder) Destroy() error {
	fmt.Printf("destroying transcode session: %s\n", s.OutputDir)
	pgid, err := syscall.Getpgid(s.cmd.Process.Pid)
	if err != nil {
		// Probably the process was already terminated
		fmt.Printf("error getting transcode session group id: %v\n", err)
	} else {
		syscall.Kill(-pgid, syscall.SIGTERM)
	}

	// Clean output dir once the process is done
	s.processWaitGroup.Wait()
	err = os.RemoveAll(s.OutputDir)
	fmt.Printf("destroyed transcode session: %s\n", s.OutputDir)

	return err
}
