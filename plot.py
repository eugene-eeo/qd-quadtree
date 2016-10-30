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
        m[entry['q']][entry['d']] = (entry['nodes'], entry['scanned'])

    with PdfPages('results.pdf') as pdf:
        for q in sorted(m):
            d = list(m[q].keys())
            N = [a for a, _ in m[q].values()]
            T = [a for _, a in m[q].values()]
            fig, ax1 = plt.subplots()
            fig.suptitle('$ q=%d $' % q)
            ax1.plot(d, N, 'b')
            ax1.set_xlabel('$ d $')
            ax1.set_ylabel('nodes')

            ax2 = ax1.twinx()
            ax2.plot(d, T, 'g')
            ax2.set_ylabel('total scanned')

            plt.grid(True)
            pdf.savefig()
            plt.close()


if __name__ == '__main__':
    main()
