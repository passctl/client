# Passctl - CLI Client
This is the command-line client for passctl.

## Install
### Dependencies 
- `python3`
- `pip`
### Installation
```
pip install -U https://github.com/passctl/client/archive/refs/heads/main.zip
```

## Usage
### Set a server
```
passctl server http://someserver.com
```
### Add entries
```
passctl add someentry
```
### List entries
```
passctl list
```
### Read entries
```
passctl get someentry
```
### Delete entries
```
passctl del someentry
```
### Backup your vault
```
passctl backup somefile
```
### Load from backup
```
passctl load somefile
```
