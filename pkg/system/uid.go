package system

import (
	"bytes"
	"net"

	"github.com/brownhash/golog"
)

func GetUid() (uint64, error) {
    interfaces, err := net.Interfaces()

    if err != nil {
		golog.Debug(err)
        return uint64(0), err
    }

    for _, i := range interfaces {
        if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {

            // Skip locally administered addresses
            if i.HardwareAddr[0]&2 == 2 {
                continue
            }

            var mac uint64
            for index, addr := range i.HardwareAddr {
                if index >= 8 {
                    break
                }
                mac <<= 8
                mac += uint64(addr)
            }

            return mac, nil
        }
    }

    return uint64(0), nil
}