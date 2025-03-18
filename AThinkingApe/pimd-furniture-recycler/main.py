from ppadb.client import Client as AdbClient
import time

# See README.md for more info about these constants.

RECYCLE_POS = "221 1071"
HOLD_RECYCLE_POS = "544 1715"


def connect():
    client = AdbClient()
    devices = client.devices()

    if len(devices) == 0:
        print('No devices')
        quit()

    device = devices[0]

    print(f'Connected to the first device: {device}')

    return device, client


def start_recycle(recycle_num=1):
    device, client = connect()

    for i in range(recycle_num):
        device.shell(f"input tap {RECYCLE_POS}")  # recycle button
        time.sleep(1)  # wait for recycle screen to load
        device.shell(
            f"input touchscreen swipe {HOLD_RECYCLE_POS} {HOLD_RECYCLE_POS} 5000"
        )  # hold down recycle button
        print(f"Recycled {i + 1}/{recycle_num}")
        time.sleep(1)  # wait to get back to recycle screen


if __name__ == '__main__':
    num_to_recycle = input("Amount of furniture's to recycle: ")
    start_recycle(int(num_to_recycle))
