from semantic_version import Spec, Version

def main():
    vs = ["0.2.0", "0.1.6", "0.1.5", "0.1.4", "0.1.3", "0.1.2", "0.1.1", "0.1.0"]
    s = Spec('>=0.1.3,<0.2.0')
    print('versions:',  vs)
    print('rule:', s.specs, '\n')
    versions = [Version(v) for v in vs]
    filtered = [str(v) for v in s.filter(versions)]
    print('filtered:', filtered)
    print('selected:', s.select(versions))

if __name__ == '__main__':
    main()