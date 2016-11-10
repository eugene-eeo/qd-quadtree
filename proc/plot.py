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


colormap = plt.cm.gist_ncar
legend_opts = {
    'columnspacing': 1.0,
    'labelspacing':  0.0,
    'handletextpad': 0.5,
    'handlelength':  1.5,
    }


def plotitem(ax, i, data, Qs, Ds, colors):
    for q in Qs:
        N = []
        for d in Ds:
            N.append(data[q][d][i])
        ax.semilogy(Ds, N, label='$ q = %d $' % q, marker='o')
    for i, line in enumerate(ax.lines):
        line.set_color(colors[i])


def main():
    m = defaultdict(dict)
    for entry in line_parser(sys.stdin):
        m[entry['q']][entry['d']] = (entry['nodes'], entry['scanned'])

    with PdfPages('results.pdf') as pdf:
        Q = sorted(m)
        D = list(m[next(iter(m))].keys())
        num_plots = len(D)
        colors = [colormap(i) for i in np.linspace(0, 1, num_plots)]

        fig, ax = plt.subplots()
        fig.suptitle('Total nodes vs $ q $')
        ax.set_xlabel('$ d $')
        ax.set_ylabel('total nodes')
        plotitem(ax, 0, m, Q, D, colors)
        ax.legend(loc='upper left', **legend_opts)
        plt.grid(True)
        pdf.savefig()
        plt.close()

        fig, ax = plt.subplots()
        fig.suptitle('Total scanned vs $ q $')
        ax.set_xlabel('$ d $')
        ax.set_ylabel('total scanned')
        plotitem(ax, 1, m, Q, D, colors)
        ax.legend(loc='upper right', ncol=4, **legend_opts)
        plt.grid(True)
        pdf.savefig()
        plt.close()


if __name__ == '__main__':
    main()
