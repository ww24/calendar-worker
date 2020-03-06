#!/bin/ash

set -eo pipefail

if [ ! -f "$GOOGLE_APPLICATION_CREDENTIALS" ] && [ -n "$SERVICE_ACCOUNT" ]; then
    echo -n "$SERVICE_ACCOUNT" | base64 -d > "$GOOGLE_APPLICATION_CREDENTIALS"
fi

calendar-notifier -config /usr/local/etc/calendar-notifier/config.yml
