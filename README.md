# Cross-Chain-Crawler


![Powered by the Cosmos SDK](https://cdn-images-1.medium.com/max/1600/1*vzymD5hP-EqqdSwL2uDpJA.png)

The Cross-chain crawler is using Web of Trust and Google NL APIs to rate as many sites and services on various blockchains. The application is built using the Cosmos SDK and IBC Protocol.

By virtue of using the Cosmos SDK, the crawler will be an application where records of every analysis and reputation reports are maintained with a high degree of security and reliability, since these records are public and immutable ( any changes affect all other blocks, but this requires the consensus of the network ). Bounties may be created by users that can request for a certain site or service to be rated. Thus, the IBC protocol is used and will allow for transactions between any 2 chains to take place, such that any rating request can be paid with the desired cryptocurrency.

The advantage is perhaps one blockchain technology requires: popularisation. Since this analysis will provide secure information and a reliable report on any blockchain and all applications built on it, users will have a higher degree of confidence in using them. The benefit for the Cosmos community is evident as well: people will trust the network and use it more. Thus, the Cosmos SDK takes part in what can become a Blockchain of Trust ( or should one say : An internet of trusted blockchains ? *wink* ). 

This application is registered for the Cross-chain hackathon organized by Cosmos on Gitcoin: https://gitcoin.co/issue/cosmos/ics/408/4212  

It was made with <3 , *hard work* and **a lot** of passion by a student in-between finals and assignments.

## Plans, improvements and proposals

**Popularization**: I plan on programming a Twitter bot called 'Blockchain Spider' that regularly posts the results from crawling the internet of blockchains and thus contributes to the popularisation of the app and , implicitly, the Cosmos community.  

**Reward system**: Future development aims at rewarding users with Reputation NFTs, which may act as a badge of trust.  

**Token-curated registry**: This application can certainly contribute to such a system as the TCR, where all blockchain apps and sites are submitted for evaluation and kept in a list of trusted entities.


## Requirements and specifications

1. The scaffold tool provided by Cosmos was used to make this app: https://github.com/cosmos/scaffold
2. Cosmos SDK : https://github.com/cosmos/cosmos
3. IBC Protocol : https://github.com/cosmos/ics/tree/master/ibc 
4. Google Natural Language API: https://cloud.google.com/natural-language/docs/basics
5. Web of Trust API : https://support.mywot.com/hc/en-us/sections/360004477734-API-
6. GoLang 1.14
7. Gaia : https://hub.cosmos.network/master/gaia-tutorials/installation.html
8. Click library for Python: provides CLI functionalities in a pythonic way

If you run on Windows, then the make command can be installed with chocolatey: choco install make
