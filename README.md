# lastpost [![Build Status](https://travis-ci.org/justincampbell/lastpost.svg?branch=master)](https://travis-ci.org/justincampbell/lastpost)

> Fetches the latest item from an RSS feed, and attempts to display it nicely.

## Installation

1. Download the latest package for your platform from the [Releases page](https://github.com/justincampbell/lastpost/releases/latest).
2. Untar the package with `tar -zxvf lastpost*.tar.gz`.
3. Move the extracted `lastpost` file to a directory in your `$PATH` (for most systems, this will be `/usr/local/bin/`).

Or, if you have a [Go development environment](https://golang.org/doc/install):

```
go get github.com/justincampbell/lastpost
```

## Examples

### Get the workout of the day from your gym's blog

```
$ lastpost http://crossfitwc.com/category/wod/feed/
WOD 4/29/16 (18 hours ago)

15 minutes to a heavy clean
Then...
WWPF
For Time:
100 Cal Row, (pet rock)
75 C2B Pullups, Partner hangs from bar
50 Handstand Pushups, partner holds handstand
25 Ground To OH, 185/133
Sleep and hydration are essential for recovery
```

Any `lastpost` command could also be set as a shell alias (in `~/.profile` or `~/.bash_profile`):

```sh
wod() { lastpost http://crossfitwc.com/category/wod/feed/; }
```

### Check the status of a web service

```
$ lastpost http://status.aws.amazon.com/rss/ec2-us-east-1.rss
Service is operating normally: [RESOLVED] Network Connectivity

Between 1:59 PM and 2:01 PM PDT, some EC2 instances experienced network connectivity issues between two Availability Zones in the US-EAST-1 Region. This issue has been resolved and the service is operating normally.
```

```
$ lastpost https://status.github.com/messages.rss
[good] Everything operating normally. (5 days ago)

Everything operating normally.
```

```
$ lastpost https://status.heroku.com/feed
Resolved: Emergency EU platform maintenance scheduled on May 3 at 17:45 UTC

This maintenance is complete.
```

### Display a weather report

```
$ lastpost http://w1.weather.gov/xml/current_obs/KPHL.rss
Overcast and 52 F at Philadelphia, Philadelphia International Airport, PA

Winds are East at 10.4 MPH (9 KT). The pressure is 1007.9 mb and the humidity is 83%.
The wind chill is 48. Last Updated on May 4 2016, 8:54 am EDT.
```
