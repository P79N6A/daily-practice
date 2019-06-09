#!/usr/bin/env bash

nasm -f bin c13_1_mbr.asm -o c13_1_mbr.bin
nasm -f bin c13_2_core.asm -o c13_2_core.bin
nasm -f bin c13_3_user.asm -o c13_3_user.bin

# dd if=c13_1_mbr.bin of=~/VirtualBox\ VMs/LearnASM/LearnASMNew.vhd conv=notrunc
# dd if=c13_2_core.bin of=~/VirtualBox\ VMs/LearnASM/LearnASMNew.vhd seek=1  conv=notrunc
# dd if=c13_3_user.bin of=~/VirtualBox\ VMs/LearnASM/LearnASMNew.vhd seek=50  conv=notrunc
# dd if=c13_diskdata.txt of=~/VirtualBox\ VMs/LearnASM/LearnASMNew.vhd seek=100  conv=notrunc

python fixvhdwr.py c13_1_mbr.bin ~/VirtualBox\ VMs/LearnASM/LearnASMNew.vhd 0
python fixvhdwr.py c13_2_core.bin ~/VirtualBox\ VMs/LearnASM/LearnASMNew.vhd 1
python fixvhdwr.py c13_3_user.bin ~/VirtualBox\ VMs/LearnASM/LearnASMNew.vhd 50
python fixvhdwr.py c13_diskdata.txt ~/VirtualBox\ VMs/LearnASM/LearnASMNew.vhd 100
