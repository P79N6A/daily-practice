#!/usr/bin/env bash

nasm -f bin c13_1_mbr.asm -o c13_mbr.bin
nasm -f bin c13_2_core.asm -o c13_core.bin
nasm -f bin c13_3_user.asm -o c13_user.bin

python fixvhdwr.py c13_mbr.bin ~/VirtualBox\ VMs/LearnASM/LearnASM.vhd 0
python fixvhdwr.py c13_core.bin ~/VirtualBox\ VMs/LearnASM/LearnASM.vhd 1
python fixvhdwr.py c13_user.bin ~/VirtualBox\ VMs/LearnASM/LearnASM.vhd 50