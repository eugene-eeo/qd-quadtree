qd-quadtree
===========

an experiment to find the best ``q`` and ``d`` factors for quadtree-ifying
a triangle mesh in order to efficiently determine the bounding triangle.
See `this link <https://eugene-eeo.github.io/notes/triangle-mesh.html>`_ for an
introduction to the ideas behind the algorithm. Run the simulations,
generate nice graphs::

    $ make install # first time
    $ make run

- generates a 2D mesh
- partitions the quadtree with the desired qd factors and records the:

  - no. of nodes in the tree
  - total triangles scanned to locate 2000 points

- generates subjectively nice graphs
