from ppadb.client import Client as AdbClient
import time
import schedule

# https://blog.testproject.io/2021/08/10/useful-adb-commands-for-android-testing/
LAVATORY_X_POX = 508
LAVATORY_Y_POS = 1546

CONTRIBUTE_X_POS = 540
CONTRIBUTE_Y_POS = 1880

RSS_SWIPE_START_X = 406
RSS_SWIPE_START_Y = 996

RSS_SWIPE_END_X = 430
RSS_SWIPE_END_Y = 989

device = None


def connect():
    _client = AdbClient(host="127.0.0.1", port=5037)

    devices = _client.devices()

    if len(devices) == 0:
        print('No devices')
        quit()

    _device = devices[0]

    print(f'Connected to {_device}')

    return _device, _client


def send_rss():
    for i in range(2):
        # Tap the lavatory
        device.shell(f'input tap {LAVATORY_X_POX} {LAVATORY_Y_POS}')
        time.sleep(1)
        # Tap contribute
        device.shell(f'input tap {CONTRIBUTE_X_POS} {CONTRIBUTE_Y_POS}')
        time.sleep(1)
        # swipe to max rss
        device.shell(f'input swipe {RSS_SWIPE_START_X} {RSS_SWIPE_START_Y} {RSS_SWIPE_END_X} {RSS_SWIPE_END_Y} 700')
        time.sleep(1)
        # tap send, same location as contribute
        device.shell(f'input tap {CONTRIBUTE_X_POS} {CONTRIBUTE_Y_POS}')
        time.sleep(2)


if __name__ == '__main__':
    device, _ = connect()

    # 4 seconds, to give marches time to come back
    schedule.every(4).seconds.do(send_rss)
    while True:
        schedule.run_pending()
