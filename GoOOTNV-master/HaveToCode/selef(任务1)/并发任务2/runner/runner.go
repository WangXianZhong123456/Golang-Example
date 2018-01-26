// 监控程序
//主要完成3个功能
//1 程序顺序执行任务
//2 程序超时监控
//3 程序监听中断
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	tasks     []func(int)
	complete  chan error
	timeout   <-chan time.Time
	interrupt chan os.Signal
}

var ErrTimeOut = errors.New("执行任务超时")
var ErrInterrupt = errors.New("执行者被中断")

//创建线程池
func New(tm time.Duration) *Runner {
	return &Runner{
		//complete:make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}

func (r *Runner) Add(task ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}
