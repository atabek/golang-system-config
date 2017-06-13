# This file describes the network interfaces available on your system
# and how to activate them. For more information, see interfaces(5).

source /etc/network/interfaces.d/*

# The loopback network interface
auto lo
iface lo inet loopback

# The primary network interface
allow-hotplug {{.Netinfo.Eth0}}
#iface {{.Netinfo.Eth0}} inet dhcp

iface     {{.Netinfo.Eth0}} inet static
address   {{.Netinfo.IpAddr}}
network   {{.Netinfo.Network}}
netmask   {{.Netinfo.Netmask}}
broadcast {{.Netinfo.Broadcast}}
gateway   {{.Netinfo.Gateway}}
dns-nameservers {{.Netinfo.Dns}}
