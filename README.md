# [MTMR](https://github.com/Toxblh/MTMR) meets [ifconfig.co](https://ifconfig.co/)

![image](https://user-images.githubusercontent.com/418868/67294329-5d867f00-f4ee-11e9-9a92-4b18773c2c12.jpg)

## Build

```bash
git clone https://github.com/a0s/ip-requester.git
cd ip-requester
go build
./ip-requester
```

## Usage

MTMR's config:

```json
{
  {
    "type": "shellScriptTitledButton",
    "width": 110,
    "refreshInterval": 5,
    "source": {
      "inline": "~/ip-requester/ip-requester"
    },
    "action": "appleScript",
    "actionAppleScript": {
      "inline": ""
    },
    "align": "right",
    "bordered": false
  }
}
```
