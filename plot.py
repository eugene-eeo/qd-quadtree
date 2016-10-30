#!/usr/bin/env python
import sys
import matplotlib.pyplot as plt
from matplotlib.backends.backend_pdf import PdfPages
from collections import defaultdict


def line_parser(lines):
    for line in lines:
        data = {}
        for item in line.split():
            key, value = item.split('=')
            data[key] = int(value)
        yield data


def main():
    m = defaultdict(dict)
    for entry in line_parser(sys.stdin):
        m[entry['q']][entry['d']] = entry['nodes']

    with PdfPages('results.pdf') as pdf:
        for q in sorted(m):
            plt.title('$ q=%d $' % q)
            plt.xlabel('$ d $')
            plt.ylabel('nodes')
            plt.plot(
                list(m[q].keys()),
                list(m[q].values()),
                #'ro'
                )
            plt.grid(True)
            pdf.savefig()
            plt.close()


if __name__ == '__main__':
    main()
