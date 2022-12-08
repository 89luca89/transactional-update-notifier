# Transactional Update Notifier

This tool can be used in conjunction with `transactional-update`'s notify method in
order to notify all graphically logged in users about updates performed.

## Build and installation

All you're going to need is [Go](https://go.dev/) `>= 1.18` and
[GNU Make](https://www.gnu.org/software/make/)

``` bash
make
sudo make install
```

## Usage

### Daemon

**Transactional Update Notifier** should be run as a user's Systemd unit. 

This will wait for messages over dbus at `org.opensuse.tukit.Updated` and trigger the graphical
notification when receiving the signal.

Graphical notifications are performed using user's dbus session.

``` console
~$: transactional-update-notifier daemon
```

Or using `systemctl`:

``` console
~$: systemctl --user enable --now transactional-update-notifier
```

**Note:** After installing **Transactional Update Notifier** using `make`, the
`96-transactional-update-notifier.preset` preset file should enable the unit
service by default on next boot and all you needed to do is to start it with:

``` bash
~$: systemctl --user start transactional-update-notifier
```

### Client

**Transactional Update Notifier** can be run from anywhere and executed by anyone,
it will send messages over dbus on `org.opensuse.tukit.Updated` and all listening services
will trigger a graphical notification.

``` console
~#: transactional-update-notifier client
```

Be aware that the file `org.opensuse.tukit.Updated.conf` should be put into `/etc/dbus-1/system.d/` in
order to protect the `org.opensuse.tukit.Updated` name and allow only `root` to emit signals on this
interface.
