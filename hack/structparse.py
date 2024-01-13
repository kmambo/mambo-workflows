import re
import sys
from pathlib import Path

outfilen = Path.joinpath(Path.home(), "tmp", "collected_types.go")
filen = sys.argv[1]
struct_re = re.compile(r'type ([A-Za-z]+) struct {')
struct_end = re.compile(r'^}')
comments_re = re.compile(r'^// ')
started, ended = False, False
with open(filen, "rt") as gofile:
    comments = []
    for line in gofile:
        if comments_re.match(line):
            comments.append(line)
        elif struct_re.match(line):
            started = True
            comments.append("// +kubebuilder:object:root=true\n")
            comments.append(line)
        elif started and not struct_end.match(line):
            comments.append(line)
        elif started and struct_end.match(line):
            comments.append(line)
            ended = True
        else:
            comments.clear()
        if ended:
            comments.append("\n")
            break

if comments:
    with open(outfilen, "at") as out:
        out.writelines(comments)
