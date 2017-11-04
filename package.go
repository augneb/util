package utils

import (
	"bytes"
	"context"
	"net"
	"strconv"
	"github.com/pkg/errors"
)

type ReadPackageCall func(pack []byte) bool

func ReadPackage(ctx context.Context, conn net.Conn, call ReadPackageCall, packEof ...interface{}) error {
	msgBuffer := make([]byte, 0)

	isEofModel := len(packEof) > 0

	eof := []byte{}
	switch packEof[0].(type) {
	case string:
		eof = []byte(packEof[0].(string))
	case []byte:
		eof = packEof[0].([]byte)
	}
	eofLen := len(eof)

	for {
		select {
		case <-ctx.Done():
			return errors.New("timeout")
		default:
		}

		buffer := make([]byte, 256)

		n, err := conn.Read(buffer)
		if err != nil {
			return err
		}

		buffer = buffer[:n]

		if len(msgBuffer) > 0 {
			msgBuffer = append(msgBuffer, buffer...)
		} else {
			msgBuffer = buffer
		}

		if isEofModel {
			for {
				idx := bytes.Index(msgBuffer, eof)
				if idx == -1 {
					break
				}

				doNext := call(msgBuffer[:idx])
				if !doNext {
					return nil
				}

				idx += eofLen
				msgBuffer = msgBuffer[idx:]
			}
		} else {
			for {
				// 数据不足
				if len(msgBuffer) < 4 {
					break
				}

				// 前4个字节是包长度信息
				packLen, err := strconv.Atoi(string(msgBuffer[:4]))
				if err != nil {
					return err
				}

				l := packLen + 4

				// 数据不足
				if len(msgBuffer) < l {
					break
				}

				doNext := call(msgBuffer[:l])
				if !doNext {
					return nil
				}

				msgBuffer = msgBuffer[l:]
			}
		}
	}

	return nil
}
