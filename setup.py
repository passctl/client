from setuptools import setup, find_packages

VERSION = "0.0.2" 
DESCRIPTION = "Passctl CLI client"

setup(
        name="passctl", 
        version=VERSION,
        author="ngn13",
        description=DESCRIPTION,
        packages=find_packages(),
        install_requires=["pycryptodome", "rich", "requests"], 
        entry_points={'console_scripts': ['passctl=passctl.__init__:main']}
)
