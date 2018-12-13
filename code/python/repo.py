import os
import time
from urllib.parse import urlencode
from colorama import Fore, Back, Style
from git import Repo

def dir_exists(dir):
    return os.path.isdir(dir) and os.path.exists(dir)

def get_repo():
    local_dir = os.path.abspath(os.path.join(os.getcwd(), '../../'))
    if not dir_exists(local_dir):
        repo_url = 'https://github.com/rockdragon/julia-programming/'
        Repo.clone_from(repo_url, local_dir, branch='master')
    return Repo(local_dir)

def pull():
    get_repo().remotes.origin.pull()

def print_commit(commit):
    print('{0}commit {3} ({1}HEAD -> {2}master{0})'.format(Fore.YELLOW, \
        Fore.LIGHTBLUE_EX, Fore.GREEN, commit.hexsha,))
    print('{0}Author: {1} <{2}>'.format(Fore.WHITE, \
        commit.author.name, commit.author.email))
    print('Date: {0}'. format(commit.committed_datetime.timestamp())) 
    print('\n\t{0}'. format(commit.message))

def main():
    repo = get_repo()
    head_commit = repo.head.commit
    print_commit(head_commit)

    for k in repo.iter_commits():
        print_commit(k)

if __name__ == '__main__':
    main()