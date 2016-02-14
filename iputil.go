package iputil

import "net"

type IP struct {
	net.IP
}

type IPNet struct {
	*net.IPNet
}

func Parse(s string) (IP, *IPNet, error) {
	ip := net.ParseIP(s)
	if ip != nil {
		v4 := ip.To4()
		if v4 != nil {
			ip = v4
		}
		return IP{ip}, nil, nil
	}
	ip, rawnet, err := net.ParseCIDR(s)
	net := IPNet{rawnet}
	return IP{ip}, &net, err
}

func (ip IP) Next() IP {
	return ip.rel(true)
}

func (ip IP) Prev() IP {
	return ip.rel(false)
}

func (ip IP) rel(b bool) IP {
	next := IP{ip.IP}
	for i := len(next.IP) - 1; i > -1; i-- {
		if b {
			next.IP[i]++
		} else {
			next.IP[i]--
		}
		if next.IP[i] != 0 {
			break
		}
	}
	return next
}

func (ipn *IPNet) NetworkIP() IP {
	return IP{ipn.IPNet.IP}
}

func (ipn *IPNet) BroadcastIP() IP {
	bcast := ipn.IP
	for i, v := range ipn.IPNet.Mask {
		bcast[i] |= (v ^ 255)
	}
	return IP{bcast}
}
