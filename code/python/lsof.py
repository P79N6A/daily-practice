import socket
import errno
import sys

def port_occupied(port):
    is_occpuied = False
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    try:
        s.bind(("127.0.0.1", int(port)))
    except socket.error as e:
        if e.errno == errno.EADDRINUSE:
            is_occpuied = True
    s.close()
    return is_occpuied


if __name__ == '__main__':
    length = len(sys.argv)
    if length != 2:
        print('usage: python lsof.py PORT_NUMBER')
        sys.exit(-1)
    port = sys.argv[1]
    print('PORT: {0}, Occupied: {1}'.format(port, port_occupied(port)))