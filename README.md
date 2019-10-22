# MTMR meets ifconfig.co

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