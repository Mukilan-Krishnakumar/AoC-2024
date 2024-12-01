import re

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

if len(list_1) != len(list_2):
    raise Exception("Unequal Lists, cannot compare")

diff = 0
for num in range(len(list_1)):
    diff += abs(list_1[num] - list_2[num])

print(diff)
