<div align="center">
    <h1>passctl | CLI client</h1>
    <img src="assets/showcase.gif">
    <h3>Passctl commad-line client written in python</h3>
</div>

### Installation 
1. Install python(3) and pip.
#### Arch Linux
```
pacman -S python python-pip
```
#### Debian
```
apt install python3 python3-pip
```
2. Run the the following pip install command:
```
pip install -U https://github.com/passctl/client/archive/refs/heads/main.zip --break-system-packages
```
3. You are good to go!

### Usage
Run `passctl` or `passctl help` to get a full command list:
```
┏━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ Command ┃ Description                    ┃
┡━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┩
│ add     │ Add an entry to the vault      │
│ get     │ Read an entry from the vault   │
│ del     │ Remove an entry from the vault │
│ list    │ List the entries in the vault  │
│ server  │ Set a passctl server           │
│ pull    │ Pull the vault from the server │
│ push    │ Push local vault to the server │
└─────────┴────────────────────────────────┘
```
