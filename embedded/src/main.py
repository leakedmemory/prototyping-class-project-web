import machine
import network
import secrets
import utime
import uping


ROUTER_IP = "192.168.1.1"

PING_COUNT = 8
PING_TIMEOUT = 5000
PING_INTERVAL = 10
PING_SIZE = 64


def wlan_connect(success_led: machine.Pin, fail_led: machine.Pin):
    success_led.value(1)
    fail_led.value(1)
    utime.sleep(1)

    wlan = network.WLAN(network.STA_IF)
    success_led.value(1)
    fail_led.value(0)
    utime.sleep(1)

    wlan.active(True)
    success_led.value(0)
    fail_led.value(1)
    utime.sleep(1)

    wlan.connect(secrets.ROUTER_SSID, secrets.ROUTER_PASSWORD)
    success_led.value(0)
    fail_led.value(0)
    utime.sleep(1)

    return wlan


def main() -> None:
    success_led = machine.Pin(28, machine.Pin.OUT)
    fail_led = machine.Pin(0, machine.Pin.OUT)

    success_led.value(0)
    fail_led.value(0)
    utime.sleep(2)

    wlan = wlan_connect(success_led, fail_led)
    utime.sleep(3)

    while not wlan.isconnected():
        success_led.value(0)
        fail_led.value(1)
        utime.sleep(3)
        wlan = wlan_connect(success_led, fail_led)

    success_led.value(1)
    fail_led.value(0)
    utime.sleep(3)

    success_led.value(0)
    fail_led.value(0)
    utime.sleep(2)

    while True:
        p = uping.ping(
            ROUTER_IP,
            count=PING_COUNT,
            timeout=PING_TIMEOUT,
            interval=PING_INTERVAL,
            quiet=True,
            size=PING_SIZE
        )

        # rx >= tx
        if p[1] >= p[0] * 0.75:
            success_led.value(1)
            fail_led.value(0)
        else:
            success_led.value(0)
            fail_led.value(1)

        utime.sleep(1)

        success_led.value(0)
        fail_led.value(0)
        utime.sleep(0.25)


if __name__ == "__main__":
    main()
