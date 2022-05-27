#/bin/env sh
export CO_MIGRATIONS=file:///migrations
/confetti migrate
/confetti serve
