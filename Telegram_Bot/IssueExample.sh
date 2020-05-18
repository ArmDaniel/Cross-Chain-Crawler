
#Clarification: the user does not have to type those CLI commands, the bot translates "issue 'url'" into the following command
#and executes it


#Creates a scavenge(just an example, there can be any url served as input)
scavengeCLI tx dcrawl createScavenge 69foo "How safe is ethereum.org?" --from bot

#lists the scavenge
scavengeCLI q dcrawl get e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

: '
Example:
solutionHash is an artefact, it will be modified

{
  "creator": "cosmos1q0z497ezps254cv76c8702f3nwfu5f8nxheq2r",
  "description": "How safe is ethereum.org?",
  "solutionHash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
  "reward": [
    {
      "denom": "foo",
      "amount": "69"
    }
  ],
  "solution": "",
  "scavenger": ""
}
'
