import os
import click

@click.group
def main():

    pass

@main.command
@click.argument('url')
def issue(url):
    os.system("scavengeCLI tx dcrawl createScavenge 69foo 'Rating issued for:"+url+" --from bot")
    


if __name__ == '__main__':
    main()
