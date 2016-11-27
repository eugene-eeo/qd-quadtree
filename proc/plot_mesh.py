import sys
import json
import matplotlib.pyplot as plt


def main():
    data = json.load(sys.stdin)
    points = data['points']
    simplices = data['simplices']
    X = [x for x,y in points]
    Y = [y for x,y in points]
    plt.triplot(X, Y, simplices, color='black')
    plt.plot(X, Y, '.', color='black')
    plt.savefig('triangulation.png', dpi=200, bbox_inches='tight')


if __name__ == '__main__':
    main()
