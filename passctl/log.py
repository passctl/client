from rich.console import Console
from getpass import getpass
import readline
console = Console()

def inp(text, start="", end="", pwd=False) -> str:
    console.print(f"{start}[bold][blue]■ [/blue]{text}[/]: {end}", end="")

    if pwd:
        return getpass("")
    return input()

def success(text, start="", end="") -> None:
    console.print(f"{start}[green bold]■[/] {text}{end}")

def info(text, start="", end="") -> None:
    console.print(f"{start}[cyan bold]■[/] {text}{end}")

def error(text, start="", end="") -> None:
    console.print(f"{start}[red bold]■[/] {text}{end}")
    exit()

def banner() -> None:
    console.print("""[bold]
                ┓
        ┏┓┏┓┏┏┏╋┃   github.com/passctl
        ┣┛┗┻┛┛┗┗┗   version 2.2
        ┛[/]""")
