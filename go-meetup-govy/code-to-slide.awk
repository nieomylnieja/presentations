#!/usr/bin/env -S awk -f

BEGIN                     { s = 2 }
s < 1  && !/\/\/ CO$/     { gsub(/\t/, "  "); print }
s > 0                     { printf "/// %s\n", $0 }
/\/\/ CO$/                { printf "/// %s\n", $0 }
/^\)/                     { s = s - 1 }
/^$/ && s == 1            { s = 0 }
