// func asmCommunicate(iax, ibx, icx, idx, idi, isi uint32) (oax, obx, ocx, odx, odi, osi uint32)
TEXT ·asmCommunicate(SB),NOSPLIT,$0-48
        MOVL iax+0(FP), AX
        MOVL ibx+4(FP), BX
        MOVL icx+8(FP), CX
        MOVL idx+12(FP), DX
        MOVL idi+16(FP), DI
        MOVL isi+20(FP), SI
        
        INL
        
        MOVL AX, oax+24(FP)
        MOVL BX, obx+28(FP)
        MOVL CX, ocx+32(FP)
        MOVL DX, odx+36(FP)
        MOVL DI, odi+40(FP)
        MOVL SI, osi+44(FP)
        
        RET
