package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

type Instruction struct {
	operator uint16
	operand  uint16
}

const (
	op_inc_dp = iota
	op_dec_dp
	op_inc_val
	op_dec_val
	op_out
	op_in
	op_jmp_fwd
	op_jmp_bck
)

const data_size int = 65535

func compile_bf(input string) (program []Instruction, err error) {
	var pc, jmp_pc uint16 = 0, 0
	jmp_stack := make([]uint16, 0)
	for _, c := range input {
		switch c {
		case '>':
			program = append(program, Instruction{op_inc_dp, 0})
		case '<':
			program = append(program, Instruction{op_dec_dp, 0})
		case '+':
			program = append(program, Instruction{op_inc_val, 0})
		case '-':
			program = append(program, Instruction{op_dec_val, 0})
		case '.':
			program = append(program, Instruction{op_out, 0})
		case ',':
			program = append(program, Instruction{op_in, 0})
		case '[':
			program = append(program, Instruction{op_jmp_fwd, 0})
			jmp_stack = append(jmp_stack, pc)
		case ']':
			if len(jmp_stack) == 0 {
				return nil, errors.New("Compilation error.")
			}
			jmp_pc = jmp_stack[len(jmp_stack)-1]
			jmp_stack = jmp_stack[:len(jmp_stack)-1]
			program = append(program, Instruction{op_jmp_bck, jmp_pc})
			program[jmp_pc].operand = pc
		default:
			pc--
		}
		pc++
	}
	if len(jmp_stack) != 0 {
		return nil, errors.New("Compilation error.")
	}
	return
}

func execute_bf(program []Instruction) string {
	data := make([]int16, data_size)
	var buffer bytes.Buffer
	var data_ptr uint16 = 0
	reader := bufio.NewReader(os.Stdin)
	for pc := 0; pc < len(program); pc++ {
		switch program[pc].operator {
		case op_inc_dp:
			data_ptr++
		case op_dec_dp:
			data_ptr--
		case op_inc_val:
			data[data_ptr]++
		case op_dec_val:
			data[data_ptr]--
		case op_out:
			buffer.WriteString(fmt.Sprintf("%c", data[data_ptr]))
		case op_in:
			read_val, _ := reader.ReadByte()
			data[data_ptr] = int16(read_val)
		case op_jmp_fwd:
			if data[data_ptr] == 0 {
				pc = int(program[pc].operand)
			}
		case op_jmp_bck:
			if data[data_ptr] > 0 {
				pc = int(program[pc].operand)
			}
		default:
			panic("Unknown operator.")
		}
	}
	log.Println("Result: ", buffer.String())
	return buffer.String()
}

func Execute(input string) (string, error) {
	program, err := compile_bf(input)

	if err != nil {
		return "", nil
	}

	return execute_bf(program), nil
}
