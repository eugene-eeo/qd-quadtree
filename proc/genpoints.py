#!/usr/bin/env python
import sys
import numpy as np
from scipy.spatial import Delaunay
from json import dumps


def main():
    n = int(sys.argv[1]) if len(sys.argv) > 1 else 2000
    # take points ~ N(0, 2.5^2)
    points = 2.5 * np.random.randn(n, 2)
    #points = np.random.uniform(size=(n,2))
    triangulation = Delaunay(points)
    print(dumps({
        "points": points.tolist(),
        "simplices": triangulation.simplices.tolist(),
    }))


if __name__ == '__main__':
    main()
