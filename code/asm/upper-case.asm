assume cs:codesg

datasg segment
    db "Beginner's All-purpose Symbolic Instruction Code.", 0
datasg end

codesg segment
    begin:      mov ax, datasg
                mov ds, ax
                mov si, 0
                call letterc
                mov ax, 4c00h
                int 21h

    letterc:    mov cl, [si]
                mov ch, 0
                jcxz ok
                cmp byte ptr [si], 97
                jb letterc
                cmp byte ptr [si], 122
                ja letterc
                and byte ptr [si], 11011111b
                inc si
                jmp short letterc
    ok:         ret

codesg end
end begin