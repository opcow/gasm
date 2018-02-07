    PROCESSOR 6502
    LIST ON

OP1 = $02
OP2 = $4
SOP1 = $-3
SOP2 = $64
; LOCATIONS FOR 0 PAGE VARS
VARS = $80
M1 = $80
M2 = $81
S1 = $82
S2 = $83
RL = $84 ; RESULT LOW
RH = $85 ; RESULT HIGH
TMP = $86

    .ORG $8000
START
    CLD                     ; CLEAR DECIMAL MODE
    LDX #TMP-M1             ; SETUP LOOP
    LDA #0
CLEAR
    STA VARS,X              ; ZERO VARS
    DEX
    BPL CLEAR
    LDA #SOP1
    STA M1
    LDA #SOP2
    STA M2
    JSR MULT8
    RTS
MULT8 SUBROUTINE mult8
    LDA #0
    STA RL
    STA RH
    BIT M1
    BPL .POS
    LDA #$FF
.POS
    STA TMP                 ; FOR HIGH BYTE OF M2 POWS OF 2
    LDX #7                  ; 7 BIT SIGNED
.LOOP
    LSR M1                  ; SHIFT THIS OP EACH LOOP
    BCC .SKIP               ; IF BIT 0 WAS O DON'T ADD A POW OF 2
    LDA RL
    CLC
    ADC M2                  ; ADD POW OF 2 LOW TO RESULT
    STA RL                  ; STORE IT IN RES LOW BYTE
    LDA RH
    ADC TMP                 ; ADD POW OF 2 HIGH
    STA RH                  ; STORE
.SKIP
    ASL M2                  ; SHIFT FOR NEXT POWER OF 2
    ROL TMP                 ; SAME FOR HIGH BYTE
    DEX
    BNE .LOOP
    LDA RH                  ; EXTEND SIGN BIT
    SBC #1
    STA RH
.EXIT
    RTS
    DC   $12, $34, $56, $78
    DC.L $12345678
    DC   $ff, $ee, $ab

    .ORG $8200
UMULT8 SUBROUTINE umult8
    LDA #0
    STA RL
    STA RH
    STA TMP
    LDX #8                  ; 8 BIT UNSIGNED
.LOOP
    LSR M1                  ; SHIFT THIS OP EACH LOOP
    BCC .SKIP               ; IF BIT 0 WAS O DON'T ADD A POW OF 2
    LDA RL
    CLC
    ADC M2                  ; ADD POW OF 2 LOW TO RESULT
    STA RL                  ; STORE IT IN RES LOW BYTE
    LDA RH
    ADC TMP                 ; ADD POW OF 2 HIGH
    STA RH                  ; STORE
.SKIP
    ASL M2                  ; SHIFT FOR NEXT POWER OF 2
    ROL TMP                 ; SAME FOR HIGH BYTE
    DEX
    BNE .LOOP
    RTS

    .ORG $8300
    LDA #0
    STA RL
    STA RH
    STA TMP
