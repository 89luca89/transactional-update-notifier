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

``` bash
transactional-update-notifier daemon
```

Or using `systemctl`:

``` bash
systemctl --user enable --now transactional-update-notifier
```

**Note:** After installing **Transactional Update Notifier** using `make`, the
`96-transactional-update-notifier.preset` preset file should enable the unit
service by default on next boot and all you needed to do is to start it with:

``` bash
systemctl --user start transactional-update-notifier
```

### Client

**Transactional Update Notifier** can be run from anywhere and executed by anyone,
it will find the notifier socket and notify the user.

``` bash
transactional-update-notifier client
```
