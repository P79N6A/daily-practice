import os
from urllib.parse import urlparse
from semantic_version import Spec, Version

def is_typeof(o, repr):
    return type(o).__name__ == repr


def is_url(url):
    return urlparse(url).scheme != ''


def is_tce():
    return os.environ.get('TCE_PSM') == 'ee.infrastructure.devops'


def dir_exists(dir):
    return os.path.isdir(dir) and os.path.exists(dir)


# 从版本列表中挑出符合语义化规则的最大版本
def pick_version(versions, rule):
    spec = Spec(rule)
    vs = [Version(v) for v in versions]
    return str(spec.select(vs))


def version_compare(a, b):
    va = [int(v) for v in a.split('.')]
    vb = [int(v) for v in b.split('.')]
    for i in range(len(va)):
        if i == 2 or va[i] != vb[i]:
            return va[i] - vb[i]
    return 0


def cmp_to_version():
    class Comparator:
        def __init__(self, obj):
            self.obj = obj

        def __lt__(self, other):
            return version_compare(self.obj, other.obj) < 0

        def __gt__(self, other):
            return version_compare(self.obj, other.obj) > 0

        def __eq__(self, other):
            return version_compare(self.obj, other.obj) == 0

        def __le__(self, other):
            return version_compare(self.obj, other.obj) <= 0

        def __ge__(self, other):
            return version_compare(self.obj, other.obj) >= 0

        def __ne__(self, other):
            return version_compare(self.obj, other.obj) != 0

    return Comparator


def find_first(func, iterable):
    for x in filter(func, iterable):
        return x
    return None
