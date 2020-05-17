import requests
import json

from google.cloud import language
from google.oauth2 import service_account
from google.cloud.language import enums
from google.cloud.language import types

# uses a google service account via the credentials provided in the JSON
client = language.LanguageServiceClient.from_service_account_json('credentials.json')


def analyse(client, url, invalid_types = ['OTHER'], **data):

        html = load_text_from_url(url, **data)

        if not html:
            return None

        document = types.Document(
        content=html,
        type=language.enums.Document.Type.HTML )

        features = {'extract_syntax': True,
                'extract_entities': True,
                'extract_document_sentiment': True,
                'extract_entity_sentiment': True,
                'classify_text': True
                }

        response = client.annotate_text(document=document, features=features)
        sentiment = response.document_sentiment
        entities = response.entities

        response = client.classify_text(document)
        categories = response.categories

        def get_type(type):
            return client.enums.Entity.Type(entity.type).name

        result = {}

        result['sentiment'] = []
        result['entities'] = []
        result['categories'] = []

        if sentiment:
            result['sentiment'] = [{ 'magnitude': sentiment.magnitude, 'score':sentiment.score }]

        for entity in entities:
            if get_type(entity.type) not in invalid_types:
                result['entities'].append({'name': entity.name, 'type': get_type(entity.type), 'salience': entity.salience, 'wikipedia_url': entity.metadata.get('wikipedia_url', '-')  })

        for category in categories:
            result['categories'].append({'name':category.name, 'confidence': category.confidence})


        return result

#here we gather all text from given source, and this function is called in the analysis one
def load_text_from_url(url, **data):

        timeout = data.get('timeout', 20)

        results = []

        try:

            print("Grabbing text from source: {}".format(url))
            response = requests.get(url, timeout=timeout)

            text = response.text
            status = response.status_code

            if status == 200 and len(text) > 0:
                return text

            return None


        except Exception as e:
            print('Error with: {0}.'.format(url))
        return None


def result_format(url):

    analysis = analyse(client,url)

    text_category = analysis['categories'][0]['name']
    text_entities_list = []

    for i in range(len(analysis['entities'])):
        text_entities_list.append(analysis['entities'][i]['type'])

    #to remove duplicates, we convert list to a dictionary with keys formed from list elements 
    #and then convert back to list - dictionaries do not accept duplicates for keys
        
    text_entities = list(dict.fromkeys(text_entities_list))

    #crude cases here for sentiment analysis, perhaps more cases should be included
    #must include magnitude parameter

    if(analysis['sentiment'][0]['score'] == 0.0):
        sent = 'Neutral'
    elif(analysis['sentiment'][0]['score'] > 0.0):
        sent = 'Positive'
    else:
        sent = 'Negative'

    print(f"Main subject: {text_category} \nEntities: {text_entities} \nSentiment: {sent} \n")

#update the github repo to CLI format

#any url may be introduced, but as an example, we shall analyse the ethereum site
url = "https://ethereum.org/what-is-ethereum/"
result_format(url)

"""
EXAMPLE:
Grabbing text from source: https://ethereum.org/what-is-ethereum/
Main subject: /Finance/Investing/Currencies & Foreign Exchange 
Entities: ['ORGANIZATION', 'PERSON', 'WORK_OF_ART', 'LOCATION', 'CONSUMER_GOOD', 'EVENT', 'DATE', 'NUMBER'] 
Sentiment: Positive 
"""
