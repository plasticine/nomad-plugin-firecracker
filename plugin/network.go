package plugin

import (
    "crypto/rand"
    "fmt"
    "net"
    "os"

    "github.com/vishvananda/netlink"
)

// NetworkInterface defines a network interface.
type NetworkInterface struct {
    Name     string
    HardAddr string
    Addrs    []netlink.Addr
}

// TapIface defines a TAP interface
type TapIface struct {
    ID    string
    Name  string
    Iface NetworkInterface
    VMFds []*os.File
}

// NetworkInterfacePair defines a pair between VM and virtual network interfaces.
type NetworkInterfacePair struct {
    TapIface  TapIface
    VirtIface NetworkInterface
}

func createTapNetwork(idx int64, allocID string) (netIfacePair *NetworkInterfacePair, err error) {
    netHandle, err := netlink.NewHandle()
    if err != nil {
        return &NetworkInterfacePair{}, err
    }
    defer netHandle.Delete()

    randomMacAddr, err := generateRandomPrivateMacAddr()
    if err != nil {
        return &NetworkInterfacePair{}, fmt.Errorf("Could not generate random mac address: %s", err)
    }

    netIfacePair = &NetworkInterfacePair{
        TapIface: TapIface{
            ID:   allocID,
            Name: fmt.Sprintf("eth%d", idx),
            Iface: NetworkInterface{
                Name: fmt.Sprintf("tap%d_nomad", idx),
            },
        },
        VirtIface: NetworkInterface{
            Name:     fmt.Sprintf("eth%d", idx),
            HardAddr: randomMacAddr,
        },
    }

    fmt.Printf("netIfacePair = %#v \n", netIfacePair)

    tapLink, fds, err := createLink(netHandle, netIfacePair.TapIface.Iface.Name)
    if err != nil {
        return &NetworkInterfacePair{}, fmt.Errorf("Could not create TAP interface: %s", err)
    }

    netIfacePair.TapIface.VMFds = fds

    // Save the MAC address to the TAP so that it can later be used
    // to build the QMP command line. This MAC address has to be
    // the one inside the VM in order to avoid any firewall issues. The
    // bridge created by the network plugin on the host actually expects
    // to see traffic from this MAC address and not another one.
    netIfacePair.TapIface.Iface.HardAddr = tapLink.Attrs().HardwareAddr.String()

    if err := netHandle.LinkSetMTU(tapLink, tapLink.Attrs().MTU); err != nil {
        return &NetworkInterfacePair{}, fmt.Errorf("Could not set TAP MTU %d: %s", tapLink.Attrs().MTU, err)
    }

    if err := netHandle.LinkSetUp(tapLink); err != nil {
        return &NetworkInterfacePair{}, fmt.Errorf("Could not enable TAP %s: %s", netIfacePair.TapIface.Iface.Name, err)
    }

    return netIfacePair, nil
}

func createLink(netHandle *netlink.Handle, name string) (newLink netlink.Link, fds []*os.File, err error) {
    newLink = &netlink.Tuntap{
        LinkAttrs: netlink.LinkAttrs{Name: name},
        Mode:      netlink.TUNTAP_MODE_TAP,
        Flags:     netlink.TUNTAP_VNET_HDR,
    }

    if err := netHandle.LinkAdd(newLink); err != nil {
        return nil, fds, fmt.Errorf("LinkAdd() failed for %s name %s: %s", newLink.Type(), name, err)
    }

    tuntapLink, ok := newLink.(*netlink.Tuntap)
    if ok {
        fds = tuntapLink.Fds
    }

    newLink, err = netHandle.LinkByName(name)
    if err != nil {
        return nil, fds, fmt.Errorf("LinkByName() failed for name %s: %s", name, err)
    }

    return newLink, fds, err
}

func generateRandomPrivateMacAddr() (string, error) {
    buf := make([]byte, 6)
    _, err := rand.Read(buf)
    if err != nil {
        return "", err
    }

    // Set the local bit for local addresses
    // Addresses in this range are local mac addresses:
    // x2-xx-xx-xx-xx-xx , x6-xx-xx-xx-xx-xx , xA-xx-xx-xx-xx-xx , xE-xx-xx-xx-xx-xx
    buf[0] = (buf[0] | 2) & 0xfe

    hardAddr := net.HardwareAddr(buf)
    return hardAddr.String(), nil
}
