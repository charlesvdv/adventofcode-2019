package main

import (
	"errors"
	"log"
)

const (
	PARAM_MODE_POSITION  = 0
	PARAM_MODE_IMMEDIATE = 1
)

type Input = int
type Output = []int

type Processor struct {
	initialProgram []int
	program        []int
	pc             int
	out            Output
	paramsMode     parametersMode
}

func NewProcessor(program []int) Processor {
	return Processor{
		initialProgram: append([]int{}, program...),
		program:        append([]int{}, program...),
		pc:             0,
		out:            []int{},
		paramsMode:     emptyParametersMode(),
	}
}

func (proc *Processor) Run(input Input) (Output, error) {
	for true {
		instruction := proc.program[proc.pc]
		opcode, paramsMode := parseInstruction(instruction)
		// fmt.Printf("instruction: %v, opcode: %v, paramsMode: %v\n", instruction, opcode, paramsMode.values)
		proc.paramsMode = paramsMode
		// programEndIndex := int(math.Min(float64(proc.pc+5), float64(len(proc.program))))
		// fmt.Printf("pc: %v, next data: %v, params mode: %v, output: %v\n", proc.pc, proc.program[proc.pc:programEndIndex], proc.paramsMode.values, proc.out)
		proc.pc += 1

		// fmt.Printf("PROGRAM COUNTER: %v\n", proc.pc)
		switch opcode {
		case 1:
			proc.add()
		case 2:
			proc.multiply()
		case 3:
			proc.input(input)
		case 4:
			proc.output()
		case 5:
			proc.jumpIfTrue()
		case 6:
			proc.jumpIfFalse()
		case 7:
			proc.lessThan()
		case 8:
			proc.equals()
		case 99:
			output := proc.out
			proc.restore()
			return output, nil
		default:
			return Output{}, errors.New("unknown opcode")
		}
	}

	return proc.out, errors.New("unexpected finish")
}

func (proc *Processor) restore() {
	proc.program = append([]int{}, proc.initialProgram...)
	proc.pc = 0
	proc.out = []int{}
	proc.paramsMode = emptyParametersMode()
}

func (proc *Processor) add() {
	outputAddress := proc.program[proc.pc+2]
	// fmt.Printf("run add %v + %v to %v\n", proc.getInput(0), proc.getInput(1), outputAddress)
	proc.program[outputAddress] = proc.getInput(0) + proc.getInput(1)
	proc.pc += 3
}

func (proc *Processor) multiply() {
	outputAddress := proc.program[proc.pc+2]
	proc.program[outputAddress] = proc.getInput(0) * proc.getInput(1)
	proc.pc += 3
}

func (proc *Processor) input(input Input) {
	// fmt.Printf("pc: %v, address: %v\n", proc.pc, proc.program[proc.pc])
	address := proc.program[proc.pc]
	// fmt.Printf("run input: set value %v to address %v\n", input, address)
	proc.program[address] = input
	proc.pc += 1
}

func (proc *Processor) output() {
	proc.out = append(proc.out, proc.getInput(0))
	proc.pc += 1
}

func (proc *Processor) jumpIfTrue() {
	if proc.getInput(0) != 0 {
		proc.pc = proc.getInput(1)
	} else {
		proc.pc += 2
	}
}

func (proc *Processor) jumpIfFalse() {
	if proc.getInput(0) == 0 {
		proc.pc = proc.getInput(1)
	} else {
		proc.pc += 2
	}
}

func (proc *Processor) lessThan() {
	value := 0
	if proc.getInput(0) < proc.getInput(1) {
		value = 1
	}
	outAddress := proc.program[proc.pc+2]
	proc.program[outAddress] = value
	proc.pc += 3
}

func (proc *Processor) equals() {
	value := 0
	if proc.getInput(0) == proc.getInput(1) {
		value = 1
	}
	// fmt.Printf("equals: %v == %v outputs %v to address %v\n", proc.getInput(0), proc.getInput(1), value, proc.program[proc.pc+2])
	outAddress := proc.program[proc.pc+2]
	proc.program[outAddress] = value
	// fmt.Printf("after modification: %v\n", proc.program)
	proc.pc += 3
}

func (proc *Processor) getInput(index int) int {
	val := proc.program[proc.pc+index]
	switch proc.paramsMode.get(index) {
	case PARAM_MODE_IMMEDIATE:
		// fmt.Printf("input immediate: %v\n", val)
		return val
	case PARAM_MODE_POSITION:
		// fmt.Printf("input position: address: %v and value: %v\n", val, proc.program[val])
		return proc.program[val]
	}
	log.Fatal("unknown param mode")
	return 0
}

func parseInstruction(instruction int) (int, parametersMode) {
	digits := numberToDigits(instruction)
	if len(digits) <= 2 {
		return instruction, emptyParametersMode()
	}
	startOpCodeIndex := len(digits) - 2
	return digitsToNumber(digits[startOpCodeIndex:len(digits)]), newParametersMode(digits[0:(startOpCodeIndex)])
}

type parametersMode struct {
	values []int
}

func newParametersMode(modes []int) parametersMode {
	// reverse modes
	for i := len(modes)/2 - 1; i >= 0; i-- {
		opp := len(modes) - 1 - i
		modes[i], modes[opp] = modes[opp], modes[i]
	}
	return parametersMode{
		values: modes,
	}
}

func emptyParametersMode() parametersMode {
	return newParametersMode([]int{})
}

func (paramMode parametersMode) get(index int) int {
	if index < len(paramMode.values) {
		return paramMode.values[index]
	}
	return PARAM_MODE_POSITION
}
