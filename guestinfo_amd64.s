// Copyright 2016 Yann Hodique
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
