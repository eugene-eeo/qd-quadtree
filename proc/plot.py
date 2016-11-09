#!/usr/bin/env python
import sys
import matplotlib.pyplot as plt
from matplotlib.backends.backend_pdf import PdfPages
from collections import defaultdict
import numpy as np


def line_parser(lines):
    for line in lines:
        data = {}
        for item in line.split():
            key, value = item.split('=')
            data[key] = int(value)
        yield data


legend_opts = {
    'columnspacing': 1.0,
    'labelspacing':  0.0,
    'handletextpad': 0.0,
    'handlelength':  1.5,
    }


def main():
    m = defaultdict(dict)
    for entry in line_parser(sys.stdin):
        m[entry['q']][entry['d']] = (entry['nodes'], entry['scanned'])

    with PdfPages('results.pdf') as pdf:
        d = list(m[next(iter(m))].keys())
        num_plots = len(d)
        colormap = plt.cm.gist_ncar
        colors = [colormap(i) for i in np.linspace(0, 1, num_plots)]

        fig, ax = plt.subplots()
        fig.suptitle('Total nodes vs $ q $')
        ax.set_xlabel('$ d $')
        ax.set_ylabel('total nodes')

        for q in sorted(m):
            N = []
            for k in d:
                n, _ = m[q][k]
                N.append(n)
            ax.semilogy(d, N, label='$ q = %d $' % q, marker='o')

        for i, line in enumerate(ax.lines):
            line.set_color(colors[i])

        ax.legend(loc='upper left', **legend_opts)
        plt.grid(True)
        pdf.savefig()
        plt.close()

        fig, ax = plt.subplots()
        fig.suptitle('Total scanned vs $ q $')
        ax.set_xlabel('$ d $')
        ax.set_ylabel('total scanned')

        for q in sorted(m):
            N = []
            for k in d:
                _, n = m[q][k]
                N.append(n)
            ax.semilogy(d, N, label='$ q = %d $' % q, marker='o')

        for i, line in enumerate(ax.lines):
            line.set_color(colors[i])

        ax.legend(loc='upper right', ncol=4, **legend_opts)
        plt.grid(True)
        pdf.savefig()
        plt.close()


if __name__ == '__main__':
    main()
