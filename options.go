package xaqt

import (
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// An Option is a function that performs some kind of configuration on
// the context.
// No current implementations return errors, but it is included so guards
// can be added as desired.
// Idea taken from https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
type option func(*Context) error

// defaultOptions provides some useful defaults if the user provides none.
func defaultOptions(c *Context) error {
	pwd, _ := os.Getwd()
	c.path = filepath.Join(pwd, "data")

	if runtime.GOOS == "darwin" {
		c.folder = "/tmp"
	}

	c.image = "frenata/xaqt-sandbox"

	c.timeout = time.Second * 5
	return nil
}

// Timeout configures how long evaluation should run before it is killed.
func Timeout(t time.Duration) option {
	return func(c *Context) error {
		c.timeout = t
		return nil
	}
}

// Image configures which docker image should be used for evaluation.
func Image(i string) option {
	return func(c *Context) error {
		c.image = i
		return nil
	}
}

// Path configures the folder with the execution script and "Payload" dir.
func Path(p string) option {
	return func(c *Context) error {
		c.path = p
		return nil
	}
}

// TargetFolder configures where the result directory should be created.
func TargetFolder(f string) option {
	return func(c *Context) error {
		c.folder = f
		return nil
	}
}