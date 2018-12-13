text = "For Base packages, check if the package you seek is listed in the built-in package manager on github, or check METADATA for registered Julia packages, then use the built-in package manager to install it after checking the requirements for respective versions and dont forget the Easter eggs!"
pattern = r"([^\s]+)"
iterator = eachmatch(pattern, text)

for regex in iterator
    print(regex.match, "\t")
end
