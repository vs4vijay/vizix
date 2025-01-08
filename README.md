# VIZIX

An operating system for fun and learning

## Development Flow (Write an OS)
- Create Boot Sector
- 

## Development Tools Required
- qemu
- nasm
- gcc 

PS: On Mac, Using homebrew: `brew install qemu nasm`

## Running
- `nasm -f bin Vizix.asm -o Vizix.bin`
- `qemu Vizix.bin`

## Bootsector Example:
```
; Infinite loop
loop:
    jmp loop 

; Fill with 510 zeros minus the size of the previous code
times 510-($-$$) db 0

; Magic number
dw 0xaa55
```

## Registers:
- General Purpose Registers:
  - `ax`; Accumulator (EAX)
  - `bx`; Base (EBX)
  - `cx`; Counter (ECX)
  - `dx`; Data (EDX)
- Stack Registers:
  - `bp`; Base Pointer, Stores the base address of stack (EBP)
  - `sp`; Stack Pointer, Stores the top address of stack, sp gets decremented (ESP)
- Pointer Registers:
  - `si`; Source Index, (ESI)
  - `di`; Destination Index, (EDI)
  - `ip`; Instruction Pointer
- Segmentation: Obsolete memory protection technique
  - Segment Register: Can be altered using registers
    - `cs`; Code Segment; Can not be altered
    - `ds`; Data Segment
    - `ss`; Stack Segment
    - `es`; Extra Segment
    - `fs`, `gs`; General Purpose Segment
- Special Registers:
  - `CR0,2,3`
  - `DR0,1,2,3,6,7`
  - `TR4,5,6,7`


## Bootloaders:
- Geek Loading: Everything in Boot Record
- One-stage Loading
- Two-stage Loading: `GRUB` (GRand Unified Bootloader)



### References:
- Little OS Book - https://littleosbook.github.io/
- https://wiki.osdev.org/Bare_Bones
- https://github.com/cfenollosa/os-tutorial
- https://www.cs.bham.ac.uk/~exr/lectures/opsys/10_11/lectures/os-dev.pdf
- https://www.cs.cmu.edu/~410/
- https://wiki.osdev.org/Main_Page
- http://stanislavs.org/helppc/idx_interrupt.html
- https://legacy.gitbook.com/book/samypesse/how-to-create-an-operating-system/details
- https://samypesse.gitbook.io/how-to-create-an-operating-system/
- https://github.com/SamyPesse/How-to-Make-a-Computer-Operating-System
- https://github.com/arjun024/mkernel
- http://3zanders.co.uk/2017/10/13/writing-a-bootloader/
- http://mikeos.sourceforge.net/write-your-own-os.html
- http://osblog.stephenmarz.com/index.html
- https://os.phil-opp.com/
- https://www.rpi4os.com/
- https://s-matyukevich.github.io/raspberry-pi-os/
- https://www.fysnet.net/osdesign_book_series.htm
- https://pages.cs.wisc.edu/~remzi/OSTEP/


### Hardware
- https://eater.net/8bit/
- https://nand.arhan.sh/
- https://www.nand2tetris.org/
- https://nandgame.com/
- https://www.megaprocessor.com/
- https://gitlab.com/x653/nand2tetris-fpga/
