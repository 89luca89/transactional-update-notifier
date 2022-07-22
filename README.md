# Transactional Update Notifier

This tool can be used in conjunction with `transactional-update`'s notify method in
order to notify all graphically logged in users about updates performed.

# Build

`go build`

# Usage

## Daemon

this should run as an user's systemd unit. 

`transactional-update-notifier daemon`

## Client

this can run from anywhere and executed by anyone, this will find all the notifier
socket and notify them.

`transactional-update-notifier client`


