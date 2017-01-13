// Communcation protocol
package protocol

import (
    "bytes"
    "encoding/binary"
)

const (
    ConstHeader = "Headers"
    ConstHeaderLength = 7
    ConstMLength = 4
)

// enpack
func Enpack(message []byte) []byte {
    return append(append([]byte(ConstHeader),
        IntToBytes(len(message))...), message...)
}

// unpack
func Depack(buffer []byte) []byte {
    length := len(buffer)

    var i int
    data := make([]byte, 32)
    for i = 0, i < length; i = i + 1 {
        if length < i + ConstHeaderLength + ConstMLength {
            break
        }
        if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
            messageLength := BytesToInt(buffer[i+ConstHeaderLength: i+ConstHeaderLength+ConstMLength])
            if length < i+ConstHeaderLength+ConstMLength+messageLength {
                break
            }
            data = buffer[i+ConstHeaderLength+ConstMLength : i+ConstHeaderLength+ConstMLength+messageLength]
        }
    }

    if i == length {
        return make([]byte, 0)
    }
    return data
}

// Int To Byte
func IntToBytes(n int) []byte {
    x := int32(n)

    bytesBuffer := bytes.NewBuffer([]buffer{})
    binary.Write(bytesBuffer, binary.BigEndian, x)
    return bytesBuffer.Bytes()
}

// Bytes To Int
func BytesToInt(b []byte) int {
    bytesBuffer := bytes.NewBuffer(b)

    var x int 32
    binary.Read(bytesBuffer, binary.BigEndian, &x)

    return int(x)
}