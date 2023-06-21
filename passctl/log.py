from colorama import Fore, Back, Style, init
init(autoreset=True)

def inp(text, start="", end="\n"):
    print(f"{start}{Style.BRIGHT+Back.BLUE+Fore.WHITE}  INPUT  {Back.RESET+Fore.RESET} {text}", end=end)

def success(text, start="", end="\n"):
    print(f"{start}{Style.BRIGHT+Back.MAGENTA+Fore.WHITE} SUCCESS {Back.RESET+Fore.RESET} {text}", end=end)

def error(text, start="", end="\n"):
    print(f"{start}{Style.BRIGHT+Back.RED+Fore.WHITE}  ERROR  {Back.RESET+Fore.RESET} {text}", end=end)
    exit()
