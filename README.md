# slow

```
NAME
	slow, version ` + VERSION + ` - Non-interactive pager (stdout delay)
	Pipe fast things (such as lsmod or dmesg) into it.

AUTHOR
	Copyright (c) 2016 aerth [aerth@sdf.org]

SYNOPSIS
	[FILE] | slow -d 300

USAGE
	lsmod | slow 
	tree / | slow
	
TIPS
	Custom Delay: cat main.go | slow -d 100
	Delay is in microseconds. Default is 300. 
	The exception to '-d' flag is setting to 1-5 which are in seconds.
```