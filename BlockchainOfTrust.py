import requests

"""

WoT API for trust rating


"""


url = "https://scorecard.api.mywot.com/v3/targets?t=ethereum.org"
headers = {'x-user-id':'USER-ID-HERE' ,'x-api-key':'API-KEY-HERE'}

wot = requests.get(url,headers=headers)
format_wot = wot.json()

try:
    print("The site: " + format_wot[0]['target'] + " is rated as: " + format_wot[0]['safety']['status'])
    print("Additional information: ")
    for i in range(len(format_wot[0]['categories'])):
        print(format_wot[0]['categories'][i]['name'])
except:
    print("Not enough information")


"""
Example:
The site ethereum.org is rated as: SAFE
Additional information:
good site

"""
