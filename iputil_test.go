package iputil

import "testing"

func TestParse2(t *testing.T) {
	var (
		s   string
		ip  IP
		net *IPNet
		err error
	)
	s = "10.0.0.1"
	ip, net, err = Parse(s)
	if ip.String() != s {
		t.Errorf("got=%+v, expected=%+v", ip.String(), s)
	}
	if net != nil {
		t.Errorf("got=%+v, expected=nil", net)
	}
	if err != nil {
		t.Errorf("got=%+v, expected=nil", err)
	}

	s = "fe80::1"
	ip, net, err = Parse(s)
	if ip.String() != "fe80::1" {
		t.Errorf("got=%+v, expected=%+v", ip.String(), "fe80::1")
	}
	if net != nil {
		t.Errorf("got=%+v, expected=nil", net)
	}
	if err != nil {
		t.Errorf("got=%+v, expected=nil", err)
	}

	s = "10.0.0.0/24"
	ip, net, err = Parse(s)
	if ip.String() != "10.0.0.0" {
		t.Errorf("got=%+v, expected=%+v", ip.String(), "10.0.0.0")
	}
	if net.String() != s {
		t.Errorf("got=%+v, expected=%+v", net, s)
	}
	if err != nil {
		t.Errorf("got=%+v, expected=nil", err)
	}
}

func TestNext(t *testing.T) {
	var (
		got string
		exp string
		ip  IP
	)
	ip, _, _ = Parse("10.0.0.1")
	exp = "10.0.0.2"
	got = ip.Next().String()
	if got != exp {
		t.Errorf("got=%+v, expected=%+v", got, exp)
	}

	ip, _, _ = Parse("10.0.0.255")
	exp = "10.0.1.0"
	got = ip.Next().String()
	if got != exp {
		t.Errorf("got=%+v, expected=%+v", got, exp)
	}

	ip, _, _ = Parse("255.255.255.255")
	exp = "0.0.0.0"
	got = ip.Next().String()
	if got != exp {
		t.Errorf("got=%+v, expected=%+v", got, exp)
	}
}

func TestNetworkIP(t *testing.T) {
	var (
		got string
		exp string
		net *IPNet
	)
	_, net, _ = Parse("10.0.0.1/24")
	exp = "10.0.0.0"
	got = net.NetworkIP().String()
	if got != exp {
		t.Errorf("got=%+v, expected=%+v", got, exp)
	}

	_, net, _ = Parse("10.0.0.1/32")
	exp = "10.0.0.1"
	got = net.NetworkIP().String()
	if got != exp {
		t.Errorf("got=%+v, expected=%+v", got, exp)
	}

	_, net, _ = Parse("fe80::1/32")
	exp = "fe80::"
	got = net.NetworkIP().String()
	if got != exp {
		t.Errorf("got=%+v, expected=%+v", got, exp)
	}
}

func TestBroadcastIP(t *testing.T) {
	var (
		got string
		exp string
		net *IPNet
	)
	_, net, _ = Parse("10.0.0.1/24")
	exp = "10.0.0.255"
	got = net.BroadcastIP().String()
	if got != exp {
		t.Errorf("got=%+v, expected=%+v", got, exp)
	}

	_, net, _ = Parse("fe80::/64")
	exp = "fe80::ffff:ffff:ffff:ffff"
	got = net.BroadcastIP().String()
	if got != exp {
		t.Errorf("got=%+v, expected=%+v", got, exp)
	}
}
