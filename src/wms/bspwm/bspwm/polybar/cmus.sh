#!/bin/bash

prepend_zero () {
        seq -f "%02g" $1 $1
}

artist=$(echo -n $(cmus-remote -C 'format_print %F'))

if [[ $artist = *[!\ ]* ]]; then
#        song=$(echo -n $(cmus-remote -C status | grep title | cut -c 11-))
        position=$(cmus-remote -C status | grep position | cut -c 10-)
        minutes1=$(prepend_zero $(($position / 60)))
        seconds1=$(prepend_zero $(($position % 60)))
        duration=$(cmus-remote -C status | grep duration | cut -c 10-)
        minutes2=$(prepend_zero $(($duration / 60)))
        seconds2=$(prepend_zero $(($duration % 60)))
        echo -n "$artist"
else
        echo
fi

