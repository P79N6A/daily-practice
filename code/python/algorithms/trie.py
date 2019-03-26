class TrieNode:
    def __init__(self):
        self.children = {}

    def set_child(self, c):
        self.children[ord(c.lower())] = TrieNode()

    def get_child(self, c):
        return self.children.get(ord(c.lower()))


class Trie:
    def __init__(self):
        self._root = TrieNode()

    def insert(self, s):
        node = self._root
        for i in range(len(s)):
            c = s[i]
            if not node.get_child(c):
                node.set_child(c)
            node = node.get_child(c)

    def search(self, s):
        node = self._root
        for i in range(len(s)):
            c = s[i]
            if not node.get_child(c):
                return False
            node = node.get_child(c)

        return node != None


if __name__ == "__main__":
    words = ['Cookies', 'a', 'mechanism', 'for', 'persisting', 'data', 'locally', 'in',
            'browser', 'can', 'be', 'incorporated', 'into', 'your', 'React', 'project',
            'matter', 'If', 'have', 'React', 'Router', 'React', 'Redux', 'Components']
    output = ['Not Found', 'Found']

    t = Trie()

    for w in words:
        t.insert(w)


    keywords = ['react', 'redux', 'component', 'angular', 'vue.js']
    for k in keywords:
        print('{:<10}: {}'.format(k, output[t.search(k)]))

