#!/usr/bin/env python

import os
import json

message = os.environ["MESSAGE"]
buttons = json.loads(os.environ["BUTTONS"])

result=[
  {
    "color": "good",
    "fallback": message,
    "text": message
  }
]

actions=[]

for action in buttons:
  for key, val in action.items():
    try:
      with open(val) as f:
        action[key] = f.read().strip()
    except:
      pass
  actions.append(action)

result[0]["actions"] = actions

with open('slack-message-output/attachments.json', 'w') as outfile:
    json.dump(result, outfile)
