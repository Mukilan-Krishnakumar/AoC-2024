import re
import collections

pattern = r"^(\d+)   (\d+)"

list_1, list_2 = list(), list()
matches = [
    (list_1.append(int(match.group(1))), list_2.append(int(match.group(2))))
    for match in (
        re.search(pattern, x.split("\n")[0], re.MULTILINE)
        for x in (open("./input.txt", "r").readlines())
    )
]

list_1 = sorted(list_1)
list_2 = sorted(list_2)

freq_list_2 = collections.Counter(list_2)

similarity_score = 0

for elem in list_1:
    val = freq_list_2.get(elem, None)
    if val:
        similarity_score += elem * val

print(similarity_score)
