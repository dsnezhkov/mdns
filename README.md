# Purpose
mdns is a scanner to query services that respond to mDNS, one or many (over 12K). 
List of services is based on https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml

mDNS or Multicast DNS can be used to discover services on the local network without the use of an authoritative DNS server. This enables peer-to-peer discovery. It is important to note that many networks restrict the use of multicasting, which prevents mDNS from functioning. Notably, multicast cannot be used in any sort of cloud, or shared infrastructure environment. However it works well in most office, home, or private infrastructure environments.

It found its use in semi-passive fingerpirnting of local networks without reliance on more aggressive scanning techniques (port scanning).
IT sends one multicast packet to specified service.

# Install

### Compile:
`go build`  or download release (OSX)

TODO: X-platform build doc

# Usage

## Single Service Query

```bash
$ ./mdns  -service "_ssh._tcp"  -delay 100ms
[i] Looking up service :  _ssh._tcp
[i] Ctrl-C when done listening for responses

[+] Response: &{tibet._ssh._tcp.local. tibet.local. 192.168.88.8 <nil> 22  [] 10 192.168.88.8 true true}
^C
All done
```

## Multi-Service Query

```bash
$ cat ./mdns_service.some
_echo._udp:7
_chargen._tcp:19
_ipp._tcp:631
_telnet._tcp:23
_deos._tcp:76
_http._tcp:80
_bpmd._udp:3593
_ssh._tcp:22
```


### Sync
```bash
$ time ./mdns  -servicepath  ./mdns_service.some   -delay 1000ms
[i] Looking up service _echo._udp (port:7)
[i] Looking up service _chargen._tcp (port:19)
[i] Looking up service _ipp._tcp (port:631)
[+] Response: &{Brother\ DCP-8080DN._ipp._tcp.local. BRN001BA97BDF18.local. 192.168.88.29 <nil> 631 txtvers=1|qtotal=1|pdl=application/vnd.hp-PCL|rp=duerqxesz5090|ty=Brother DCP-8080DN|product=(Brother DCP-8080DN)|adminurl=http://BRN001BA97BDF18.local./|priority=50|usb_MFG=Brother|usb_MDL=DCP-8080DN|Color=F|Copies=T|Duplex=F|PaperCustom=T|Binary=T|Transparent=T|TBCP=F [txtvers=1 qtotal=1 pdl=application/vnd.hp-PCL rp=duerqxesz5090 ty=Brother DCP-8080DN product=(Brother DCP-8080DN) adminurl=http://BRN001BA97BDF18.local./ priority=50 usb_MFG=Brother usb_MDL=DCP-8080DN Color=F Copies=T Duplex=F PaperCustom=T Binary=T Transparent=T TBCP=F] 10 192.168.88.29 true true}
[i] Looking up service _telnet._tcp (port:23)
[i] Looking up service _deos._tcp (port:76)
[i] Looking up service _http._tcp (port:80)
[+] Response: &{Brother\ DCP-8080DN._http._tcp.local. BRN001BA97BDF18.local. 192.168.88.29 <nil> 80  [] 10 192.168.88.29 true true}
[i] Looking up service _bpmd._udp (port:3593)
[+] Response: &{media-m._afpovertcp._tcp.local. media-m.local. <nil> fe80::8de:cf1c:5ec1:c59b 548  [] 4500 fe80::8de:cf1c:5ec1:c59b true true}
[i] Looking up service _ssh._tcp (port:22)
[+] Response: &{tibet._ssh._tcp.local. tibet.local. 192.168.88.8 <nil> 22  [] 10 192.168.88.8 true true}
[+] Response: &{media-m._ssh._tcp.local. media-m.local. <nil> fe80::8de:cf1c:5ec1:c59b 22  [] 10 fe80::8de:cf1c:5ec1:c59b true true}
All done

real	0m16.044s
user	0m0.007s
sys	0m0.014s
```

### Async

```bash
$ time ./mdns  -servicepath  ./mdns_service.some  -async  -delay 1000ms
[i] Looking up service _echo._udp (port:7)
[i] Looking up service _chargen._tcp (port:19)
[i] Looking up service _ipp._tcp (port:631)
[i] Looking up service _telnet._tcp (port:23)
[+] Response: &{Brother\ DCP-8080DN._ipp._tcp.local. BRN001BA97BDF18.local. 192.168.88.29 <nil> 631 txtvers=1|qtotal=1|pdl=application/vnd.hp-PCL|rp=duerqxesz5090|ty=Brother DCP-8080DN|product=(Brother DCP-8080DN)|adminurl=http://BRN001BA97BDF18.local./|priority=50|usb_MFG=Brother|usb_MDL=DCP-8080DN|Color=F|Copies=T|Duplex=F|PaperCustom=T|Binary=T|Transparent=T|TBCP=F [txtvers=1 qtotal=1 pdl=application/vnd.hp-PCL rp=duerqxesz5090 ty=Brother DCP-8080DN product=(Brother DCP-8080DN) adminurl=http://BRN001BA97BDF18.local./ priority=50 usb_MFG=Brother usb_MDL=DCP-8080DN Color=F Copies=T Duplex=F PaperCustom=T Binary=T Transparent=T TBCP=F] 10 192.168.88.29 true true}
[i] Looking up service _deos._tcp (port:76)
[i] Looking up service _http._tcp (port:80)
[i] Looking up service _bpmd._udp (port:3593)
[+] Response: &{Brother\ DCP-8080DN._http._tcp.local. BRN001BA97BDF18.local. 192.168.88.29 <nil> 80  [] 10 192.168.88.29 true true}
[i] Looking up service _ssh._tcp (port:22)
All done

real	0m8.022s
user	0m0.006s
sys	0m0.011s
```

### ACIICAST:
https://asciinema.org/a/JF2HXIYVfsPxaRCI8jbejuiVM

