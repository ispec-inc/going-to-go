"""
Does the following:
1. Inits git if used
2. Deletes dockerfiles if not going to be used
3. Deletes config utils if not needed
"""
from __future__ import print_function
import os
import shutil
from subprocess import Popen

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

def remove_file(filename):
    """
    generic remove file from project dir
    """
    fullpath = os.path.join(PROJECT_DIRECTORY, filename)
    if os.path.exists(fullpath):
        os.remove(fullpath)

def init_git():
    """
    Initialises git on the new project folder
    """
    GIT_COMMANDS = [
        ["git", "init"],
        ["git", "add", "."],
        ["git", "commit", "-a", "-m", "Initial Commit."]
    ]

    for command in GIT_COMMANDS:
        git = Popen(command, cwd=PROJECT_DIRECTORY)
        git.wait()


def remove_firebase_files():
    """
    Removes files needed for docker if it isn't going to be used
    """
    for filename in [
            "pkg/registry/firebase_repository.go",
            "pkg/infra/dao/firebase_auth.go",
            ]:
        os.remove(os.path.join(
            PROJECT_DIRECTORY, filename
        ))
    shutil.rmtree("pkg/firebase")

def remove_base_registry():
    """
    Removes files needed for docker if it isn't going to be used
    """
    for filename in [
            "pkg/registry/repository.go",
            ]:
        os.remove(os.path.join(
            PROJECT_DIRECTORY, filename
        ))

if '{{ cookiecutter.use_git }}'.lower() == 'y':
    init_git()
else:
    remove_file(".gitignore")

if '{{ cookiecutter.use_firebase_auth }}'.lower() == 'y':
    remove_base_registry()
else:
    remove_firebase_files()
