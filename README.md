<img src="https://user-images.githubusercontent.com/13212227/67617365-74461200-f81d-11e9-9733-d64fac6df46a.png" width=100%> 

# RAS-Fuzzer (RAndom Subdomain Fuzzer)
<img src="https://img.shields.io/github/license/hahwul/ras-fuzzer.svg"> <a href="https://twitter.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/static/v1.svg?label=follow&message=hahwul&color=black"></a>
## Install 
go get
```cassandraql
$ go get github.com/hahwul/ras-fuzzer
```

or 

clone and build
```cassandraql
$ git clone https://github.com/hahwul/ras-fuzzer
$ cd ras-fuzzer
$ go build
```

## Usage 
### Options
```cassandraql
Usage of ./ras-fuzzer:
  -callback string
    	runnning command if me find
    	- Pattern: **PARAM**
    	- e.g: **PARAM**
  -length int
    	Max Length of domain (default 8)
  -target string
    	Target domain (e.g hahwul.com)
  -verbose int
    	(Not Supported) Show/Unshow Log(1=show log, 0=only result) (default 1)

```

### Default Fuzzing
```cassandraql
./ras-fuzzer -target google.com
____ ___  ___       ____ _    ___  ___  ____ ____
| . \|  \ | _\  ___ |  _\|| \ |_ \ |_ \ | __\| . \
|  <_| . \[__ \|___\| _\ ||_|\| __]| __]|  ]_|  <_
|/\_/|/\_/|___/     |/   |___/|___/|___/|___/|/\_/      by hahwul
* Fuzzing Information
* Your Target: *.google.com
* Max Length : 8
* Dictionary  : [a b c d e f g h i j k l m n o p q r s t u v w x y z 0 1 2 3 4 5 6 7 8 9]
* Combinatorial: (36+8-1)!/((36-1)!*8!) = ༼ つ ◕_◕ ༽つ <= !@#.. many case..
-----------------------------------------------------------------------------------------
[+] w.google.com
[+] d.google.com
[+] vr.google.com
[+] gg.google.com
[+] yp.google.com
[+] id.google.com
[+] 1.google.com
....
```
<img src="https://user-images.githubusercontent.com/13212227/67593375-0954f680-f79d-11e9-8149-87762348cd91.png">

### Callback
send message to slack (with callback options)
```cassandraql
./ras-fuzzer -target google.com -callback "curl -X POST --data-urlencode 'payload={\"channel\": \"#your-slack-channel\", \"username\": \"ras-fuzzer\", \"text\": \"Find subdomain : **PARAM**\", \"icon_emoji\": \":ghost:\"}' https://hooks.slack.com/services/your-slack-webhook-address"
```
<img src="https://user-images.githubusercontent.com/13212227/67593463-2689c500-f79d-11e9-9fd4-ea4ebc0c4be1.png">
<img src="https://user-images.githubusercontent.com/13212227/67593464-2689c500-f79d-11e9-9814-886f94f90c97.png">

make file (with callback options)
```cassandraql
./ras-fuzzer -target google.com -callback "touch **PARAM**"
...
[+] gg.google.com
[+] 1.google.com 
...


$ ls | grep google.com
1.google.com
gg.google.com
```

### Video
[![asciicast](https://asciinema.org/a/277040.svg)](https://asciinema.org/a/277040)
