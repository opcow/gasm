;**********************
;*** file mult.com ***
;**********************
    PROCESSOR 6502

    .ORG $8000 ; seg 00 | 80 bytes

    CLD
    LDX #$06
    LDA #$00
sub_8005 
    STA $80,X
    DEX
    BPL sub_8005 
    LDA #$fd
    STA $80
    LDA #$64
    STA $81
    JSR sub_8016 
    RTS
sub_8016 
    LDA #$00
    STA $84
    STA $85
    BIT $80
    BPL sub_8022 
    LDA #$ff
sub_8022 
    STA $86
    LDX #$07
sub_8026 
    LSR $80
    BCC sub_8037 
    LDA $84
    CLC
    ADC $81
    STA $84
    LDA $85
    ADC $86
    STA $85
sub_8037 
    ASL $81
    ROL $86
    DEX
    BNE sub_8026 
    LDA $85
    SBC #$01
    STA $85
    RTS
    .byte $12, $34, $56, $78, $78, $56, $34, $12
    .byte $ff, $ee, $ab

    .ORG $8200 ; seg 01 | 35 bytes

    LDA #$00
    STA $84
    STA $85
    STA $86
    LDX #$08
sub_820a 
    LSR $80
    BCC sub_821b 
    LDA $84
    CLC
    ADC $81
    STA $84
    LDA $85
    ADC $86
    STA $85
sub_821b 
    ASL $81
    ROL $86
    DEX
    BNE sub_820a 
    RTS

    .ORG $8300 ; seg 02 | 8 bytes

    LDA #$00
    STA $84
    STA $85
    STA $86
