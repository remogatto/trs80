// Copyright 2012 Lawrence Kesteloot

package main

import (
	"time"
)

const timerHz = 30

func getTimerCh() <-chan time.Time {
	return time.Tick(time.Second / timerHz)
}

func (cpu *cpu) timerInterrupt(state bool) {
	if state {
		cpu.irqLatch |= timerIrqMask
	} else {
		cpu.irqLatch &^= timerIrqMask
	}
}

func (vm *vm) handleTimer() {
	vm.cpu.timerInterrupt(true)
	vm.cpu.diskMotorOffInterrupt(vm.checkDiskMotorOff())
}
