import sys
import numpy as np
from scipy.spatial import Delaunay
from json import dumps


def main():
    n = int(sys.argv[1]) if len(sys.argv) > 1 else 1000
    # take points ~ N(0, 2.5^2)
    points = 2.5 * np.random.randn(n, 2)
    triangulation = Delaunay(points)
    for point in points[triangulation.simplices]:
        [a, b, c] = point
        print(dumps([list(k) for k in [a, b, c]]))


if __name__ == '__main__':
    main()
