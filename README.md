# wasgubata
###### ./wgb knows geolocations...

[![Go Report Card](https://goreportcard.com/badge/github.com/naltun/wasgubata)](https://goreportcard.com/report/github.com/naltun/wasgubata) [![Shields.io](https://img.shields.io/badge/free%20software-support%20free%2Flibre%20software-yellow.svg)](https://en.wikipedia.org/wiki/Free_software) [![Shields.io](https://img.shields.io/badge/license-GPLv2-blue.svg)](https://opensource.org/licenses/GPL-2.0) [![Shields.io](https://img.shields.io/badge/developed%20on-GNU%2FLinux-purple.svg)](https://www.debian.org/releases/jessie/amd64/ch01s02.html.en)

### wasguwhat?

wasgubata fetches geolocation information on an IP address or domain.

![wasgubata in action](https://i.imgur.com/Cjct6gp.png)

### Installation
wasgubata is programmed in Go(lang). As such, please install and set up Go. Be aware that wasgubata is _not_ tested on earlier versions of Go, although it should work on any computer running 1.7+.

wasgubata also requires the `dig' tool to be installed. dig can typically be found with the ISC DNS tools (aka bind tools).

If you have `make' installed, you can simply run the command. If not, run:

```shell
$ go build
```

### TODO
I want to implement native functionality for performing what the `dig` CLI accomplishes.

### Who
wasgubata is developed by Noah Altunian (github.com/naltun).

### License
Strange name, loves freedom. wasgubata is freedom-respecting software, and is licensed under the terms of the GNU GPL v2.

Love your free/libre software. For more information on Free/Libre, Open Source Software, please visit [Wikipedia](https://en.wikipedia.org/wiki/Free_software_movement).
