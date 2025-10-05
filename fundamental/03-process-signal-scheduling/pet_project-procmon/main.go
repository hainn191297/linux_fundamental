package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("usage: procmon -- <command> [args...]")
		os.Exit(2)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("parent PID=%d\n", os.Getpid())
	start := time.Now()

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "start error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("child PID=%d\n", cmd.Process.Pid)

	// graceful shutdown on SIGINT/SIGTERM
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()

	select {
	case s := <-sigc:
		fmt.Printf("received signal: %s, terminating child...\n", s)
		// try graceful
		_ = cmd.Process.Signal(syscall.SIGTERM)
		// if still running after a timeout, force kill
		select {
		case err := <-done:
			reportResult(err, start)
		case <-time.After(2 * time.Second):
			_ = cmd.Process.Kill()
			fmt.Println("child killed")
		}
	case err := <-done:
		reportResult(err, start)
	}
}

func reportResult(err error, start time.Time) {
	dur := time.Since(start)
	if err == nil {
		fmt.Printf("exit ok | duration=%s\n", dur)
		return
	}
	// try to extract exit status
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			if status.Signaled() {
				fmt.Printf("exited by signal=%s | duration=%s\n", status.Signal(), dur)
				return
			}
			fmt.Printf("exit code=%d | duration=%s\n", status.ExitStatus(), dur)
			return
		}
	}
	fmt.Printf("exit error=%v | duration=%s\n", err, dur)
}
