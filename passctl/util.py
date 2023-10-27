from os import makedirs, path

def passdir():
    localdir = path.join(path.expanduser("~"), ".local", "share", "passctl")
    makedirs(localdir, exist_ok=True)
    return localdir
