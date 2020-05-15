import requests

"""

WoT API for trust rating


"""


url = "https://scorecard.api.mywot.com/v3/targets?t=google.com"
headers = {'x-user-id':'USER-ID-HERE' ,'x-api-key':'API-KEY-HERE'}

wot = requests.get(url,headers=headers)
print(wot.json())


"""
EXAMPLE

[{'target': 'google.com', 'safety': {'status': 'SAFE', 'reputations': 94, 'confidence': 67},
'childSafety': {'reputations': 93, 'confidence': 65},
'categories': [{'id': 501, 'name': 'good site', 'confidence': 99}, {'id': 301, 'name': 'online tracking', 'confidence': 42},
{'id': 304, 'name': 'other', 'confidence': 23}]}]

TODO: give a better format to the JSON response

"""
