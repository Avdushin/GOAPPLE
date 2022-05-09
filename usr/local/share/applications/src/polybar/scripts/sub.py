#!/usr/bin/env python

import urllib.request
import json

USERNAME="BRODOBRO" #replace with channel username

data = urllib.request.urlopen("https://www.googleapis.com/youtube/v3/channels?part=statistics&forUsername="+USERNAME+"&key=AIzaSyALXS-GcI0-JZNIi1mxMxJ79zGa5S0vuQc").read()
subs = json.loads(data)["items"][0]["statistics"]["subscriberCount"]

print("{:,d}".format(int(subs)))